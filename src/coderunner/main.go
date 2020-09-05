package coderunner

import (
	"context"
	"errors"
	"strings"

	"github.com/rajankumar549/glHackathon/src/coderunner/c"
	golang "github.com/rajankumar549/glHackathon/src/coderunner/go"
	"github.com/rajankumar549/glHackathon/src/coderunner/js"
	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
)

type repo struct {
	GO coderunner.IDE
	JS coderunner.IDE
	C  coderunner.IDE
}

func New() coderunner.CodeRunner {
	return &repo{
		C:  c.New(),
		GO: golang.New(),
		JS: js.New(),
	}
}

func (r *repo) GetIDE(ctx context.Context, lan string) (coderunner.IDE, error) {

	switch strings.ToLower(lan) {
	case "c":
		return r.C, nil
	case "go", "golang":
		return r.GO, nil
	case "js", "javascript":
		return r.JS, nil
	default:
		return nil, errors.New("invalid language selected")

	}

}
