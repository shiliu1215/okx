package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"okex/config"
	"okex/pkg/ws"
	"okex/service/okx"
	"os"
	"os/signal"
)

func main() {
	okx.CreateOkxOrder()
	//初始化ws
	err := config.InitWs()
	if err != nil {
		return
	}
	defer config.CloseWs()
	//
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	err = okx.OkxGetKlineChannel()
	if err != nil {
		return
	}
	ws.MessageRead(config.BusinessConn)

	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupt signal received, closing connection...")
			// 发送关闭消息给服务器
			err := config.PublicConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close message error:", err)
			}
			// 等待服务器处理关闭连接
			return
		}
	}
}
