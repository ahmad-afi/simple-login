package user

import (
	"context"
)

type UserRepo interface {
	CreateUser(ctx context.Context, params UserEntity) (err error)
	GetListUser(ctx context.Context) (res []UserEntity, err error)
	UpdatetUser(ctx context.Context, id string, params UserEntity) (err error)
	DeletetUser(ctx context.Context, id string) (err error)
}
