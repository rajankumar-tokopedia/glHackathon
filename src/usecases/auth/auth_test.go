package auth

import (
	"context"
	"reflect"
	"testing"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func TestIsAuthenticated(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsAuthenticated(tt.args.ctx, tt.args.token)
			if got != tt.want {
				t.Errorf("IsAuthenticated() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IsAuthenticated() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
