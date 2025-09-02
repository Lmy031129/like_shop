package handler

import (
	"context"
	"errors"
	"user-srv/cmd/globar"
	__ "user-srv/cmd/proto"
	"user-srv/model"

	"gorm.io/gorm"
)

// 注册接口
func (s *Server) Shopadd(_ context.Context, in *__.ShopaddReq) (*__.ShopaddResp, error) {
	var as model.Shop
	as = model.Shop{
		UserId:  in.UserId,
		Title:   in.Title,
		Price:   float64(in.Price),
		ShopNum: int(in.ShopNum),
		Info:    in.Info,
	}
	//if in.UserId == 0 {
	//	return nil, errors.New("请前去登录 登陆之后才能够添加商品")
	//}

	if err := as.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}

	if err := as.Create(globar.DB).Error; err != nil {
		return nil, errors.New("添加失败")
	}
	return &__.ShopaddResp{
		Id: int64(as.ID),
	}, nil
}

// 注册接口
func (s *Server) ShopUpdate(_ context.Context, in *__.ShopUpdateReq) (*__.ShopUpdateResp, error) {
	var as model.Shop
	as = model.Shop{
		Model:   gorm.Model{ID: uint(in.Id)},
		UserId:  in.UserId,
		Title:   in.Title,
		Price:   float64(in.Price),
		ShopNum: int(in.ShopNum),
		Info:    in.Info,
	}

	if in.UserId == 0 {
		return nil, errors.New("请前去登录 登陆之后才能够修改商品")
	}
	if err := as.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	var u model.User
	globar.DB.Find(&u)
	if u.ID != uint(in.UserId) {
		return nil, errors.New("不是此用户发布的商品不能进行修改")
	}

	if err := as.Update(globar.DB); err != nil {
		return nil, errors.New("修改失败")
	}
	return &__.ShopUpdateResp{
		Id: int64(as.ID),
	}, nil
}

// 注册接口
func (s *Server) ShopShow(_ context.Context, in *__.ShopShowReq) (*__.ShopShowResp, error) {
	var as model.Shop
	as = model.Shop{
		Model:  gorm.Model{ID: uint(in.Id)},
		UserId: in.UserId,
	}

	if in.UserId == 0 {
		return nil, errors.New("请前去登录 登陆之后才能够修改商品")
	}
	if err := as.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	var u model.User
	globar.DB.Find(&u)
	if u.ID != uint(in.UserId) {
		return nil, errors.New("不是此用户发布的商品不能进行修改")
	}

	if err := as.Shows(globar.DB); err != nil {

	}
	return &__.ShopShowResp{
		UserId:  as.UserId,
		Title:   as.Title,
		Info:    as.Info,
		Price:   float32(as.Price),
		ShopNum: int64(as.ShopNum),
		Id:      int64(as.ID),
	}, nil
}

// 注册接口
func (s *Server) ShopDelete(_ context.Context, in *__.ShopDeleteReq) (*__.ShopDeleteResp, error) {
	var as model.Shop
	as = model.Shop{
		Model:  gorm.Model{ID: uint(in.Id)},
		UserId: in.UserId,
	}
	if in.UserId == 0 {
		return nil, errors.New("请前去登录 登陆之后才能够下架商品")
	}
	if err := as.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	var u model.User
	globar.DB.Find(&u)
	if u.ID != uint(in.UserId) {
		return nil, errors.New("不是此用户发布的商品不能进行下架")
	}

	if err := as.Delete(globar.DB); err != nil {

	}
	return &__.ShopDeleteResp{
		Id: int64(as.ID),
	}, nil
}

// 注册接口
func (s *Server) Shoplist(_ context.Context, in *__.ShoplistReq) (*__.ShoplistResp, error) {
	var as model.Shop
	if in.UserId == 0 {
		return nil, errors.New("请前去登录 登陆之后才能够下架商品")
	}
	if err := as.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	var u model.User
	globar.DB.Find(&u)
	if u.ID != uint(in.UserId) {
		return nil, errors.New("不是此用户发布的商品不能进行下架")
	}
	getcoser, err := as.Getcoser(globar.DB, in.Page, in.Size)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	return &__.ShoplistResp{
		List: getcoser,
	}, nil
}
