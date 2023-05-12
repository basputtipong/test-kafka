package main

import (
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextfnd-go.git/v2/exchangelog/mqlog"
	"test-kafka/httpserv"
	"test-kafka/infrastructure"
)

func init() {
	infrastructure.InitConfig()
	infrastructure.InitJwt()
}
func main() {
	infrastructure.InitLogger()
	infrastructure.InitError()
	infrastructure.InitMQ()
	mqlog.InitExchangeLog()
	httpserv.Run()
}
