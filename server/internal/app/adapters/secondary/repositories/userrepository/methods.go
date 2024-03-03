package userRepository

import (
	"context"
	"fmt"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain"
	"time"
)

func (repo *UserRepository) GetUserById(id int64) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	queryString := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)

	rows, err := repo.DB.QueryContext(ctx, queryString)
	if err != nil {
		return domain.User{}, err
	}

	for rows.Next() {
		var userDTO UserDTO

		err = rows.Scan(&userDTO.Id, &userDTO.Name, &userDTO.Email, &userDTO.Password)
		if err != nil {
			return domain.User{}, err
		}

		user = userDTO.ToEntity()

		return
	}

	return
}
