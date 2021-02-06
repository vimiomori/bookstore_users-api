package users

import (
	"fmt"

	"github.com/vimiomori/bookstore_users-api/datasource/mysql/users_db"
	"github.com/vimiomori/bookstore_users-api/utils/dates"
	"github.com/vimiomori/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO USERS(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	// preparing a statement is faster than executing query from client
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dates.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}
	user.ID = userID
	return nil
}
