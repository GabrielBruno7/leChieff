package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	Name      string
	Number    string
	Email     string
	Cep       string
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (o *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

type CustomerResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Number    string     `json:"number"`
	Email     string     `json:"email"`
	Cep       string     `json:"cep"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
