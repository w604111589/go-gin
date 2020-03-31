package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/w604111589/goTest/go_websocket/impl"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnnection(wsConn); err != nil {
		goto ERR
	}

	//起一个协程，发送心跳
	go func() {
		var err error
		for {
			if err = conn.WriteMessage([]byte("heartBeat")); err != nil {
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}

func main() {
	http.HandleFunc("/ws", WsHandler)
	fmt.Println("服务已开启:localhost:8082")

	http.ListenAndServe(":8082", nil)
}
