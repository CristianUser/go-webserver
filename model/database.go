package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open("postgres://user:pass@localhost:5432/prone"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&Todo{}); err != nil {
		log.Println(err)
	}

	return db, err

}
