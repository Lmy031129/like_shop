package model

import (
	__ "user-srv/cmd/proto"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	UserId  int64   `gorm:"type:int(11)" json:"userId"`
	Title   string  `gorm:"type:varchar(30)" json:"title"`
	Price   float64 `gorm:"type:decimal(10,2)" json:"price"`
	ShopNum int     `gorm:"type:int(11)" json:"shop_num"`
	Info    string  `gorm:"type:text" json:"info"`
}

func (s *Shop) GetOnfile(db *gorm.DB) error {
	return db.Model(s).Where("title = ?", s.Title).Limit(1).Find(s).Error
}
func (s *Shop) Create(db *gorm.DB) error {
	return db.Model(s).Create(s).Error
}
func (s *Shop) Update(db *gorm.DB) error {
	return db.Model(s).Where("id = ?", s.ID).Updates(&s).Error
}
func (s *Shop) Delete(db *gorm.DB) error {
	return db.Model(s).Where("id = ?", s.ID).Delete(&s).Error
}
func (s *Shop) Shows(db *gorm.DB) error {
	return db.Model(s).Where("id = ?", s.ID).Find(&s).Limit(1).Error
}

func (s *Shop) Getcoser(db *gorm.DB, Page, Size int64) (list []*__.Shoplist, err error) {
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
	err = db.Model(s).Offset(int(offset)).Limit(int(pageSize)).Find(&list).Error
	return list, err
}
