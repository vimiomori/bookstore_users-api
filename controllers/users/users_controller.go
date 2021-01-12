package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vimiomori/bookstore_users-api/domain/users"
	"github.com/vimiomori/bookstore_users-api/services"
	"github.com/vimiomori/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	res, err := services.CreateUser(user)
	if err != nil {
		// TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, res)
}

// func GetUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
