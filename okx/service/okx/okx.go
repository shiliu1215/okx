package okx

import (
	"encoding/json"
	"log"
	"okex/config"
	"okex/models"
	"okex/pkg/ws"
)

func OkxGetKlineChannel() error {
	arg := models.Args{
		Channel: "candle1H",
		InstId:  "BTC-USDT",
	}
	args := make([]models.Args, 0)
	args = append(args, arg)
	req := models.OkxReq{
		Op:   "subscribe",
		Args: args,
	}
	msg, err := json.Marshal(req)
	if err != nil {
		log.Printf("OkxGetKlineChannel json转换失败 err=%v", err)
		return err
	}
	err = ws.MessageSend(config.BusinessConn, msg)
	if err != nil {
		return err
	}
	return err
}
