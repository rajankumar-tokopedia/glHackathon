package coderunner

import (
	"context"
	"reflect"
	"testing"

	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
)

func Test_repo_GetIDE(t *testing.T) {
	type fields struct {
		GO coderunner.IDE
		JS coderunner.IDE
		C  coderunner.IDE
	}
	type args struct {
		ctx context.Context
		lan string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    coderunner.IDE
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				GO: tt.fields.GO,
				JS: tt.fields.JS,
				C:  tt.fields.C,
			}
			got, err := r.GetIDE(tt.args.ctx, tt.args.lan)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIDE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIDE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
