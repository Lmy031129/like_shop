package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile   string `gorm:"type:varchar(30)" json:"mobile,omitempty"`
	Password string `gorm:"type:varchar(32)" json:"password,omitempty"`
	Email    string `gorm:"type:varchar(30)" json:"email,omitempty"`
}

func (u *User) GetOnfile(db *gorm.DB) error {
	return db.Model(&u).Where("mobile = ?", u.Mobile).Find(&u).Limit(1).Error
}
func (u *User) Create(db *gorm.DB) error {
	return db.Model(&u).Create(&u).Error
}
func (u *User) UserShow(db *gorm.DB) error {
	return db.Model(&u).Where("mobile = ?", u.Mobile).Limit(1).Find(&u).Error
}
