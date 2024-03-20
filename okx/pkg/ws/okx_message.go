package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"okex/config"
)

func ReadOkxMessage() {
	go readOkxMessageByPublicConn()
	go readOkxMessageByPrivateConn()
	go readOkxMessageByBusinessConn()

}
func readOkxMessageByPublicConn() {
	for {
		t, message, err := config.PublicConn.ReadMessage()
		if err != nil {
			fmt.Println("err=", err.Error())
		}
		fmt.Printf("公共频道消息 消息类型%v \n消息=%v \n", t, string(message))
	}
}
func readOkxMessageByPrivateConn() {
	for {
		t, message, err := config.PrivateConn.ReadMessage()
		if err != nil {
			fmt.Println("err=", err.Error())
		}
		fmt.Printf("公共频道消息 消息类型%v \n消息=%v \n", t, string(message))
		fmt.Println()
	}
}
func readOkxMessageByBusinessConn() {
	for {
		// 读取服务器发送的消息
		t, message, err := config.BusinessConn.ReadMessage()
		if err != nil {
			fmt.Println("err=", err.Error())
		}
		fmt.Printf("公共频道消息 消息类型%v \n消息=%v \n", t, string(message))
	}
}
func SendMessage() {
	type WsHangQingReq struct {
		Op   string `json:"op"`
		Args []struct {
			Channel string `json:"channel"`
			InstId  string `json:"instId"`
		} `json:"args"`
	}
	req := WsHangQingReq{
		Op: "subscribe",
		//Op: "unsubscribe",
		Args: []struct {
			Channel string `json:"channel"`
			InstId  string `json:"instId"`
		}{
			{
				Channel: "tickers",
				InstId:  "BTC-USDT",
			},
		},
	}
	marshal, _ := json.Marshal(req)
	err := config.PublicConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		fmt.Println(err)
	}
}
