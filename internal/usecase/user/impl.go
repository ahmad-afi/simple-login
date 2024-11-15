package user

import (
	"context"
	"log"
	"net/http"
	"simple-login/internal/domain/user"
	"simple-login/internal/helper"
	"simple-login/internal/utils"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) UserUsc {
	return &UserUsecase{userRepo}
}

func (d *UserUsecase) CreateUser(ctx context.Context, params CreateUserReq) (errRes *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(params); errValidate != nil {
		log.Println(errValidate)
		return &helper.ErrorStruct{
			Err:  errValidate,
			Code: http.StatusBadRequest,
		}
	}

	// TODO PENGECEKAN BY EMAIL / USERNAME SUDAH TERPAKAI ATAU BELUM
	id, _ := uuid.NewV7()
	idStr := id.String()

	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "create user error", err)
		errRes = &helper.ErrorStruct{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
		return
	}

	err = d.userRepo.CreateUser(ctx, user.UserEntity{
		ID:       idStr,
		Name:     params.Name,
		Role:     params.Role,
		Username: params.Username,
		Password: hashedPassword,
		Email:    params.Email,
	})
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "create user error", err)
		errRes = &helper.ErrorStruct{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
		return
	}

	return
}
