package c

import (
	"context"

	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

type C struct {
}

func New() coderunner.IDE {
	return &C{}
}

func (c *C) Run(ctx context.Context, sourceCode string, testcase model.TestCase) (float32, bool, error) {
	return 80.01, true, nil
}
