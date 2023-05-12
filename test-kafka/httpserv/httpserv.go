package httpserv

import (
	"gitdev.devops.krungthai.com/hexagonal/gin.git/v2/app"
	"gitdev.devops.krungthai.com/hexagonal/log"
	"gitdev.devops.krungthai.com/hexagonal/middleware.git/v2/middleware"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nexterr-go.git/v2"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextjwt-go.git/v2"
	"github.com/spf13/viper"
	"test-kafka/infrastructure"
	"test-kafka/listener"
)

func Run() {
	a := app.NewApp()
	m := middleware.New("", "")
	a.UseMiddleware(m.DefaultMiddlewarePack()...)
	a.UseMiddleware(nextjwt.VerifyJWT)
	a.UseMiddleware(nexterr.ErrorHandler)

	bindHealthRoute(a)
	appPort := viper.GetString("app.port")
	if appPort == "" {
		appPort = "1323"
	}
	server := a.Run(":" + appPort)

	// add consumer handle
	listener.Run()

	a.RegisterShutdownListener(server, a.RegisterListeners(), exCloseFunc)
}

func exCloseFunc() {
	if err := infrastructure.FaceReflectMqReader.Close(); err != nil {
		log.Errorf("error when close kafka %v", err)
	}

	if err := infrastructure.FaceLeadMqReader.Close(); err != nil {
		log.Errorf("error when close kafka %v", err)
	}
}
