package controller

import (
	"log"
	"net/http"
	"pronesoft/server/model"
	"pronesoft/server/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var user model.User
	var login LoginCredentials

	db, err := model.Database()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("username= ? OR email= ?", login.Username, login.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Password not match"})
		return
	}

	jwtToken, err := token.GenerateToken(user.ID)

	log.Println("Reached this point", jwtToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	session := &model.Session{
		UserId: user.ID,
		Token:  jwtToken,
	}

	_, err = session.SaveSession(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   jwtToken,
		"session": session,
		"user":    user,
	})
}
