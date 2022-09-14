package config

import (
	"assignment2/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(structs.Order{})
	db.AutoMigrate(structs.Item{})
	return db
}