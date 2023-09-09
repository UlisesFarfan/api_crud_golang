package user_controller

import (
	"net/http"
	"time"

	m "github.com/api-rest-go/models"
	user_repository "github.com/api-rest-go/repository/user.repository"
	"github.com/gin-gonic/gin"
)

// User controller functions.
func GetUsers(c *gin.Context) {
	users, err := user_repository.Read()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	User, err := user_repository.ReadById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, User)
}

func PostUser(c *gin.Context) {
	var User m.User
	if err := c.BindJSON(&User); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data."})
		return
	}
	User.CreatedAt = time.Now()
	User.UpdateAt = time.Now()
	err := user_repository.Create(User)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong."})
		return
	}
	c.IndentedJSON(http.StatusCreated, User)
}

func Update(c *gin.Context) {
	var User m.User
	var UserId string
	User.UpdateAt = time.Now()
	if err := c.BindJSON(&User); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data."})
		return
	}
	UserId = c.Param("id")
	result, err := user_repository.Update(User, UserId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, result)
}

func Delete(c *gin.Context) {
	userId := c.Param("id")
	err := user_repository.Delete(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Deleted successfuly."})
}
