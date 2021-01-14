package users

import (
	"fmt"

	"github.com/vimiomori/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user User) Get() *errors.RestErr {
	res := usersDB[user.ID]
	if res == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	return nil
}
