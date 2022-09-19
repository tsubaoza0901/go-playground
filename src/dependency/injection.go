package dependency

import (
	"gorm.io/gorm"
)

// Injection ...
type Injection struct {
	db *gorm.DB
}

// NewInjection ...
func NewInjection(db *gorm.DB) Injection {
	return Injection{db}
}
