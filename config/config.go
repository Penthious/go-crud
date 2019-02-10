package config

import (
	"fmt"
	"go-crud/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=db user=postgres DB.name=go_crud_dev password=postgres sslmode=disable")
	db.AutoMigrate(&entities.User{})
	fmt.Println(db)

	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
