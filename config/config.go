package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() {
	db := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	// if err != nil {
	// 	return nil, err
	// } else {
	// 	db := session.DB(dbName)
	// 	return db, nil
	// }
}
