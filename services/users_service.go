package services

import (
	"github.com/vimiomori/bookstore_users-api/domain/users"
	"github.com/vimiomori/bookstore_users-api/utils/errors"
)

func GetUser(user_id int64) (*users.User, *errors.RestErr) {
	res := users.User{ID: user_id}
	if err := res.Get(); err != nil {
		return nil, err
	}
	return &res, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
