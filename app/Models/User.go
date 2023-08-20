package Models

import "gorm.io/gorm"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
	Role      Role   `gorm:"foreignKey:RoleID" json:"role"`
	RoleID    int    `json:"role_id"`
	gorm.Model
}

type UserRegister struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) getFullName() string {
	return u.FirstName + " " + u.LastName
}
