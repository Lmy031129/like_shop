package request

type UserReq struct {
	Mobile   string `gorm:"mobile" form:"mobile"`
	Password string `gorm:"password" form:"password"`
	Email    string `gorm:"email" form:"email"`
}
