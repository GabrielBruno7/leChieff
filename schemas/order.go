package schemas

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	Id     string
	Status string
	Notes  string
}
