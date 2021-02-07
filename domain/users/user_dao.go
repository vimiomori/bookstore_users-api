package users

import (
	"fmt"
	"strings"

	"github.com/vimiomori/bookstore_users-api/datasource/mysql/users_db"
	"github.com/vimiomori/bookstore_users-api/utils/dates"
	"github.com/vimiomori/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "users.email"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO USERS(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.ID),
			)
		}
		fmt.Println(err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.ID, err.Error()),
		)
	}

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
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
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
