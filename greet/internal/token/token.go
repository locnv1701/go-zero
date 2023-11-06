package token

import (
	"encoding/json"
	"fmt"
	"greet/common/redis"
	"time"
)

type MapToken struct {
	MapToken map[int]string
}

func AddTokenToRedis(id int, token string) {
	res, ok, err := redis.RedisCache.Get("mapToken", 10*time.Hour)
	if err != nil {
		fmt.Println("Error getting", err)
	}

	mapToken := MapToken{}

	fmt.Println("ok", ok)
	if ok {
		err = json.Unmarshal(res, &mapToken)
		if err != nil {
			fmt.Println("Unmarshal error", err)
		}
		mapToken.MapToken[id] = token
		err = redis.RedisCache.Set("mapToken", mapToken, 10*time.Hour)
		if err != nil {
			fmt.Println("Error Set", err)
		}
	} else {
		mapToken.MapToken = make(map[int]string)

		mapToken.MapToken[id] = token
		err = redis.RedisCache.Set("mapToken", mapToken, 10*time.Hour)
		if err != nil {
			fmt.Println("Error Set", err)
		}
	}

	fmt.Println("after")
	mapTokenAfter := MapToken{}

	res, ok, err = redis.RedisCache.Get("mapToken", 10*time.Hour)
	if err != nil {
		fmt.Println("Error getting", err)
	}
	if ok {
		err = json.Unmarshal(res, &mapTokenAfter)
		if err != nil {
			fmt.Println("Unmarshal error", err)
		}
	}

	fmt.Println("Map token after", mapTokenAfter)
}
