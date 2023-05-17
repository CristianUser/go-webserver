package controller

import (
	"net/http"
	"pronesoft/server/model"
	"pronesoft/server/utils/token"
	"time"

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

	jwtToken, exp, err := token.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	session := &model.Session{
		UserId:    user.ID,
		Token:     jwtToken,
		ExpiresAt: time.Unix(exp, 0),
	}

	_, err = session.SaveSession(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
		"user":  user,
	})
}

func Logout(c *gin.Context) {

	var session model.Session
	currentSessionMap := c.GetStringMap("session")

	db, err := model.Database()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := db.Where("token= ?", currentSessionMap["token"]).First(&session).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	if err := db.Delete(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted"})
}
