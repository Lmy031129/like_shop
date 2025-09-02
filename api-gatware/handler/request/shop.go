package request

type ShopReq struct {
	UserId  int64   `gorm:"userId" gorm:"userId"`
	Title   string  `gorm:"title" gorm:"title"`
	Info    string  `gorm:"info" gorm:"info"`
	Price   float64 `gorm:"price" gorm:"price"`
	ShopNum int64   `gorm:"shopNum" gorm:"shopNum"`
	Id      int64   `gorm:"id" gorm:"id"`
	Page    int64   `gorm:"page" gorm:"page"`
	Size    int64   `gorm:"size" gorm:"size"`
}
