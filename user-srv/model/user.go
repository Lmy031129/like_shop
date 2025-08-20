package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile   string `gorm:"type:varchar(30)" json:"mobile,omitempty"`
	Password string `gorm:"type:varchar(32)" json:"password,omitempty"`
	Email    string `gorm:"type:varchar(30)" json:"email,omitempty"`
}
