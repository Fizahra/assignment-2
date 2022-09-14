package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_Id      uint `gorm:"primaryKey"`
	Customer_Name string
	Ordered_At    time.Time
	Items         []Item `gorm:"foreignKey:Order_Id"`
}
