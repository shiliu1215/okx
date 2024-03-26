package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func MessageRead(coon *websocket.Conn) {
	for {
		t, message, err := coon.ReadMessage()
		if err != nil {
			fmt.Println("err=", err.Error())
		}
		fmt.Printf("公共频道消息 消息类型%v \n消息=%v \n", t, string(message))
	}
}

func MessageSend(coon *websocket.Conn, msg []byte) error {
	err := coon.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Printf("MessageSend 发送消息失败 err=%v", err)
	}
	return err
}
