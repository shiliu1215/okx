package models

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
type OkxCreateOrderReq struct {
	SprdId  string `json:"sprdId"`  //必填 spread ID，如 BTC-USDT_BTC-USDT-SWAP
	ClOrdId string `json:"clOrdId"` //客户自定义订单ID字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag     string `json:"tag"`     //订单标签字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
	Side    string `json:"side"`    //必填 订单方向 buy：买，sell：卖，可以是纯字母、纯数字，且长度在1-16位之间。
	OrdType string `json:"ordType"` //订单类型 limit：限价单 post_only：只做maker单 ioc：立即成交并取消剩余
	Px      string `json:"px"`      //必填 委托数量。反向价差的数量单位为USD，正向价差为其对应baseCcy
	Sz      string `json:"sz"`      //必填 委托价格，仅适用于limit, post_only, ioc类型的订单
}

type OkxCuoHetRadeReq struct {
	InstId  string `json:"instId"`
	TdMode  string `json:"tdMode"`
	ClOrdId string `json:"clOrdId"`
	Side    string `json:"side"`
	OrdType string `json:"ordType"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
}
