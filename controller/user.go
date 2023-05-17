package controller

import (
	"html"
	"log"
	"net/http"
	"pronesoft/server/model"
	"pronesoft/server/utils/token"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type NewUser struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	IdCard   string `json:"idCard" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	IdCard   string `json:"idCard"`
}

func GetUsers(c *gin.Context) {

	var users []model.User

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

func GetUser(c *gin.Context) {

	var user model.User

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := model.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func PostUser(c *gin.Context) {

	var user NewUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Println(err)
	}

	username := html.EscapeString(strings.TrimSpace(user.Username))

	newUser := model.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Password: string(hash),
		Username: username,
		IdCard:   user.IdCard,
	}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if _, err := newUser.SaveUser(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c *gin.Context) {

	var user model.User

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var updateUser UserUpdate

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(model.User{Name: updateUser.Name, LastName: updateUser.LastName, Password: updateUser.Password}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {

	var user model.User

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
