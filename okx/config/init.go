package config

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"okex/models"
)

var (
	PublicConn   *websocket.Conn
	PrivateConn  *websocket.Conn
	BusinessConn *websocket.Conn
)

func InitWs() error {
	var err error
	header := http.Header{}
	header.Add("x-simulated-trading", "1")
	//公共频道
	PublicConn, _, err = websocket.DefaultDialer.Dial(models.WsUrlPublicUrl, header)
	if err != nil {
		log.Printf("连接公共频道WS失败 err=%v", err)
	}
	//私有频道
	PrivateConn, _, err = websocket.DefaultDialer.Dial(models.WsUrlPrivateUrl, header)
	if err != nil {
		log.Printf("连接私有频道WS失败 err=%v", err)
	}
	//业务频道
	BusinessConn, _, err = websocket.DefaultDialer.Dial(models.WsUrlBusinessUrl, header)
	if err != nil {
		log.Printf("连接业务频道WS失败 err=%v", err)
	}
	if PublicConn == nil || PrivateConn == nil || BusinessConn == nil {
		return errors.New("创建ws失败 频道出现nil")
	}
	return nil
}
func CloseWs() {
	BusinessConn.Close()
	PrivateConn.Close()
	PublicConn.Close()
}
