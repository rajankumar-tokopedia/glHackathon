package handller

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
	mock3 "github.com/rajankumar549/glHackathon/src/mocks/coderunner"
	mock2 "github.com/rajankumar549/glHackathon/src/mocks/httpparams"
	mock "github.com/rajankumar549/glHackathon/src/mocks/model"
)

func Test_repo_PostSubmissions(t *testing.T) {

	t.Run("Invalid EventID in Payload", func(t *testing.T) {
		h := &repo{}
		req1, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/problem/submission", strings.NewReader(`{
			"event_id":"1111",
			"problem_id":2,
			"source_code":"{{AllTheCode}}",
			"lang":"c"
		}`))
		_, err := h.PostSubmissions(context.Background(), req1, nil)
		if (err != nil) != true {
			t.Errorf("PostSubmissions() error = %v, wantErr %v", err, true)
			return
		}
	})

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	modelMockRepo := mock.NewMockModel(mockCtrl)
	h := &repo{
		model: modelMockRepo,
	}
	t.Run("Save Test Code Record To DB", func(t *testing.T) {

		req1, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/problem/submission", strings.NewReader(`{
			"event_id":1111,
			"problem_id":2,
			"source_code":"{{AllTheCode}}",
			"lang":"c"
		}`))
		modelMockRepo.EXPECT().InsertSubmission(context.Background(), model.Submission{
			EventId:    1111,
			ProblemId:  2,
			Status:     model.SubmissionStatus.Proscessing,
			Language:   "c",
			SourceCode: "{{AllTheCode}}",
		}).Return(int64(0), nil)
		_, err := h.PostSubmissions(context.Background(), req1, nil)
		if (err != nil) != false {
			t.Errorf("PostSubmissions() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("Unable Save Test Code Record To DB", func(t *testing.T) {

		req1, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/problem/submission", strings.NewReader(`{
			"event_id":1111,
			"problem_id":2,
			"source_code":"{{AllTheCode}}",
			"lang":"c"
		}`))
		modelMockRepo.EXPECT().InsertSubmission(context.Background(), model.Submission{
			EventId:    1111,
			ProblemId:  2,
			Status:     model.SubmissionStatus.Proscessing,
			Language:   "c",
			SourceCode: "{{AllTheCode}}",
		}).Return(int64(0), errors.New("unable To insert this record"))
		_, err := h.PostSubmissions(context.Background(), req1, nil)
		if (err != nil) != true {
			t.Errorf("PostSubmissions() error = %v, wantErr %v", err, true)
			return
		}
	})

}

func Test_repo_GetSubmissions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	httpMockrepo := mock2.NewMockHttpParams(mockCtrl)
	modelMockRepo := mock.NewMockModel(mockCtrl)
	h := &repo{
		model: modelMockRepo,
	}
	t.Run("Invalid EventID in Payload", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/auth/event/qqqq/submissions", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("qqqq")
		_, err := h.GetSubmissions(context.Background(), req1, httpMockrepo)
		if (err != nil) != true {
			t.Errorf("GetSubmissions() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("when event id is zero", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/auth/event/0/submissions", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("0")
		_, err := h.GetSubmissions(context.Background(), req1, httpMockrepo)
		if (err != nil) != true {
			t.Errorf("GetSubmissions() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("Successfully return Submissions", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/auth/event/123/submissions", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("123")
		modelMockRepo.EXPECT().GetSubmission(context.Background(), int64(123), int64(0)).Return([]model.Submission{
			{
				EventId:    1111,
				ProblemId:  2,
				Status:     model.SubmissionStatus.Proscessing,
				Language:   "c",
				SourceCode: "{{AllTheCode}}",
			},
		}, nil)
		gotResult, err := h.GetSubmissions(context.Background(), req1, httpMockrepo)
		if (err != nil) != false {
			t.Errorf("GetSubmissions() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(gotResult, []model.Submission{
			{
				EventId:    1111,
				ProblemId:  2,
				Status:     model.SubmissionStatus.Proscessing,
				Language:   "c",
				SourceCode: "{{AllTheCode}}",
			},
		}) {
			t.Errorf("GetSubmissions() gotResult = %v", gotResult)
		}
	})

	t.Run("Unable to get result from db", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/auth/event/1234/submissions", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("1234")
		modelMockRepo.EXPECT().GetSubmission(context.Background(), int64(1234), int64(0)).Return(nil, errors.New("unable to get result from db."))
		_, err := h.GetSubmissions(context.Background(), req1, httpMockrepo)
		if (err != nil) != true {
			t.Errorf("GetSubmissions() error = %v, wantErr %v", err, true)
			return
		}
	})

}

func Test_repo_GetLeaderBoard(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	httpMockrepo := mock2.NewMockHttpParams(mockCtrl)
	modelMockRepo := mock.NewMockModel(mockCtrl)
	h := &repo{
		model: modelMockRepo,
	}

	t.Run("Invalid EventID in Payload", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/event/qqqq/leaders", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("qqqq")
		_, err := h.GetLeaderBoard(context.Background(), req1, httpMockrepo)
		if (err != nil) != true {
			t.Errorf("GetLeaderBoard() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("Unable to Get Leaders from DB", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/event/123/leaders", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("1234")
		modelMockRepo.EXPECT().GetLeaders(context.Background(), int64(1234), int64(10)).Return(nil, errors.New("unable to get leadersfrom db"))
		_, err := h.GetLeaderBoard(context.Background(), req1, httpMockrepo)
		if (err != nil) != true {
			t.Errorf("GetLeaderBoard() error = %v, wantErr %v", err, true)
			return
		}
	})
	t.Run("Get valid Leaders from DB but unable to fetch submissions", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/event/123/leaders", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("1234")
		leaders := []model.Leader{
			{
				EventId: 1234,
				GroupId: 1,
				Score:   50,
				Rank:    1,
			},
			{
				EventId: 1234,
				GroupId: 2,
				Score:   50,
				Rank:    2,
			},
		}
		modelMockRepo.EXPECT().GetLeaders(context.Background(), int64(1234), int64(10)).Return(leaders, nil)
		modelMockRepo.EXPECT().GetEventSubmissionsByGroupId(context.Background(), int64(1234), int64(1), int64(10)).Return(nil, errors.New("unable to get problems from db"))
		modelMockRepo.EXPECT().GetEventSubmissionsByGroupId(context.Background(), int64(1234), int64(2), int64(10)).Return(nil, errors.New("unable to get problems from db"))
		gotResult, err := h.GetLeaderBoard(context.Background(), req1, httpMockrepo)
		if (err != nil) != false {
			t.Errorf("GetLeaderBoard() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(gotResult, leaders) {
			t.Errorf("GetSubmissions() gotResult = %v", gotResult)
		}
	})

	t.Run("Get valid Leaders from DB", func(t *testing.T) {
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/event/123/leaders", nil)
		httpMockrepo.EXPECT().ByName("eventid").Return("1234")
		leaders := []model.Leader{
			{
				EventId: 1234,
				GroupId: 1,
				Score:   50,
				Rank:    1,
			},
			{
				EventId: 1234,
				GroupId: 2,
				Score:   50,
				Rank:    2,
			},
		}
		modelMockRepo.EXPECT().GetLeaders(context.Background(), int64(1234), int64(10)).Return(leaders, nil)
		sub1 := []model.Submission{
			{
				SubId:     1,
				EventId:   1234,
				ProblemId: 1,
				GroupId:   1,
				Language:  "go",
				Accuracy:  90.00,
				UIStatus: model.Status{
					ID:    model.SubmissionStatus.Accepted,
					Label: model.StatusToLable(model.SubmissionStatus.Accepted),
				},
			},
		}
		modelMockRepo.EXPECT().GetEventSubmissionsByGroupId(context.Background(), int64(1234), int64(1), int64(10)).Return(sub1, nil)
		sub2 := []model.Submission{
			{
				SubId:     2,
				EventId:   1234,
				ProblemId: 1,
				GroupId:   2,
				Language:  "JS",
				Accuracy:  70.55,
				UIStatus: model.Status{
					ID:    model.SubmissionStatus.Accepted,
					Label: model.StatusToLable(model.SubmissionStatus.Accepted),
				},
			},
		}
		modelMockRepo.EXPECT().GetEventSubmissionsByGroupId(context.Background(), int64(1234), int64(2), int64(10)).Return(sub2, nil)
		gotResult, err := h.GetLeaderBoard(context.Background(), req1, httpMockrepo)
		if (err != nil) != false {
			t.Errorf("GetLeaderBoard() error = %v, wantErr %v", err, true)
			return
		}
		leaders[0].Submissions = sub1
		leaders[1].Submissions = sub2
		if !reflect.DeepEqual(gotResult, leaders) {
			t.Errorf("GetSubmissions() gotResult = %v", gotResult)
		}
	})
}

func Test_repo_submissionProcess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	modelMockRepo := mock.NewMockModel(mockCtrl)
	coderunnerMockRepo := mock3.NewMockCodeRunner(mockCtrl)
	ideMockRepo := mock3.NewMockIDE(mockCtrl)
	h := &repo{
		model:      modelMockRepo,
		coderunner: coderunnerMockRepo,
	}

	t.Run("Invalid Language in Payload", func(t *testing.T) {
		coderunnerMockRepo.EXPECT().GetIDE(context.Background(), "go").Return(nil, errors.New("Unable to get ide for go"))
		err := h.submissionProcess(context.Background(), PostSubmissionPayload{
			Lang: "go",
		})
		if err == nil {
			t.Errorf("submissionProcess() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("Some of test case failed and Unable to Update on db", func(t *testing.T) {
		coderunnerMockRepo.EXPECT().GetIDE(context.Background(), "go").Return(ideMockRepo, nil)
		ideMockRepo.EXPECT().Run(context.Background(), gomock.Any(), gomock.Any()).Return(float32(0.00), model.SubmissionStatus.Rejected, nil)
		modelMockRepo.EXPECT().UpdateSubmission(context.Background(), int64(123), float32(0.00), model.SubmissionStatus.Rejected).Return(errors.New("Unable to Update on db"))
		err := h.submissionProcess(context.Background(), PostSubmissionPayload{
			subId: 123,
			Lang:  "go",
		})
		if err == nil {
			t.Errorf("submissionProcess() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("Some of test case failed", func(t *testing.T) {
		coderunnerMockRepo.EXPECT().GetIDE(context.Background(), "go").Return(ideMockRepo, nil)
		ideMockRepo.EXPECT().Run(context.Background(), gomock.Any(), gomock.Any()).Return(float32(0.00), model.SubmissionStatus.Rejected, nil)
		modelMockRepo.EXPECT().UpdateSubmission(context.Background(), int64(123), float32(0.00), model.SubmissionStatus.Rejected).Return(nil)
		err := h.submissionProcess(context.Background(), PostSubmissionPayload{
			subId: 123,
			Lang:  "go",
		})
		if err != nil {
			t.Errorf("submissionProcess() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("Error while Running Code", func(t *testing.T) {
		coderunnerMockRepo.EXPECT().GetIDE(context.Background(), "go").Return(ideMockRepo, nil)
		ideMockRepo.EXPECT().Run(context.Background(), gomock.Any(), gomock.Any()).Return(float32(0), model.SubmissionStatus.Rejected, errors.New("unable to run Code"))
		err := h.submissionProcess(context.Background(), PostSubmissionPayload{
			Lang: "go",
		})
		if err == nil {
			t.Errorf("submissionProcess() error = %v, wantErr %v", err, true)
			return
		}
	})

	t.Run("All Test case passed", func(t *testing.T) {
		coderunnerMockRepo.EXPECT().GetIDE(context.Background(), "go").Return(ideMockRepo, nil)
		ideMockRepo.EXPECT().Run(context.Background(), gomock.Any(), gomock.Any()).Return(float32(70.00), model.SubmissionStatus.Accepted, nil)
		modelMockRepo.EXPECT().UpdateSubmission(context.Background(), int64(123), float32(70.00), model.SubmissionStatus.Accepted).Return(nil)
		err := h.submissionProcess(context.Background(), PostSubmissionPayload{
			subId: 123,
			Lang:  "go",
		})
		if err != nil {
			t.Errorf("submissionProcess() error = %v, wantErr %v", err, true)
			return
		}
	})
}
