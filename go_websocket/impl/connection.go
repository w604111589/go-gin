package impl

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct{
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
	closeChan chan byte
	isClose bool
	mutex sync.Mutex
}

func  InitConnnection (wsConn *websocket.Conn)(conn *Connection,err error)  {
	conn = &Connection{
		wsConn: wsConn,
		inChan: make(chan []byte ,1000),
		outChan: make( chan []byte,1000),
		closeChan: make( chan byte,1),
		isClose: false,
	}
	go conn.ReadLoop()
	go conn.WriteLoop()
	return
}

func (conn *Connection) ReadMessage()(data []byte,err error){
	// data = <- conn.inChan
	select{
	case data = <- conn.inChan:
	case <- conn.closeChan:
		err = errors.New("连接被关闭")
	}
	return	
}


func (conn *Connection) WriteMessage(data []byte) (err error){
	// conn.outChan <- data
	select{
	case conn.outChan <- data:
	case <- conn.closeChan:
		err = errors.New("连接被关闭")
	}
	return
}

func (conn *Connection) Close() {
	//线程安全，可重入的close
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mutex.Unlock()
	return
}

func(conn *Connection) ReadLoop(){
	var (
		data []byte
		err error
	)
	for {
		if _,data,err = conn.wsConn.ReadMessage();err!=nil{
			goto ERR
		}
		// conn.inChan <- data
		select{
			case conn.inChan <- data:
			case <- conn.closeChan:
			 goto ERR
		}
	}
ERR:
	conn.Close()
}

func(conn *Connection) WriteLoop(){
	var (
		data []byte
		err error
	)
	for {
		select{
			case data = <-conn.outChan:
			case <-conn.closeChan:
		 		goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage,data);err != nil {
			goto ERR;
		}
	}
ERR:
	conn.Close()
}