package submission

import (
	"context"
	"reflect"
	"testing"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func TestGetLeaderBoardInfo(t *testing.T) {
	type args struct {
		ctx       context.Context
		modelRepo model.Model
		eventId   int64
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Leader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLeaderBoardInfo(tt.args.ctx, tt.args.modelRepo, tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLeaderBoardInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLeaderBoardInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubmissionInfo(t *testing.T) {
	type args struct {
		ctx       context.Context
		modelRepo model.Model
		eventId   int64
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Submission
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSubmissionInfo(tt.args.ctx, tt.args.modelRepo, tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubmissionInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubmissionInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveSubmissionResult(t *testing.T) {
	type args struct {
		ctx        context.Context
		modelRepo  model.Model
		accuracy   float32
		language   string
		sourceCode string
		eventId    int64
		problemId  int64
		status     int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SaveSubmissionResult(tt.args.ctx, tt.args.modelRepo, tt.args.accuracy, tt.args.language, tt.args.sourceCode, tt.args.eventId, tt.args.problemId, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveSubmissionResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveSubmissionResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSubmissionResult(t *testing.T) {
	type args struct {
		ctx          context.Context
		modelRepo    model.Model
		submissionId int64
		accuracy     float32
		status       int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateSubmissionResult(tt.args.ctx, tt.args.modelRepo, tt.args.submissionId, tt.args.accuracy, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSubmissionResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
