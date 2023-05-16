package model

import (
	//"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Done bool   `json:"done"`
}
