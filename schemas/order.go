package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         string `gorm:"type:char(36);primaryKey"`
	Status     string
	Notes      string
	CustomerID string   `gorm:"type:char(36);not null"`
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (order *Order) BeforeCreate(context *gorm.DB) (err error) {
	order.ID = uuid.New().String()
	return
}

type OrderResponse struct {
	ID        uuid.UUID  `json:"id"`
	Status    string     `json:"status"`
	Notes     string     `json:"notes"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
