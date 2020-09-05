package coderunner

import (
	"context"
	"testing"
)

func Test_repo_GetIDE(t *testing.T) {
	type args struct {
		lan string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Get Go CodeRunner",
			args: args{
				lan: "go",
			},
		},
		{
			name: "Get Go CodeRunner",
			args: args{
				lan: "Go",
			},
		},
		{
			name: "Get Go CodeRunner",
			args: args{
				lan: "GOLANG",
			},
		},
		{
			name: "Get JS CodeRunner",
			args: args{
				lan: "JS",
			},
		},
		{
			name: "Get JS CodeRunner",
			args: args{
				lan: "JaVaSCRiPt",
			},
		},
		{
			name: "Get C CodeRunner",
			args: args{
				lan: "C",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{}
			_, err := r.GetIDE(context.Background(), tt.args.lan)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIDE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
