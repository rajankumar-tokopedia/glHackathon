package model

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
	"log"
)

type repo struct {
	db *sql.DB
}

func New(dbName string) (model.Model, error) {
	db, err := connect(dbName)
	if err != nil {
		log.Printf("Model.New Unable to Connect database :: %+v", err)
		return nil, err
	}

	return &repo{
		db: db,
	}, nil
}

func (r *repo) InsertSubmission(ctx context.Context, record model.Submission) (int64, error) {
	stm, err := r.db.Prepare(INSERT_SUBMISSION)
	if err != nil {
		log.Printf("Model.InsertSubmission Unable to Prepare Stm :: %+v", err)
		return 0, err
	}
	res, err := stm.ExecContext(ctx, record.EventId, record.GroupId, record.ProblemId, record.SourceCode, record.Language, record.Status)
	if err != nil {
		log.Printf("Model.InsertSubmission Unable to Insert Record :: %+v", err)
		return 0, err
	}

	return res.LastInsertId()
}

func (r *repo) UpdateSubmission(ctx context.Context, subId int64, acc float32, status int64) error {
	fmt.Println("subId int64, acc float32, status int64", subId, acc, status)
	stm, err := r.db.Prepare(UPDATE_SUBMISSION)
	if err != nil {
		log.Printf("Model.InsertSubmission Unable to Prepare Stm :: %+v", err)
		return err
	}
	_, err = stm.ExecContext(ctx, status, acc, subId)
	if err != nil {
		log.Printf("Model.InsertSubmission Unable to Insert Record :: %+v", err)
	}

	return err
}

func (r *repo) GetSubmission(ctx context.Context, eventId int64, groupId int64) ([]model.Submission, error) {
	var allSubmissions []model.Submission
	stm, err := r.db.Prepare(SELECT_SUBMISSION)
	if err != nil {
		log.Printf("Model.GetSubmission Unable to Prepare Stm :: %+v", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, eventId, groupId, 10, 0)
	if err != nil {
		log.Printf("Model.GetSubmission Unable to Get Record :: %+v", err)
		return nil, err
	}

	if rows == nil {
		return allSubmissions, nil
	}
	defer rows.Close()

	for rows.Next() {
		sub := model.Submission{}
		err := rows.Scan(&sub.SubId, &sub.EventId, &sub.GroupId, &sub.ProblemId, &sub.SourceCode, &sub.Language, &sub.Status, &sub.Accuracy, &sub.CreatedAt, &sub.UpdatedAt)
		sub.UIStatus = model.Status{
			Label: model.StatusToLable(sub.Status),
			ID:    sub.Status,
		}
		if err != nil {
			log.Printf("Model.GetSubmission.Scan Unable Scan Record Record :: row -> %+v err -> %+v", rows, err)
			continue
		}
		allSubmissions = append(allSubmissions, sub)
	}

	return allSubmissions, nil
}

func (r *repo) GetLeaders(ctx context.Context, eventId int64, limit int64) ([]model.Leader, error) {
	var leaderBoard []model.Leader
	stm, err := r.db.Prepare(SELECT_LEADERS)
	if err != nil {
		log.Printf("Model.GetSubmission Unable to Prepare Stm :: %+v", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, eventId, model.SubmissionStatus.Accepted, 10)
	if err != nil {
		log.Printf("Model.GetSubmission Unable to Get Record :: %+v", err)
		return nil, err
	}

	if rows == nil {
		return leaderBoard, nil
	}
	defer rows.Close()
	var rank int64 = 1
	for rows.Next() {
		l := model.Leader{}
		err := rows.Scan(&l.GroupId, &l.Score)
		if err != nil {
			log.Printf("Model.GetSubmission.Scan Unable Scan Record Record :: row -> %+v err -> %+v", rows, err)
			continue
		}
		l.EventId = eventId
		l.Rank = rank
		leaderBoard = append(leaderBoard, l)
		rank++
	}

	return leaderBoard, nil
}

func (r *repo) GetEventSubmissionsByGroupId(ctx context.Context, eventId int64, groupId int64, limit int64) ([]model.Submission, error) {
	var problems []model.Submission
	stm, err := r.db.Prepare(FIND_ALL_TOP_SOLVED_PROBLEMS)
	if err != nil {
		log.Printf("Model.GetEventProblemsByGroupId Unable to Prepare Stm :: %+v", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, eventId, groupId, model.SubmissionStatus.Accepted, 10)
	if err != nil {
		log.Printf("Model.GetEventProblemsByGroupId Unable to Get Record :: %+v", err)
		return nil, err
	}

	if rows == nil {
		return problems, nil
	}
	defer rows.Close()

	for rows.Next() {
		s := model.Submission{}
		err := rows.Scan(&s.SubId, &s.ProblemId, &s.Language, &s.Accuracy, &s.Status, &s.UpdatedAt)
		if err != nil {
			log.Printf("Model.GetEventProblemsByGroupId.Scan Unable Scan Record Record :: row -> %+v err -> %+v", rows, err)
			continue
		}
		s.UIStatus = model.Status{
			ID:    s.Status,
			Label: model.StatusToLable(s.Status),
		}
		s.EventId = eventId
		problems = append(problems, s)
	}
	return problems, nil
}
