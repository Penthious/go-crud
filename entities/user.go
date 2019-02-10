package entities

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Age   int
	Name  string `gorm:"size:255"`
	Email string `gorm:"type:varchar;unique_index"`
}
