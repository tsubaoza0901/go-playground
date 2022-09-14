package dependency

import (
	"gorm.io/gorm"
)

// Injection ...
type Injection struct {
	dbConn *gorm.DB
}

// NewInjection ...
func NewInjection(dbConn *gorm.DB) Injection {
	return Injection{dbConn}
}
