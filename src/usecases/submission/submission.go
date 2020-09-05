package submission

import (
	"context"
	"fmt"
	"log"

	"github.com/rajankumar549/glHackathon/src/constants"
	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

func SaveSubmissionResult(ctx context.Context, modelRepo model.Model, accuracy float32, language, sourceCode string, eventId, problemId, status int64) (int64, error) {
	groupID, _ := ctx.Value(constants.CTX_GROUP_ID).(int64)
	userId, _ := ctx.Value(constants.CTX_GROUP_ID).(int64)
	var submission = model.Submission{
		EventId:    eventId,
		GroupId:    groupID,
		UserId:     userId,
		ProblemId:  problemId,
		Status:     status,
		Accuracy:   accuracy,
		Language:   language,
		SourceCode: sourceCode,
	}
	return modelRepo.InsertSubmission(ctx, submission)
}

func UpdateSubmissionResult(ctx context.Context, modelRepo model.Model, submissionId int64, accuracy float32, status int64) error {
	return modelRepo.UpdateSubmission(ctx, submissionId, accuracy, status)
}

func GetSubmissionInfo(ctx context.Context, modelRepo model.Model, eventId int64) ([]model.Submission, error) {
	groupID, _ := ctx.Value(constants.CTX_GROUP_ID).(int64)
	fmt.Println(eventId, groupID)
	return modelRepo.GetSubmission(ctx, eventId, groupID)
}

func GetLeaderBoardInfo(ctx context.Context, modelRepo model.Model, eventId int64) ([]model.Leader, error) {
	leaders, err := modelRepo.GetLeaders(ctx, eventId, 10)
	if err != nil {
		log.Printf("Usecase.GetLeaders Unable to get leaders for %+v. :: %+v", eventId, err)
		return nil, err
	}

	for idx, l := range leaders {
		subs, err := modelRepo.GetEventSubmissionsByGroupId(ctx, eventId, l.GroupId, 10)
		if err != nil {
			log.Printf("Usecase.GetEventProblemsByGroupId Unable to get problems for %+v.%+v :: %+v", eventId, l.GroupId, err)
			continue
		}

		leaders[idx].Submissions = subs
	}
	return leaders, nil
}
