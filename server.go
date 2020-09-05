package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"gopkg.in/paytm/grace.v1"

	"github.com/rajankumar549/glHackathon/src/apperror"
	"github.com/rajankumar549/glHackathon/src/coderunner"
	"github.com/rajankumar549/glHackathon/src/handller"
	serverIn "github.com/rajankumar549/glHackathon/src/interfaces/server"
	"github.com/rajankumar549/glHackathon/src/middleware"
	"github.com/rajankumar549/glHackathon/src/model"
	"github.com/rajankumar549/glHackathon/src/server"
)

func main() {
	// server and http routes
	appError := apperror.Init()
	appServer := server.NewHttpServer(appError)
	modelRepo, err := model.New("testingDb2")

	if err != nil || modelRepo == nil {
		log.Println("Unable to connect DB")
		os.Exit(1)
	}
	codeRunnerRepo := coderunner.New()
	handllerRepo := handller.New(codeRunnerRepo, modelRepo)
	appServer.AddMiddleware(middleware.AuthRequest)
	//Healthcheck handler
	appServer.GET("/healthcheck", func(ctx context.Context, r *http.Request, p serverIn.HttpParams) (interface{}, error) {
		return map[string]string{
			"ping": "pong",
		}, nil
	})
	//All Open Endpoints
	appServer.GET("/api/v1/event/:eventid/leaders", handllerRepo.GetLeaderBoard)

	//All Secured Endpoints
	appServer.POST("/api/v1/auth/problem/submission", handllerRepo.PostSubmissions)
	appServer.GET("/api/v1/auth/event/:eventid/submissions", handllerRepo.GetSubmissions)

	log.Fatal(grace.Serve(":7378", appServer))
}
