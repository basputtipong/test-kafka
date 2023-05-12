package infrastructure

import (
	"gitdev.devops.krungthai.com/hexagonal/log"
)

func InitLogger() {
	log.InitGlobalLogger()
}
