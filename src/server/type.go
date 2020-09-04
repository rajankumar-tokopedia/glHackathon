package server

import (
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/rajankumar549/glHackathon/src/interfaces/apperror"
	serverRepo "github.com/rajankumar549/glHackathon/src/interfaces/server"
)

type repo struct {
	error       apperror.AppError
	router      *httprouter.Router
	middlewares []serverRepo.HttpMiddleware
}

type ResponseHeader struct {
	SPT       string `json:"process_time"`
	ST        string `json:"server_time"`
	ErrCode   string `json:"error_code,omitempty"`
	StausCode int    `json:"status_code"`
	Message   string `json:"message"`
}

type Response struct {
	Header    ResponseHeader `json:"header"`
	Data      interface{}    `json:"data"`
	startTime time.Time
}
