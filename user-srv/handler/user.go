package handler

import (
	"context"
	__ "user-srv/cmd/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {

	return &__.RegisterResp{}, nil
}
