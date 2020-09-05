package code

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func TestEvaluateCode(t *testing.T) {
	type args struct {
		ctx        context.Context
		runner     coderunner.CodeRunner
		problemId  int64
		sourceCode string
		lan        string
		limit      time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantAccuracy float32
		wantStatus   int64
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAccuracy, gotStatus, err := EvaluateCode(tt.args.ctx, tt.args.runner, tt.args.problemId, tt.args.sourceCode, tt.args.lan, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvaluateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAccuracy != tt.wantAccuracy {
				t.Errorf("EvaluateCode() gotAccuracy = %v, want %v", gotAccuracy, tt.wantAccuracy)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("EvaluateCode() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func Test_getTestCasesForProblem(t *testing.T) {
	type args struct {
		ctx       context.Context
		problemId int64
	}
	tests := []struct {
		name string
		args args
		want []model.TestCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTestCasesForProblem(tt.args.ctx, tt.args.problemId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTestCasesForProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		ctx        context.Context
		runner     coderunner.CodeRunner
		lang       string
		sourceCode string
		cases      []model.TestCase
	}
	tests := []struct {
		name         string
		args         args
		wantAccuracy float32
		wantStatus   int64
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAccuracy, gotStatus, err := run(tt.args.ctx, tt.args.runner, tt.args.lang, tt.args.sourceCode, tt.args.cases)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAccuracy != tt.wantAccuracy {
				t.Errorf("run() gotAccuracy = %v, want %v", gotAccuracy, tt.wantAccuracy)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("run() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
