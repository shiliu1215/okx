package models

import (
	"fmt"
	"okex/utils"
)

func GetTimeUnixByOkx() {
	okxTime := utils.GetIsoTime()

	timestamp := okxTime
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = "cd25cf4b-16f0-424d-88b2-23bad119557a"
	header["OK-ACCESS-SIGN"] = utils.GetOkxSign(timestamp, "/api/v5/account/balance?ccy=BTC")
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = "fAlem13579?"

	resp, err := utils.SendGetRequest("https://www.okx.com/api/v5/account/balance", header)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(resp))
}

type OkxReq struct {
	Op   string `json:"op"`
	Args []Args `json:"args"`
}
type Args struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type OkxKChannelResp struct {
	Arg  Args       `json:"arg"`
	Data [][]string `json:"data"`
}
