package model

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func Test_repo_InsertSubmission(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	h := &repo{db: db}
	t.Run("Invalid Query", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO submission").WillReturnError(errors.New("Unable to prepare Query"))
		got, err := h.InsertSubmission(context.Background(), model.Submission{})
		if err == nil {
			t.Errorf("InsertSubmission() unexpected no error")
			return
		}
		if got != 0 {
			t.Errorf("InsertSubmission() got %v want 0", got)
			return
		}
	})

	t.Run("Successfully add Record", func(t *testing.T) {
		record := model.Submission{
			EventId:    1234,
			GroupId:    1,
			ProblemId:  1,
			SourceCode: "{{Some Code}}",
			Language:   "go",
			Status:     model.SubmissionStatus.Accepted,
		}
		mock.ExpectPrepare(`INSERT INTO submission`).ExpectExec().WithArgs(record.EventId, record.GroupId, record.ProblemId, record.SourceCode, record.Language, record.Status).WillReturnResult(sqlmock.NewResult(1, 1))

		got, err := h.InsertSubmission(context.Background(), record)
		if err != nil {
			t.Errorf("InsertSubmission() unexpected no error")
			return
		}
		if got != 1 {
			t.Errorf("InsertSubmission() got %v want 1", got)
			return
		}
	})
	t.Run("Unable to save record in db", func(t *testing.T) {
		record := model.Submission{
			EventId:    1234,
			GroupId:    1,
			ProblemId:  1,
			SourceCode: "{{Some Code}}",
			Language:   "go",
			Status:     model.SubmissionStatus.Accepted,
		}
		mock.ExpectPrepare(`INSERT INTO submission`).ExpectExec().WithArgs(record.EventId, record.GroupId, record.ProblemId, record.SourceCode, record.Language, record.Status).WillReturnError(errors.New("Unable to save record to db"))

		got, err := h.InsertSubmission(context.Background(), record)
		if err == nil {
			t.Errorf("InsertSubmission() unexpected no error")
			return
		}
		if got != 0 {
			t.Errorf("InsertSubmission() got %v want 1", got)
			return
		}
	})
}

func Test_repo_UpdateSubmission(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	h := &repo{db: db}
	t.Run("Invalid Query", func(t *testing.T) {
		mock.ExpectPrepare("UPDATE submission").WillReturnError(errors.New("Unable to prepare Query"))
		err := h.UpdateSubmission(context.Background(), 123, 0.0, model.SubmissionStatus.Rejected)
		if err == nil {
			t.Errorf("UpdateSubmission() unexpected no error")
			return
		}
	})

	t.Run("Successfully updated", func(t *testing.T) {
		mock.ExpectPrepare(`UPDATE submission `).ExpectExec().WithArgs(model.SubmissionStatus.Accepted, float32(56.0), int64(1234)).WillReturnResult(sqlmock.NewResult(1, 1))

		err := h.UpdateSubmission(context.Background(), int64(1234), float32(56.0), model.SubmissionStatus.Accepted)
		if err != nil {
			t.Errorf("UpdateSubmission() unexpected  error %+v", err)
			return
		}
	})
	t.Run("Unable to Save in DB", func(t *testing.T) {

		mock.ExpectPrepare(`UPDATE submission `).ExpectExec().WithArgs(model.SubmissionStatus.Accepted, float32(56.0), int64(1234)).WillReturnError(errors.New("unable to update record to db"))
		err := h.UpdateSubmission(context.Background(), int64(1234), float32(56.0), model.SubmissionStatus.Accepted)
		if err == nil {
			t.Errorf("UpdateSubmission() unexpected  error %+v", err)
			return
		}
	})
}

func Test_repo_GetSubmission(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	h := &repo{db: db}
	sqlRegex := "SELECT (.+) FROM submission (WHERE event_id = \\? AND group_id = \\? LIMIT \\? OFFSET)"
	t.Run("Invalid Query", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).WillReturnError(errors.New("Unable to prepare Query"))
		_, err := h.GetSubmission(context.Background(), 123, 1)
		if err == nil {
			t.Errorf("GetSubmission() unexpected no error")
			return
		}
	})

	t.Run("Successfully Get Submissions", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).ExpectQuery().
			WithArgs(int64(123), int64(1), int64(10), int64(0)).
			WillReturnRows(sqlmock.NewRows([]string{"sub_id", "event_id", "group_id", "problem_id", "source_code", "language", "status", "accuracy", "created_at", "updated_at"}).
				AddRow(1, 123, 1, 1, "{{code}}", "go", 1, 70.0, "2020-09-05 15:42:33", "2020-09-05 15:42:33"))

		gotResult, err := h.GetSubmission(context.Background(), 123, 1)
		if err != nil {
			t.Errorf("GetSubmission() unexpected  error %+v", err)
			return
		}
		if !reflect.DeepEqual(gotResult, []model.Submission{
			{
				SubId:      1,
				EventId:    123,
				GroupId:    1,
				ProblemId:  1,
				SourceCode: "{{code}}",
				Language:   "go",
				Status:     1,
				Accuracy:   70.0,
				CreatedAt:  "2020-09-05 15:42:33",
				UpdatedAt:  "2020-09-05 15:42:33",
				UIStatus: model.Status{
					Label: "accepted",
					ID:    1,
				},
			},
		}) {
			t.Errorf("GetSubmissions() gotResult = %v", gotResult)
		}

	})
	t.Run("No Row Found", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).ExpectQuery().
			WithArgs(int64(123), int64(1), int64(10), int64(0)).WillReturnRows(sqlmock.NewRows([]string{"sub_id", "event_id", "group_id", "problem_id", "source_code", "language", "status", "accuracy", "created_at", "updated_at"}))

		gotResult, err := h.GetSubmission(context.Background(), 123, 1)
		if err != nil {
			t.Errorf("GetSubmission() unexpected  error %+v", err)
			return
		}
		var expected []model.Submission
		if !reflect.DeepEqual(gotResult, expected) {
			t.Errorf("GetSubmissions() gotResult = %v", gotResult)
		}

	})
}

func Test_repo_GetEventSubmissionsByGroupId(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	h := &repo{db: db}
	sqlRegex := "SELECT (.+) FROM submission (WHERE event_id = \\? AND group_id = \\? AND status = \\? GROUP BY problem_id ORDER BY max_accuracy DESC LIMIT \\?)"
	t.Run("Invalid Query", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).WillReturnError(errors.New("Unable to prepare Query"))
		_, err := h.GetEventSubmissionsByGroupId(context.Background(), 123, 1, 10)
		if err == nil {
			t.Errorf("GetEventSubmissionsByGroupId() unexpected no error")
			return
		}
	})

	t.Run("Successfully Get Submissions by GroupId", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).ExpectQuery().
			WithArgs(int64(123), int64(1), int64(1), int64(10)).
			WillReturnRows(sqlmock.NewRows([]string{"sub_id", "problem_id", "language", "max_accuracy", "status", "updated_at"}).
				AddRow(1, 1, "go", 60.0, 1, "2020-09-05 15:42:33"))

		gotResult, err := h.GetEventSubmissionsByGroupId(context.Background(), 123, 1, 10)
		if err != nil {
			t.Errorf("GetEventSubmissionsByGroupId() unexpected  error %+v", err)
			return
		}
		if !reflect.DeepEqual(gotResult, []model.Submission{
			{
				SubId:     1,
				EventId:   123,
				ProblemId: 1,
				Language:  "go",
				Status:    1,
				Accuracy:  60.0,
				UpdatedAt: "2020-09-05 15:42:33",
				UIStatus: model.Status{
					Label: "accepted",
					ID:    1,
				},
			},
		}) {
			t.Errorf("GetEventSubmissionsByGroupId() gotResult = %v", gotResult)
		}

	})
	t.Run("No Row Found", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).ExpectQuery().
			WithArgs(int64(123), int64(1), int64(1), int64(10)).WillReturnRows(sqlmock.NewRows([]string{"sub_id", "problem_id", "language", "max_accuracy", "status", "updated_at"}))

		gotResult, err := h.GetEventSubmissionsByGroupId(context.Background(), 123, 1, 10)
		if err != nil {
			t.Errorf("GetEventSubmissionsByGroupId() unexpected  error %+v", err)
			return
		}
		var expected []model.Submission
		if !reflect.DeepEqual(gotResult, expected) {
			t.Errorf("GetEventSubmissionsByGroupId() gotResult = %v", gotResult)
		}

	})
}

func Test_repo_GetLeaders(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	h := &repo{db: db}
	sqlRegex := "SELECT group_id, SUM(.+) AS score FROM (.+) AS submission_processed GROUP BY submission_processed.group_id ORDER BY score DESC LIMIT \\?"
	t.Run("Invalid Query", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).WillReturnError(errors.New("Unable to prepare Query"))
		_, err := h.GetLeaders(context.Background(), 123, 10)
		if err == nil {
			t.Errorf("GetLeaders() unexpected no error")
			return
		}
	})

	t.Run("Successfully Get Leaders ", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).ExpectQuery().
			WithArgs(int64(123), int64(1), int64(10)).
			WillReturnRows(sqlmock.NewRows([]string{"group_id", "score"}).
				AddRow(1, 120.00).
				AddRow(2, 60.00))

		gotResult, err := h.GetLeaders(context.Background(), 123, 10)
		if err != nil {
			t.Errorf("GetLeaders() unexpected  error %+v", err)
			return
		}
		if !reflect.DeepEqual(gotResult, []model.Leader{
			{
				EventId: 123,
				GroupId: 1,
				Score:   120.00,
				Rank:    1,
			},
			{
				EventId: 123,
				GroupId: 2,
				Score:   60.00,
				Rank:    2,
			},
		}) {
			t.Errorf("GetLeaders() gotResult = %v", gotResult)
		}

	})
	t.Run("No Row Found", func(t *testing.T) {
		mock.ExpectPrepare(sqlRegex).ExpectQuery().
			WithArgs(int64(123), int64(1), int64(10)).WillReturnRows(sqlmock.NewRows([]string{"group_id", "score"}))

		gotResult, err := h.GetLeaders(context.Background(), 123, 10)
		if err != nil {
			t.Errorf("GetLeaders() unexpected  error %+v", err)
			return
		}
		var expected []model.Leader
		if !reflect.DeepEqual(gotResult, expected) {
			t.Errorf("GetLeaders() gotResult = %v", gotResult)
		}

	})
}
