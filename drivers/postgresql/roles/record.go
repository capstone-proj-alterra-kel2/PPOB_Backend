package roles

import ()

type Role struct {
	RoleID   uint   `json:"role_id" gorm:"size:100;primaryKey"`
	RoleName string `json:"role_name" `
}
