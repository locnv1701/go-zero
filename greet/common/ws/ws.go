package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"greet/common/helper"
	"greet/common/redis"
	"greet/internal/token"
	"net/http"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// ClientList is a map used to manage a map of clients
type ClientList map[*Client]bool

type Client struct {
	Connection *websocket.Conn
	// egress is used to avoid concurrent writes on the WebSocket
	Msg chan string
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Connection: conn,
		Msg:        make(chan string),
	}
}

var (
	Clients ClientList
	Msg     chan string
)

func init() {
	Clients = ClientList{}
	Msg = make(chan string)
	go ListEvent()
	go BoardCast()
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	} // check the http request

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("HandleWS upgrader.Upgrade(w, r, nil)")
		return
	}

	// Create New Client
	client := NewClient(conn)
	// Add the newly created client to the manager
	Clients[client] = true
}

func BoardCast() {
	for {
		// time.Sleep(2 * time.Second)
		// fmt.Println("len clients", len(Clients))
		message, ok := <-Msg
		if ok {
			for clientEle := range Clients {
				err := clientEle.Connection.WriteJSON(message)
				if err != nil {
					if !errors.Is(err, syscall.EPIPE) { // check err not broken pipe; broken pipe happen when client disconnect	ws
						fmt.Println("BoardCast clientEle.WriteJSON(message)")
					}
					clientEle.Connection.Close()
					delete(Clients, clientEle)
				}
			}
		} else {
			message := "ping"
			for clientEle := range Clients {
				err := clientEle.Connection.WriteJSON(message)
				if err != nil {
					if !errors.Is(err, syscall.EPIPE) {
						fmt.Println("BoardCast clientEle.WriteJSON(message)")
					}
					clientEle.Connection.Close()
					delete(Clients, clientEle)
				}
			}
		}
	}
}

func ListEvent() {

	// count := 0
	// for {
	// 	time.Sleep(1 * time.Second)
	// 	count += 1
	// 	Msg <- strconv.Itoa(count)
	// }

	for {
		res, ok, err := redis.RedisCache.Get("mapToken", 10*time.Hour)
		if err != nil {
			fmt.Println("Error getting", err)
		}

		mapToken := token.MapToken{}

		fmt.Println("ok", ok)
		if ok {
			err = json.Unmarshal(res, &mapToken)
			if err != nil {
				fmt.Println("Unmarshal error", err)
			}
		}
		for id, jwt := range mapToken.MapToken {
			jwtStruct, err := helper.ParseJwtToken(jwt, "AccessSecret")
			if err != nil {
				fmt.Println("Error parsing", err, jwt)
				if strings.Contains(err.Error(), "token is expired") {
					msg := "User id: " + strconv.Itoa(id) + " token expired"
					Msg <- msg
					delete(mapToken.MapToken, id)

					err = redis.RedisCache.Set("mapToken", mapToken, 10*time.Hour)
					if err != nil {
						fmt.Println("Error Set", err)
					}
				}
				continue
			}
			msg := "User id: " + strconv.Itoa(id) + " expireAt: " + jwtStruct.ExpiresAt.Time.String()
			Msg <- msg
		}
		time.Sleep(10 * time.Second)
	}
}
