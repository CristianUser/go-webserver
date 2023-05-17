package model

import (
	//"time"

	"errors"
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"not null;unique"`
	IdCard   string `json:"idCard"`
}

type Session struct {
	gorm.Model
	UserId         uint   `json:"userId"`
	User           User   `json:"user"`
	Token          string `json:"token"`
	LastTimeActive string `json:"lastTimeActive"`
}

func (s *Session) SaveSession(tx *gorm.DB) (*Session, error) {

	var err error

	err = tx.Create(&s).Error
	if err != nil {
		return &Session{}, err
	}
	return s, nil
}

func (u *User) SaveUser(tx *gorm.DB) (*User, error) {

	var err error

	err = tx.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func GetUserByID(uid uint) (User, error) {

	var u User
	db, err := Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	// u.PrepareGive()

	return u, nil

}
