package okx

import (
	"encoding/json"
	"fmt"
	"log"
	"okex/config"
	"okex/models"
	"okex/pkg/ws"
	"okex/utils"
	"strconv"
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

// 获取币种列表
func GetOkxCurrencyList() {
	uu := "/api/v5/account/balance"

	_, body, err := utils.OkxGet(uu)
	if err != nil {
		return
	}

	fmt.Println("Response:", string(body))

}
func CreateOkxOrder() {

	snow, err := utils.NewWorker(1)
	if err != nil {
		log.Printf("初始化订单结构失败err=%v", err)
	}

	orderId := snow.GetId()
	u := "/api/v5/trade/order"
	req := models.OkxCuoHetRadeReq{
		InstId:  "BTC-USDT",
		TdMode:  "cash",
		ClOrdId: strconv.FormatInt(orderId, 10),
		Side:    "buy",
		OrdType: "market",
		Px:      "",
		Sz:      "5",
	}
	_, body, err := utils.OkxPost(u, req)
	if err != nil {
		log.Printf("发送POST请求失败err=%v", err)
		return
	}
	fmt.Println(string(body))
}
