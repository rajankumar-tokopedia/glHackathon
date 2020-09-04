package golang

import (
	"context"

	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

type golang struct {
}

func New() coderunner.IDE {
	return &golang{}
}

func (c *golang) Run(ctx context.Context, sourceCode string, testcase model.TestCase) (float32, bool, error) {
	return 60.01, true, nil
}
