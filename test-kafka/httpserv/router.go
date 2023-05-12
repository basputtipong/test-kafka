package httpserv

import (
	"gitdev.devops.krungthai.com/hexagonal/foundation.git/v2/port/app"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextfnd-go.git/v2"
	"test-kafka/infrastructure"
)

func bindHealthRoute(a app.App) {
	a.Register(
		app.AppMethodGET,
		"/health",
		nextfnd.HealthHandler(func(ctx app.AppContext) error {
			return infrastructure.PingAllDb()
		}),
	)
}
