package handler

import (
	"context"
	"errors"
	"user-srv/cmd/globar"
	__ "user-srv/cmd/proto"
	"user-srv/model"
)

type Server struct {
	__.UnimplementedUserServer
}

// 注册接口
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	var u model.User
	u.Mobile = in.Mobile
	u.Email = in.Email
	u.Password = in.Password
	if err := u.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	if err := globar.DB.Create(&u).Error; err != nil {
		return nil, errors.New("添加失败")
	}
	return &__.RegisterResp{
		Id: int64(u.ID),
	}, nil
}

// 登录接口
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {
	var u model.User
	u.Mobile = in.Mobile
	u.Email = in.Email
	u.Password = in.Password
	if err := u.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	if u.Password != in.Password {
		return nil, errors.New("密码不一致")
	}
	return &__.LoginResp{
		Id: int64(u.ID),
	}, nil
}

// 登录接口
func (s *Server) UserShow(_ context.Context, in *__.UserShowReq) (*__.UserShowResp, error) {
	var u model.User
	u.Mobile = in.Mobile
	if err := u.UserShow(globar.DB); err != nil {
		return nil, errors.New("数据获取失败")
	}
	return &__.UserShowResp{
		Id:       int64(u.ID),
		Password: u.Password,
		Email:    u.Email,
		Mobile:   u.Mobile,
	}, nil
}
