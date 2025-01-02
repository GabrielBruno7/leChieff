package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductOrder struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	OrderID   string `gorm:"type:char(36);not null"`
	ProductID string `gorm:"type:char(36);not null"`
	Quantity  int    `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (productOrder *ProductOrder) BeforeCreate(context *gorm.DB) (err error) {
	productOrder.ID = uuid.New().String()
	return
}

type ProductOrderResponse struct {
	ID        uuid.UUID  `json:"id"`
	OrderID   string     `json:"order_id"`
	ProductID string     `json:"product_id"`
	Quantity  string     `json:"quantity"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
