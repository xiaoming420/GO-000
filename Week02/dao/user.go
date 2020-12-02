package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var db *sql.DB

type User struct {
	ID    int
	Phone string
}

func GetUserById(id int) (user *User, err error) {
	user = &User{ID: id}
	err = db.QueryRow("select * from user where id=?", id).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// GO认为是错误，但是上层也需要判断。
			return user, nil
		} else {
			err = errors.Wrap(err, "failed to find user")
			return nil, err
		}
	}
	return
}
