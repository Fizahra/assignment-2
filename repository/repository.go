package repository

import (
	"assignment2/structs"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetOrders() ([]structs.Order, error)
	Create(order structs.Order) (structs.Order, error)
	Update(order structs.Order) (structs.Order, error)
	Delete(order structs.Order) (structs.Order, error)
}

type Orderepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Orderepository {
	return Orderepository{
		db: db,
	}
}

func (or *Orderepository) GetOrders() ([]structs.Order, error) {
	var orders []structs.Order
	err := or.db.Preload("Items").Find(&orders)
	if err != nil {
		return orders, err.Error
	}
	return orders, nil
}

func (or *Orderepository) FindByID(id int) (structs.Order, error) {
	var order structs.Order
	err := or.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (or *Orderepository) Create(order structs.Order) (structs.Order, error) {
	fmt.Println(order)
	err := or.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (or *Orderepository) Update(order structs.Order) (structs.Order, error) {
	err := or.db.Updates(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (or *Orderepository) Delete(order structs.Order) (structs.Order, error) {
	err := or.db.Delete(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
