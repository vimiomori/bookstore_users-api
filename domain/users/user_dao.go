package users

import (
	"fmt"

	"github.com/vimiomori/bookstore_users-api/utils/dates"
	"github.com/vimiomori/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	res := usersDB[user.ID]
	if res == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = res.ID
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.Email = res.Email
	user.DateCreated = res.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s is already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	user.DateCreated = dates.GetNowString()
	usersDB[user.ID] = user
	return nil
}
