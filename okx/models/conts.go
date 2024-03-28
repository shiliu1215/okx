package models

/*
实盘API交易地址如下：
REST：https://www.okx.com/
WebSocket公共频道：wss://ws.okx.com:8443/ws/v5/public
WebSocket私有频道：wss://ws.okx.com:8443/ws/v5/private
WebSocket业务频道：wss://ws.okx.com:8443/ws/v5/business
AWS 地址如下：
REST：https://aws.okx.com
WebSocket公共频道：wss://wsaws.okx.com:8443/ws/v5/public
WebSocket私有频道：wss://wsaws.okx.com:8443/ws/v5/private
WebSocket业务频道：wss://wsaws.okx.com:8443/ws/v5/business
模拟盘交易
目前可以进行V5 API的模拟盘交易，部分功能不支持如提币、充值、申购赎回等。
模拟盘API交易地址如下：
REST：https://www.okx.com
WebSocket公共频道：wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999
WebSocket私有频道：wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999
WebSocket业务频道：wss://wspap.okx.com:8443/ws/v5/business?brokerId=9999
*/
const (
	ApiUrl              = "https://www.okx.com"
	RESTUrl             = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	WsUrlPublicUrl      = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	WsUrlPrivateUrl     = "wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999"
	WsUrlBusinessUrl    = "wss://wspap.okx.com:8443/ws/v5/business?brokerId=9999"
	AWSWsUrlPublicUrl   = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	AWSWsUrlPrivateUrl  = "wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999"
	AWSWsUrlBusinessUrl = "wss://wspap.okx.com:8443/ws/v5/business?brokerId=9999"
)
const (
	SecretKey  = "9F69653B7E2906F30464DE4C8FD94864"
	AccessKey  = "cd25cf4b-16f0-424d-88b2-23bad119557a"
	PASSPHRASE = "fAlem13579?"
)
