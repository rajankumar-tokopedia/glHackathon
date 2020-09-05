package server

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/rajankumar549/glHackathon/src/interfaces/apperror"
	server "github.com/rajankumar549/glHackathon/src/interfaces/server"
)

func TestNewHttpServer(t *testing.T) {
	type args struct {
		err apperror.AppError
	}
	tests := []struct {
		name string
		args args
		want server.HttpServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpServer(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_next(t *testing.T) {
	type args struct {
		ctx         context.Context
		httpHandler server.HttpHandler
		server      repo
		r           *http.Request
		p           server.HttpParams
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := next(tt.args.ctx, tt.args.httpHandler, tt.args.server, tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("next() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_AddMiddleware(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		middleware server.HttpMiddleware
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_DELETE(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		path        string
		httpHandler server.HttpHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_GET(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		path        string
		httpHandler server.HttpHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func Test_repo_OPTIONS(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		path        string
		httpHandler server.HttpHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_POST(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		path        string
		httpHandler server.HttpHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_PUT(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		path        string
		httpHandler server.HttpHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_ServeFiles(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		path string
		root http.FileSystem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_ServeHTTP(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_repo_routerHandler(t *testing.T) {
	type fields struct {
		error       apperror.AppError
		router      *httprouter.Router
		middlewares []server.HttpMiddleware
	}
	type args struct {
		httpHandler server.HttpHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   httprouter.Handle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &repo{
				error:       tt.fields.error,
				router:      tt.fields.router,
				middlewares: tt.fields.middlewares,
			}
			if got := s.routerHandler(tt.args.httpHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("routerHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
