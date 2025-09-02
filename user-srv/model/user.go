package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile   string `gorm:"type:varchar(30)" json:"mobile"`
	Password string `gorm:"type:varchar(32)" json:"password"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
}

func (u *User) GetOnfile(db *gorm.DB) error {
	return db.Model(u).Where("mobile = ?", u.Mobile).Limit(1).Find(&u).Error
}
func (u *User) Create(db *gorm.DB) error {
	return db.Model(u).Create(&u).Error
}
func (u *User) UserShow(db *gorm.DB) error {
	return db.Model(u).Where("mobile = ?", u.Mobile).Limit(1).Find(&u).Error
}
