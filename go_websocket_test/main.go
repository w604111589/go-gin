package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
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
