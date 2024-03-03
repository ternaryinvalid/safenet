package userRepository

import (
	"database/sql"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain"
)

type UserDTO struct {
	Id       sql.NullInt64
	Name     sql.NullString
	Email    sql.NullString
	Password sql.NullString
}

func (u *UserDTO) ToEntity() domain.User {
	return domain.User{
		Id:       u.Id.Int64,
		Name:     u.Name.String,
		Email:    u.Email.String,
		Password: u.Password.String,
	}
}
