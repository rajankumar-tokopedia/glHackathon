package coderunner

import (
	"context"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

type CodeRunner interface {
	GetIDE(ctx context.Context, lan string) (IDE, error)
}
type IDE interface {
	Run(ctx context.Context, sourceCode string, testcase model.TestCase) (float32, int64, error)
}
