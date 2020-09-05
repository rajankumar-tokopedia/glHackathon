package handller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rajankumar549/glHackathon/src/apperror"
	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
	serverIn "github.com/rajankumar549/glHackathon/src/interfaces/server"
	"github.com/rajankumar549/glHackathon/src/usecases/code"
	"github.com/rajankumar549/glHackathon/src/usecases/submission"
)

type repo struct {
	coderunner coderunner.CodeRunner
	model      model.Model
}

func New(cr coderunner.CodeRunner, model model.Model) *repo {
	return &repo{
		coderunner: cr,
		model:      model,
	}
}

func (h *repo) PostSubmissions(ctx context.Context, r *http.Request, p serverIn.HttpParams) (result interface{}, err error) {
	var (
		payload = PostSubmissionPayload{}
	)

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Inavlid Payload %+v", err)
		err = apperror.BadError("INV_PAYLOAD", "Request payload is Invalid")
		return
	}
	subId, err := submission.SaveSubmissionResult(ctx, h.model, 0, payload.Lang, payload.SourceCode, payload.EventId, payload.ProblemId, model.SubmissionStatus.Proscessing)
	if err != nil {
		log.Printf("Handler.PostSubmissions Unable to save submission to system : %+v", err)
		err = apperror.InternalServerError("SVR_FAILED", "Unable to process this request.")
		return nil, err
	}
	payload.subId = subId
	//Unblocking call which will work in separate like thread
	if subId > 0 {
		go h.submissionProcess(ctx, payload)
	}

	return map[string]interface{}{
		"sub_id": subId,
		"status": model.StatusToLable(model.SubmissionStatus.Proscessing),
	}, nil
}

func (h *repo) GetSubmissions(ctx context.Context, r *http.Request, p serverIn.HttpParams) (result interface{}, err error) {
	eventIdParam := p.ByName("eventid")

	eventId, err := strconv.ParseInt(eventIdParam, 10, 64)
	if err != nil || eventId <= 0 {
		log.Printf("Inavlid SubIDIn Req %+v", err)
		return nil, apperror.BadError("INV_PAYLOAD", "Invalid submission id in request url")
	}

	submissions, err := submission.GetSubmissionInfo(ctx, h.model, eventId)
	if err != nil {
		log.Printf("Handler.GetSubmissions Unable to get submission : %+v", err)
		return nil, apperror.InternalServerError("SVR_FAILED", "Unable to process this request.")
	}

	return submissions, nil
}

func (h *repo) GetLeaderBoard(ctx context.Context, r *http.Request, p serverIn.HttpParams) (result interface{}, err error) {
	eventIdParam := p.ByName("eventid")

	eventId, err := strconv.ParseInt(eventIdParam, 10, 64)
	if err != nil {
		log.Printf("Inavlid SubIDIn Req %+v", err)
		return nil, apperror.BadError("INV_PAYLOAD", "Invalid submission id in request url")
	}

	leaders, err := submission.GetLeaderBoardInfo(ctx, h.model, eventId)
	if err != nil {
		log.Printf("Handler.GetSubmissions Unable to get submission : %+v", err)
		return nil, apperror.InternalServerError("SVR_FAILED", "Unable to process this request.")
	}

	return leaders, nil
}

func (h *repo) submissionProcess(ctx context.Context, payload PostSubmissionPayload) error {
	accuracy, status, err := code.EvaluateCode(ctx, h.coderunner, payload.ProblemId, payload.SourceCode, payload.Lang, time.Second*60)
	if err != nil {
		log.Printf("Inavlid Payload %+v", err)
		return err
	}

	err = submission.UpdateSubmissionResult(ctx, h.model, payload.subId, accuracy, status)
	if err != nil {
		log.Printf("Handler.summisionProcess Unable to update submission status : %+v", err)
		return err
	}

	return nil
}
