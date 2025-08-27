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
func (s *Server) Videoadd(_ context.Context, in *__.VideoaddReq) (*__.VideoaddResp, error) {
	var v model.Video
	v = model.Video{
		UserId:   int(in.Userid),
		Title:    in.Title,
		VideoUrl: in.VideoUrl,
		Info:     in.Info,
		ImageUrl: in.ImageUrl,
	}
	if err := v.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("查询失败")
	}
	if err := v.Create(globar.DB); err != nil {
		return nil, errors.New("添加失败")
	}

	return &__.VideoaddResp{
		Id: int64(v.ID),
	}, nil
}

// 注册接口
func (s *Server) Videoupdate(_ context.Context, in *__.VideoupdateReq) (*__.VideoupdateResp, error) {
	var v model.Video
	v = model.Video{
		Model:  gorm.Model{ID: uint(in.Id)},
		UserId: int(in.Userid),
	}
	if err := v.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("查询失败")
	}
	if err := v.Update(globar.DB); err != nil {
		return nil, errors.New("修改失败")
	}

	return &__.VideoupdateResp{
		Id: int64(v.ID),
	}, nil
}

// 注册接口
func (s *Server) Videodelete(_ context.Context, in *__.VideodeleteReq) (*__.VideodeleteResp, error) {
	var v model.Video
	v = model.Video{
		Model:  gorm.Model{ID: uint(in.Id)},
		UserId: int(in.Userid),
	}
	if err := v.GetOnfile(globar.DB); err != nil {
		return nil, errors.New("查询失败")
	}
	if err := v.Delete(globar.DB); err != nil {
		return nil, errors.New("删除失败")
	}

	return &__.VideodeleteResp{
		Id: int64(v.ID),
	}, nil
}

// 注册接口
func (s *Server) Videolist(_ context.Context, in *__.VideolistReq) (*__.VideolistResp, error) {
	var v model.Video
	getcover, err := v.Getcover(globar.DB, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	return &__.VideolistResp{
		List: getcover,
	}, nil
}
