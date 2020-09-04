package main

import (
	"context"
	"log"
	"net/http"

	"gopkg.in/paytm/grace.v1"

	"github.com/rajankumar549/glHackathon/src/apperror"
	serverIn "github.com/rajankumar549/glHackathon/src/interfaces/server"
	"github.com/rajankumar549/glHackathon/src/server"
)

func main() {
	// server and http routes
	appError := apperror.Init()
	appServer := server.NewHttpServer(appError)

	//appServer.AddMiddleware(middleware.SetupContext)
	//models := usecase.Init()
	////models := model.InitMockModule()
	//h := handler.Init(models)
	////Healthcheck handler
	appServer.GET("/healthcheck", func(ctx context.Context, r *http.Request, p serverIn.HttpParams) (interface{}, error) {
		return map[string]string{
			"ping": "pong",
		}, nil
	})

	//appServer.POST("/api/v1/user/bookings", h.GetBookings)
	//appServer.POST("/api/v1/cabs/near", h.GetNearByCabs)
	//appServer.POST("/api/v1/cab/book", h.BookCab)

	log.Fatal(grace.Serve(":7378", appServer))
}
