package schemas

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	Value       float32  `sql:"type:decimal(10,2);"`
	Type        FoodType `sql:"type:ENUM('RUSK', 'CAKE', 'PANETTONE')" gorm:"column:food_type"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
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

func (self *FoodType) Scan(value interface{}) error {
	*self = FoodType(value.([]byte))
	return nil
}

func (self FoodType) Value() (driver.Value, error) {
	return string(self), nil
}
