package model

import (
	__ "user-srv/cmd/proto"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId   int    `gorm:"type:int(11)"      json:"userId"`
	Title    string `gorm:"type:varchar(30)"  json:"title"`
	VideoUrl string `gorm:"type:varchar(200)" json:"videoUrl"`
	Info     string `gorm:"type:varchar(30)"  json:"info"`
	ImageUrl string `gorm:"type:varchar(255)" json:"imageUrl"`
}

func (v *Video) GetOnfile(db *gorm.DB) error {
	return db.Model(v).Where("title = ?", v.Title).Limit(1).Find(v).Error
}
func (v *Video) Create(db *gorm.DB) error {
	return db.Model(v).Create(v).Error
}
func (v *Video) Update(db *gorm.DB) error {
	return db.Model(v).Where("id = ?", v.ID).Updates(v).Error
}
func (v *Video) Delete(db *gorm.DB) error {
	return db.Model(v).Where("id = ?", v.ID).Delete(v).Error
}
func (v *Video) Getcover(db *gorm.DB, Page, Size int64) (list []*__.Videolist, err error) {
	page := Page
	if page <= 0 {
		page = 1
	}
	pageSize := Size
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	err = db.Model(v).Offset(int(offset)).Limit(int(pageSize)).Find(&list).Error
	return list, err
}
