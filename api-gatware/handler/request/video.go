package request

type VideoReq struct {
	ID       int    `json:"id" form:"id"`
	UserId   int    `json:"userId" form:"userId"`
	Title    string `json:"title" form:"title"`
	VideoUrl string `json:"videoUrl" form:"videoUrl"`
	Info     string `json:"info" form:"info"`
	ImageUrl string `json:"imageUrl" form:"imageUrl"`
	Page     int64  `json:"page" form:"page"`
	Size     int64  `json:"size" form:"size"`
}
