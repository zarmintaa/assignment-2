package models

import (
	"time"
)

type Orders struct {
	Order_id      uint      `gorm:"primaryKey" json:"id"`
	Customer_name string    `gorm:"not null;type varchar(191)" json:"customer_name" binding:"required"`
	Items         []Items   `gorm:"foreignKey:OrderID; references:Order_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	Ordered_at    time.Time `gorm:"foreignKey:Order_id" json:"ordered_at"`
}
