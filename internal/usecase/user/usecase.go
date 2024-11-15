package user

import (
	"context"
	"simple-login/internal/helper"
)

type UserUsc interface {
	CreateUser(ctx context.Context, params CreateUserReq) (errRes *helper.ErrorStruct)
}
