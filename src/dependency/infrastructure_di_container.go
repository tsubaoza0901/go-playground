package dependency

import (
	"go-playground/m/v1/infrastructure"
)

// InitManageDBConn ...
func (i Injection) InitManageDBConn() infrastructure.ManageDBConn {
	return infrastructure.NewManageDBConn(i.dbConn)
}
