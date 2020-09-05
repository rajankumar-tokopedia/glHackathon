package model

import (
	"context"
	"time"
)

type Model interface {
	InsertSubmission(context.Context, Submission) (int64, error)
	UpdateSubmission(context.Context, int64, float32, int64) error
	GetLeaders(ctx context.Context, eventId int64, howMany int64) ([]Leader, error)
	GetSubmission(ctx context.Context, eventId int64, groupId int64) ([]Submission, error)
	GetEventProblemsByGroupId(ctx context.Context, eventId int64, groupId int64, limit int64) ([]Submission, error)
}

type Submission struct {
	SubId      int64   `json:"sub_id,omitempty"`
	EventId    int64   `json:"event_id,omitempty"`
	GroupId    int64   `json:"group_id,omitempty""`
	ProblemId  int64   `json:"problem_id,omitempty"`
	SourceCode string  `json:"source_code,omitempty"`
	Language   string  `json:"language,omitempty"`
	Status     int64   `json:"-"`
	UIStatus   Status  `json:"status,omitempty"`
	Accuracy   float32 `json:"accuracy,omitempty"`
	CreatedAt  string  `json:"updated_at,omitempty"`
	UpdatedAt  string  `json:"created_at,omitempty"`
	UserId     int64   `json:"user_id,omitempty"`
}

type User struct {
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Status   Status `json:"status"`
	GroupId  int64  `json:"group_id"`
}

type Leader struct {
	EventId  int64        `json:"event_id" db:"event_id"`
	GroupId  int64        `json:"group_id" db:"group_id"`
	Score    float32      `json:"score"`
	Rank     int64        `json:"rank"`
	Problems []Submission `json:"top_submissions"`
}

var SubmissionStatus = struct {
	Accepted    int64
	Rejected    int64
	Proscessing int64
}{
	Rejected:    -1,
	Proscessing: 0,
	Accepted:    1,
}

type Status struct {
	Label string `json:"label"`
	ID    int64  `json:"id"`
}

var UserAccountStatus = struct {
	Verified Status
	Awaiting Status
	Canceled Status
	Blocked  Status
}{
	Verified: Status{
		Label: "verified",
		ID:    1,
	},
	Awaiting: Status{
		Label: "awaiting",
		ID:    2,
	},
	Canceled: Status{
		Label: "canceled",
		ID:    3,
	},
	Blocked: Status{
		Label: "blocked",
		ID:    4,
	},
}

type TestCase struct {
	Input []interface{} `json:"input" db:"input"`
	Ouput []interface{} `json:"output" db:"output"`
	TTL   time.Time     `json:"ttl" db:"ttl"`
}

func StatusToLable(inp int64) string {
	switch inp {
	case -1:
		return "rejected"
	case 1:
		return "accepted"
	default:
		return "prosecessing"
	}
}
