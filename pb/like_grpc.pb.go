// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: like.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LikeClient is the client API for Like service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeClient interface {
	// 点赞/取消赞
	DoLike(ctx context.Context, in *DoLikeReq, opts ...grpc.CallOption) (*DoLikeResp, error)
	// 获取用户是否点赞
	GetUserLike(ctx context.Context, in *GetUserLikedReq, opts ...grpc.CallOption) (*GetUserLikedResp, error)
	// 获取目标点赞数
	GetTargetLikes(ctx context.Context, in *GetTargetLikesReq, opts ...grpc.CallOption) (*GetTargetLikesResp, error)
	GetTopK(ctx context.Context, in *GetTopKReq, opts ...grpc.CallOption) (*GetTopKResp, error)
	AddScore(ctx context.Context, in *AddScoreReq, opts ...grpc.CallOption) (*AddScoreResp, error)
	DailyUpdate(ctx context.Context, in *DailyUpdateReq, opts ...grpc.CallOption) (*DailyUpdateResp, error)
}

type likeClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeClient(cc grpc.ClientConnInterface) LikeClient {
	return &likeClient{cc}
}

func (c *likeClient) DoLike(ctx context.Context, in *DoLikeReq, opts ...grpc.CallOption) (*DoLikeResp, error) {
	out := new(DoLikeResp)
	err := c.cc.Invoke(ctx, "/like.like/DoLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) GetUserLike(ctx context.Context, in *GetUserLikedReq, opts ...grpc.CallOption) (*GetUserLikedResp, error) {
	out := new(GetUserLikedResp)
	err := c.cc.Invoke(ctx, "/like.like/GetUserLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) GetTargetLikes(ctx context.Context, in *GetTargetLikesReq, opts ...grpc.CallOption) (*GetTargetLikesResp, error) {
	out := new(GetTargetLikesResp)
	err := c.cc.Invoke(ctx, "/like.like/GetTargetLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) GetTopK(ctx context.Context, in *GetTopKReq, opts ...grpc.CallOption) (*GetTopKResp, error) {
	out := new(GetTopKResp)
	err := c.cc.Invoke(ctx, "/like.like/GetTopK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) AddScore(ctx context.Context, in *AddScoreReq, opts ...grpc.CallOption) (*AddScoreResp, error) {
	out := new(AddScoreResp)
	err := c.cc.Invoke(ctx, "/like.like/AddScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) DailyUpdate(ctx context.Context, in *DailyUpdateReq, opts ...grpc.CallOption) (*DailyUpdateResp, error) {
	out := new(DailyUpdateResp)
	err := c.cc.Invoke(ctx, "/like.like/DailyUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServer is the server API for Like service.
// All implementations must embed UnimplementedLikeServer
// for forward compatibility
type LikeServer interface {
	// 点赞/取消赞
	DoLike(context.Context, *DoLikeReq) (*DoLikeResp, error)
	// 获取用户是否点赞
	GetUserLike(context.Context, *GetUserLikedReq) (*GetUserLikedResp, error)
	// 获取目标点赞数
	GetTargetLikes(context.Context, *GetTargetLikesReq) (*GetTargetLikesResp, error)
	GetTopK(context.Context, *GetTopKReq) (*GetTopKResp, error)
	AddScore(context.Context, *AddScoreReq) (*AddScoreResp, error)
	DailyUpdate(context.Context, *DailyUpdateReq) (*DailyUpdateResp, error)
	mustEmbedUnimplementedLikeServer()
}

// UnimplementedLikeServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServer struct {
}

func (UnimplementedLikeServer) DoLike(context.Context, *DoLikeReq) (*DoLikeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLike not implemented")
}
func (UnimplementedLikeServer) GetUserLike(context.Context, *GetUserLikedReq) (*GetUserLikedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLike not implemented")
}
func (UnimplementedLikeServer) GetTargetLikes(context.Context, *GetTargetLikesReq) (*GetTargetLikesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetLikes not implemented")
}
func (UnimplementedLikeServer) GetTopK(context.Context, *GetTopKReq) (*GetTopKResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopK not implemented")
}
func (UnimplementedLikeServer) AddScore(context.Context, *AddScoreReq) (*AddScoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddScore not implemented")
}
func (UnimplementedLikeServer) DailyUpdate(context.Context, *DailyUpdateReq) (*DailyUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DailyUpdate not implemented")
}
func (UnimplementedLikeServer) mustEmbedUnimplementedLikeServer() {}

// UnsafeLikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServer will
// result in compilation errors.
type UnsafeLikeServer interface {
	mustEmbedUnimplementedLikeServer()
}

func RegisterLikeServer(s grpc.ServiceRegistrar, srv LikeServer) {
	s.RegisterService(&Like_ServiceDesc, srv)
}

func _Like_DoLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).DoLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.like/DoLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).DoLike(ctx, req.(*DoLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_GetUserLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLikedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetUserLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.like/GetUserLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetUserLike(ctx, req.(*GetUserLikedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_GetTargetLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetLikesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetTargetLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.like/GetTargetLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetTargetLikes(ctx, req.(*GetTargetLikesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_GetTopK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopKReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetTopK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.like/GetTopK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetTopK(ctx, req.(*GetTopKReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_AddScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddScoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).AddScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.like/AddScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).AddScore(ctx, req.(*AddScoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_DailyUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DailyUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).DailyUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.like/DailyUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).DailyUpdate(ctx, req.(*DailyUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Like_ServiceDesc is the grpc.ServiceDesc for Like service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Like_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "like.like",
	HandlerType: (*LikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoLike",
			Handler:    _Like_DoLike_Handler,
		},
		{
			MethodName: "GetUserLike",
			Handler:    _Like_GetUserLike_Handler,
		},
		{
			MethodName: "GetTargetLikes",
			Handler:    _Like_GetTargetLikes_Handler,
		},
		{
			MethodName: "GetTopK",
			Handler:    _Like_GetTopK_Handler,
		},
		{
			MethodName: "AddScore",
			Handler:    _Like_AddScore_Handler,
		},
		{
			MethodName: "DailyUpdate",
			Handler:    _Like_DailyUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like.proto",
}