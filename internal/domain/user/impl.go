package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserDomain struct {
	pg *sqlx.DB
}

func NewUserDomain(pg *sqlx.DB) UserRepo {
	return &UserDomain{pg}
}

func (d *UserDomain) CreateUser(ctx context.Context, params UserEntity) (err error) {
	_, err = d.pg.NamedExecContext(ctx, `INSERT INTO users (id, name, role, username, password, email) 
	VALUES(:id, :name, :role, :username, :password, :email)`, params)
	return
}

func (d *UserDomain) GetListUser(ctx context.Context) (res []UserEntity, err error) {
	query := `select id, name, role, username, password, email, created_at, updated_at from users where deleted_at is null`
	err = d.pg.SelectContext(ctx, &res, query)
	return
}

func (d *UserDomain) UpdatetUser(ctx context.Context, id string, params UserEntity) (err error) {
	query := `update product set name = $1 where id = $2`
	_, err = d.pg.ExecContext(ctx, query, params.Name, params.ID)
	return
}

func (d *UserDomain) DeletetUser(ctx context.Context, id string) (err error) {
	query := `update product set updated_at = now() where id = id`
	_, err = d.pg.ExecContext(ctx, query, id)
	return
}
