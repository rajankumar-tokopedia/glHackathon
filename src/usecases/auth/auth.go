package auth

import (
	"context"

	"github.com/rajankumar549/glHackathon/src/interfaces/model"
)

var fakeTokenToUser = map[string]model.User{
	"shdfhfvjdbgh44nre": {
		UserId:   123,
		Name:     "Rajan Kumar",
		Email:    "rajankumar549@gmail.com",
		UserName: "rajankumar549",
		Status:   model.UserAccountStatus.Verified,
		GroupId:  1,
	},
	"dhjfghefbkjerg555632bdsg3": {
		UserId:   1234,
		Name:     "Rajan Kumar2",
		Email:    "rajankumar5492@gmail.com",
		UserName: "rajankumar5492",
		Status:   model.UserAccountStatus.Verified,
		GroupId:  2,
	},
}

func IsAuthenticated(ctx context.Context, token string) (bool, model.User) {
	user, isExist := fakeTokenToUser[token]
	if !isExist {
		return false, user
	}
	switch user.Status {
	case model.UserAccountStatus.Verified:
		return true, user
	default:
		return false, user
	}

}
