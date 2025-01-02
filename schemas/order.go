package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         string `gorm:"type:char(36);primaryKey"`
	Status     string
	CustomerID string         `gorm:"type:char(36);not null"`
	Customer   Customer       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index;" json:"-"`
}

func (order *Order) BeforeCreate(context *gorm.DB) (err error) {
	order.ID = uuid.New().String()
	return
}

type OrderResponse struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}
