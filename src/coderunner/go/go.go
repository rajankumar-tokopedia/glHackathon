package golang

import (
	"context"

	"github.com/rajankumar549/glHackathon/src/helper"
	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

type golang struct {
}

func New() coderunner.IDE {
	return &golang{}
}

func (c *golang) Run(ctx context.Context, sourceCode string, testcase model.TestCase) (float32, int64, error) {

	//TODO :: find accuracy of go code by running code with test cases

	//For now we are just sending random accuracy [1-100] with status as accepted
	return helper.GetRandomFloats(), model.SubmissionStatus.Accepted, nil
}
