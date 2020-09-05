package code

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rajankumar549/glHackathon/src/interfaces/coderunner"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func EvaluateCode(ctx context.Context, runner coderunner.CodeRunner, problemId int64, sourceCode string, lan string, limit time.Duration) (accuracy float32, status int64, err error) {
	testcases := getTestCasesForProblem(ctx, problemId)
	if len(testcases) == 0 {
		return 0, model.SubmissionStatus.Rejected, errors.New("no test case found for this problem")
	}

	return run(ctx, runner, lan, sourceCode, testcases)
}

func getTestCasesForProblem(ctx context.Context, problemId int64) []model.TestCase {

	return []model.TestCase{
		{
			Input: []interface{}{1, 2, 3},
			Ouput: []interface{}{2, 3, 4},
		},
	}
}

func run(ctx context.Context, runner coderunner.CodeRunner, lang, sourceCode string, cases []model.TestCase) (accuracy float32, status int64, err error) {
	ide, err := runner.GetIDE(ctx, lang)
	if err != nil || ide == nil {
		return 0, model.SubmissionStatus.Proscessing, fmt.Errorf("unable to find ide for lang %+v", lang)
	}
	totalNoCases := len(cases)
	var acc float32 = 0
	var faildAnyTest bool = false
	for _, tcase := range cases {
		thisTestAccuracy, status, err := ide.Run(ctx, sourceCode, tcase)
		if err != nil {
			return 0, status, errors.New("TesCase Failed")
		}

		if status == model.SubmissionStatus.Accepted {
			acc += thisTestAccuracy
		} else if status == model.SubmissionStatus.Rejected {
			faildAnyTest = true
		}
	}

	if totalNoCases > 0 {
		acc = (acc / float32(totalNoCases))
	}

	if faildAnyTest {
		return acc, model.SubmissionStatus.Rejected, nil
	}

	return acc, model.SubmissionStatus.Accepted, nil

}
