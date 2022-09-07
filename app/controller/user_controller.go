package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-challenge-financial-chat/app/db"
	"go-challenge-financial-chat/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context) {

	//TODO move to model
	var userInput struct {
		Nickname string
		Password string
	}

	if err := c.BindJSON(&userInput); err != nil {
		return
	}

	t := time.Now()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Nickname), 8)
	user := models.User{
		Nickname: userInput.Nickname,
		Password: hashedPassword,
	}
	user.CreatedOn = t
	user.UpdatedOn = t
	//TODO can be improved
	if r := db.CONN.Create(&user); r.Error != nil {
		var err models.ValidationErrors
		if ok := errors.As(r.Error, &err); ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		}

		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetUser(c *gin.Context) {
	userID := c.Param("userID")

	var user models.User
	//TODO can be improved
	if r := db.CONN.First(&user, userID); r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, &user)
}
