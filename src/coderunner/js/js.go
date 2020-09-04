package js

import (
	"context"

	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

type js struct {
}

func New() coderunner.IDE {
	return &js{}
}

func (c *js) Run(ctx context.Context, sourceCode string, testcase model.TestCase) (float32, bool, error) {
	return 90.01, true, nil
}
