package handller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rajankumar549/glHackathon/src/apperror"
	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	serverIn "github.com/rajankumar549/glHackathon/src/interfaces/server"
	"github.com/rajankumar549/glHackathon/src/usecases/code"
)

type repo struct {
	coderunner coderunner.CodeRunner
}

func New(cr coderunner.CodeRunner) *repo {
	return &repo{
		coderunner: cr,
	}
}
func (h *repo) PostSubmissions(ctx context.Context, r *http.Request, p serverIn.HttpParams) (result interface{}, err error) {
	var (
		payload PostSubmissionPayload = PostSubmissionPayload{}
	)

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Inavlid Payload %+v", err)
		err = apperror.BadError("INV_PAYLOAD", "Request payload is Invalid")
		return
	}
	go h.summisionProcess(ctx, payload)
	return
}

func (h *repo) summisionProcess(ctx context.Context, payload PostSubmissionPayload) {
	accuracy, status, err := code.EvaluateCode(ctx, h.coderunner, payload.ProblemId, payload.SourceCode, payload.Lang, time.Second*60)
	if err != nil {

	}
}
