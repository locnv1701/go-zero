package ws

import (
	"fmt"
	"greet/common/ws"
	"greet/internal/svc"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		} // check the http request

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("HandleWS upgrader.Upgrade(w, r, nil)")
			return
		}

		// Create New Client
		client := ws.NewClient(conn)
		// Add the newly created client to the manager
		ws.Clients[client] = true
	}
}
