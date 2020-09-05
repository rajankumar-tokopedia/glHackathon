package js

import (
	"context"
	"testing"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func Test_js_Run(t *testing.T) {
	type args struct {
		sourceCode string
		testcase   model.TestCase
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		want1   int64
		wantErr bool
	}{
		{
			name:  "Accuracy should less then or equals 100%",
			want1: model.SubmissionStatus.Accepted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &js{}
			got, got1, err := c.Run(context.Background(), tt.args.sourceCode, tt.args.testcase)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got > 100.00 {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
