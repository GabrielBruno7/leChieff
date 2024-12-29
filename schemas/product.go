package schemas

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          string `gorm:"type:char(36);primaryKey"`
	Name        string
	Description string
	Value       float32  `sql:"type:decimal(10,2);"`
	Type        foodType `sql:"type:ENUM('RUSK', 'CAKE', 'PANETTONE')" gorm:"column:food_type"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (product *Product) BeforeCreate(context *gorm.DB) (err error) {
	product.ID = uuid.New().String()
	return
}

type ProductResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Value       float32    `json:"value"`
	Type        string     `json:"type"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type foodType string

const (
	RUSK      foodType = "RUSK"
	CAKE      foodType = "CAKE"
	PANETTONE foodType = "PANETTONE"
)

func (self *foodType) Scan(value interface{}) error {
	*self = foodType(value.([]byte))
	return nil
}

func (self foodType) Value() (driver.Value, error) {
	return string(self), nil
}
