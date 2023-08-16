package Models

import "gorm.io/gorm"

type Role struct {
	Name        string `gorm:"type:varchar(255);not null"`
	ChildRole   *Role  `gorm:"foreignKey:ChildRoleID"`
	ChildRoleID int    `gorm:"default:null"`
	gorm.Model
}

type Permission struct {
	Name string `gorm:"type:varchar(255);not null"`
	gorm.Model
}

type RolePermission struct {
	Role         *Role `gorm:"foreignKey:RoleID"`
	RoleID       int
	Permission   *Permission `gorm:"foreignKey:PermissionID"`
	PermissionID int
	gorm.Model
}
