package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_Id      uint      `gorm:"id" json:"id"`
	Customer_Name string    `gorm:"customer_name" json:"customer_name"`
	Ordered_At    time.Time `gorm:"ordered_at" json:"ordered_at"`
	Items         []Item    `gorm:"foreignKey:Order_Id;association_foreignkey:Item_Id;" json:"items"`
}
