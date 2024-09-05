// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user/user.proto

package user

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// users
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*GetProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	ChangeProfileImage(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Void, error)
	FetchUsers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*UserResponses, error)
	ListOfFollowing(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error)
	ListOfFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error)
	ListOfFollowingByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error)
	ListOfFollowersByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error)
	DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	// subscribe
	Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowRes, error)
	Unfollow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*DFollowRes, error)
	GetUserFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error)
	GetUserFollows(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error)
	MostPopularUser(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeProfileImage(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangeProfileImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FetchUsers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*UserResponses, error) {
	out := new(UserResponses)
	err := c.cc.Invoke(ctx, "/user.UserService/FetchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowing(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error) {
	out := new(Followings)
	err := c.cc.Invoke(ctx, "/user.UserService/ListOfFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error) {
	out := new(Followers)
	err := c.cc.Invoke(ctx, "/user.UserService/ListOfFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowingByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error) {
	out := new(Followings)
	err := c.cc.Invoke(ctx, "/user.UserService/ListOfFollowingByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowersByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error) {
	out := new(Followers)
	err := c.cc.Invoke(ctx, "/user.UserService/ListOfFollowersByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowRes, error) {
	out := new(FollowRes)
	err := c.cc.Invoke(ctx, "/user.UserService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Unfollow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*DFollowRes, error) {
	out := new(DFollowRes)
	err := c.cc.Invoke(ctx, "/user.UserService/Unfollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFollows(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserFollows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) MostPopularUser(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/MostPopularUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// users
	Create(context.Context, *CreateRequest) (*UserResponse, error)
	GetProfile(context.Context, *Id) (*GetProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UserResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	ChangeProfileImage(context.Context, *URL) (*Void, error)
	FetchUsers(context.Context, *Filter) (*UserResponses, error)
	ListOfFollowing(context.Context, *Id) (*Followings, error)
	ListOfFollowers(context.Context, *Id) (*Followers, error)
	ListOfFollowingByUsername(context.Context, *Id) (*Followings, error)
	ListOfFollowersByUsername(context.Context, *Id) (*Followers, error)
	DeleteUser(context.Context, *Id) (*Void, error)
	// subscribe
	Follow(context.Context, *FollowReq) (*FollowRes, error)
	Unfollow(context.Context, *FollowReq) (*DFollowRes, error)
	GetUserFollowers(context.Context, *Id) (*Count, error)
	GetUserFollows(context.Context, *Id) (*Count, error)
	MostPopularUser(context.Context, *Void) (*UserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Create(context.Context, *CreateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *Id) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) ChangeProfileImage(context.Context, *URL) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeProfileImage not implemented")
}
func (UnimplementedUserServiceServer) FetchUsers(context.Context, *Filter) (*UserResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUsers not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowing(context.Context, *Id) (*Followings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowing not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowers(context.Context, *Id) (*Followers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowers not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowingByUsername(context.Context, *Id) (*Followings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowingByUsername not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowersByUsername(context.Context, *Id) (*Followers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowersByUsername not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) Follow(context.Context, *FollowReq) (*FollowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedUserServiceServer) Unfollow(context.Context, *FollowReq) (*DFollowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedUserServiceServer) GetUserFollowers(context.Context, *Id) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowers not implemented")
}
func (UnimplementedUserServiceServer) GetUserFollows(context.Context, *Id) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollows not implemented")
}
func (UnimplementedUserServiceServer) MostPopularUser(context.Context, *Void) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostPopularUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangeProfileImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeProfileImage(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FetchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FetchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FetchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FetchUsers(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListOfFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowing(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListOfFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowers(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowingByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowingByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListOfFollowingByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowingByUsername(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowersByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowersByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListOfFollowersByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowersByUsername(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Follow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Unfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Unfollow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFollowers(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFollows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFollows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserFollows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFollows(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_MostPopularUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).MostPopularUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/MostPopularUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).MostPopularUser(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "ChangeProfileImage",
			Handler:    _UserService_ChangeProfileImage_Handler,
		},
		{
			MethodName: "FetchUsers",
			Handler:    _UserService_FetchUsers_Handler,
		},
		{
			MethodName: "ListOfFollowing",
			Handler:    _UserService_ListOfFollowing_Handler,
		},
		{
			MethodName: "ListOfFollowers",
			Handler:    _UserService_ListOfFollowers_Handler,
		},
		{
			MethodName: "ListOfFollowingByUsername",
			Handler:    _UserService_ListOfFollowingByUsername_Handler,
		},
		{
			MethodName: "ListOfFollowersByUsername",
			Handler:    _UserService_ListOfFollowersByUsername_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _UserService_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _UserService_Unfollow_Handler,
		},
		{
			MethodName: "GetUserFollowers",
			Handler:    _UserService_GetUserFollowers_Handler,
		},
		{
			MethodName: "GetUserFollows",
			Handler:    _UserService_GetUserFollows_Handler,
		},
		{
			MethodName: "MostPopularUser",
			Handler:    _UserService_MostPopularUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
