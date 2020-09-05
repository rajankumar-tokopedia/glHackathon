package handller

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/rajankumar549/glHackathon/src/interfaces/server"
)

func Test_repo_PostSubmissions(t *testing.T) {
	type args struct {
		r *http.Request
		p server.HttpParams
	}
	json.Marshal(``)
	req1, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/problem/submission", strings.NewReader(`{
		"event_id":"1111",
		"problem_id":2,
		"source_code":"{{AllTheCode}}",
		"lang":"c"
	}`))

	req2, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/problem/submission", strings.NewReader(`{
		"event_id":1111,
		"problem_id":2,
		"source_code":"{{AllTheCode}}",
		"lang":"c"
	}`))

	tests := []struct {
		name       string
		args       args
		wantResult interface{}
		wantErr    bool
	}{
		{
			name: "Invalid EventID in Payload",
			args: args{
				r: req1,
			},
			wantErr: true,
		},
		{
			name: "Invalid EventID in Payload",
			args: args{
				r: req2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &repo{}
			gotResult, err := h.PostSubmissions(context.Background(), tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostSubmissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PostSubmissions() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
