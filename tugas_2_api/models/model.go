package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID           uint      `json:"order_id" gorm:"primaryKey;autoIncrement;"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}

type Item struct {
	gorm.Model
	ID          uint   `json:"item_id" gorm:"primaryKey;autoIncrement;"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	Order       *Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
}
