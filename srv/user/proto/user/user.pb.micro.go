// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginReply, error)
	QueryUserInformation(ctx context.Context, in *QueryUserInformationRequest, opts ...client.CallOption) (*QueryUserInformationReply, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordReply, error)
	AddRegularPassenger(ctx context.Context, in *AddRegularPassengerRequest, opts ...client.CallOption) (*AddRegularPassengerReply, error)
	QueryRegularPassengers(ctx context.Context, in *QueryRegularPassengersRequest, opts ...client.CallOption) (*QueryRegularPassengersReply, error)
	UpdateRegularPassenger(ctx context.Context, in *UpdateRegularPassengerRequest, opts ...client.CallOption) (*UpdateRegularPassengerReply, error)
	DeleteRegularPassenger(ctx context.Context, in *DeleteRegularPassengerRequest, opts ...client.CallOption) (*DeleteRegularPassengerReply, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterReply, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(RegisterReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginReply, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) QueryUserInformation(ctx context.Context, in *QueryUserInformationRequest, opts ...client.CallOption) (*QueryUserInformationReply, error) {
	req := c.c.NewRequest(c.name, "User.QueryUserInformation", in)
	out := new(QueryUserInformationReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordReply, error) {
	req := c.c.NewRequest(c.name, "User.UpdatePassword", in)
	out := new(UpdatePasswordReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AddRegularPassenger(ctx context.Context, in *AddRegularPassengerRequest, opts ...client.CallOption) (*AddRegularPassengerReply, error) {
	req := c.c.NewRequest(c.name, "User.AddRegularPassenger", in)
	out := new(AddRegularPassengerReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) QueryRegularPassengers(ctx context.Context, in *QueryRegularPassengersRequest, opts ...client.CallOption) (*QueryRegularPassengersReply, error) {
	req := c.c.NewRequest(c.name, "User.QueryRegularPassengers", in)
	out := new(QueryRegularPassengersReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateRegularPassenger(ctx context.Context, in *UpdateRegularPassengerRequest, opts ...client.CallOption) (*UpdateRegularPassengerReply, error) {
	req := c.c.NewRequest(c.name, "User.UpdateRegularPassenger", in)
	out := new(UpdateRegularPassengerReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteRegularPassenger(ctx context.Context, in *DeleteRegularPassengerRequest, opts ...client.CallOption) (*DeleteRegularPassengerReply, error) {
	req := c.c.NewRequest(c.name, "User.DeleteRegularPassenger", in)
	out := new(DeleteRegularPassengerReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Register(context.Context, *RegisterRequest, *RegisterReply) error
	Login(context.Context, *LoginRequest, *LoginReply) error
	QueryUserInformation(context.Context, *QueryUserInformationRequest, *QueryUserInformationReply) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *UpdatePasswordReply) error
	AddRegularPassenger(context.Context, *AddRegularPassengerRequest, *AddRegularPassengerReply) error
	QueryRegularPassengers(context.Context, *QueryRegularPassengersRequest, *QueryRegularPassengersReply) error
	UpdateRegularPassenger(context.Context, *UpdateRegularPassengerRequest, *UpdateRegularPassengerReply) error
	DeleteRegularPassenger(context.Context, *DeleteRegularPassengerRequest, *DeleteRegularPassengerReply) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *RegisterRequest, out *RegisterReply) error
		Login(ctx context.Context, in *LoginRequest, out *LoginReply) error
		QueryUserInformation(ctx context.Context, in *QueryUserInformationRequest, out *QueryUserInformationReply) error
		UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordReply) error
		AddRegularPassenger(ctx context.Context, in *AddRegularPassengerRequest, out *AddRegularPassengerReply) error
		QueryRegularPassengers(ctx context.Context, in *QueryRegularPassengersRequest, out *QueryRegularPassengersReply) error
		UpdateRegularPassenger(ctx context.Context, in *UpdateRegularPassengerRequest, out *UpdateRegularPassengerReply) error
		DeleteRegularPassenger(ctx context.Context, in *DeleteRegularPassengerRequest, out *DeleteRegularPassengerReply) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterReply) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *LoginReply) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) QueryUserInformation(ctx context.Context, in *QueryUserInformationRequest, out *QueryUserInformationReply) error {
	return h.UserHandler.QueryUserInformation(ctx, in, out)
}

func (h *userHandler) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordReply) error {
	return h.UserHandler.UpdatePassword(ctx, in, out)
}

func (h *userHandler) AddRegularPassenger(ctx context.Context, in *AddRegularPassengerRequest, out *AddRegularPassengerReply) error {
	return h.UserHandler.AddRegularPassenger(ctx, in, out)
}

func (h *userHandler) QueryRegularPassengers(ctx context.Context, in *QueryRegularPassengersRequest, out *QueryRegularPassengersReply) error {
	return h.UserHandler.QueryRegularPassengers(ctx, in, out)
}

func (h *userHandler) UpdateRegularPassenger(ctx context.Context, in *UpdateRegularPassengerRequest, out *UpdateRegularPassengerReply) error {
	return h.UserHandler.UpdateRegularPassenger(ctx, in, out)
}

func (h *userHandler) DeleteRegularPassenger(ctx context.Context, in *DeleteRegularPassengerRequest, out *DeleteRegularPassengerReply) error {
	return h.UserHandler.DeleteRegularPassenger(ctx, in, out)
}
