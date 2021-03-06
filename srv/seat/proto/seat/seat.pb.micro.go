// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/seat/seat.proto

package seat

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

// Api Endpoints for Seat service

func NewSeatEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Seat service

type SeatService interface {
	CountRemainingSeats(ctx context.Context, in *CountRemainingSeatsRequest, opts ...client.CallOption) (*CountRemainingSeatsReply, error)
	GetSeats(ctx context.Context, in *GetSeatsRequest, opts ...client.CallOption) (*GetSeatsReply, error)
	RollbackSeat(ctx context.Context, in *RollbackSeatRequest, opts ...client.CallOption) (*RollbackSeatReply, error)
}

type seatService struct {
	c    client.Client
	name string
}

func NewSeatService(name string, c client.Client) SeatService {
	return &seatService{
		c:    c,
		name: name,
	}
}

func (c *seatService) CountRemainingSeats(ctx context.Context, in *CountRemainingSeatsRequest, opts ...client.CallOption) (*CountRemainingSeatsReply, error) {
	req := c.c.NewRequest(c.name, "Seat.CountRemainingSeats", in)
	out := new(CountRemainingSeatsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seatService) GetSeats(ctx context.Context, in *GetSeatsRequest, opts ...client.CallOption) (*GetSeatsReply, error) {
	req := c.c.NewRequest(c.name, "Seat.GetSeats", in)
	out := new(GetSeatsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seatService) RollbackSeat(ctx context.Context, in *RollbackSeatRequest, opts ...client.CallOption) (*RollbackSeatReply, error) {
	req := c.c.NewRequest(c.name, "Seat.RollbackSeat", in)
	out := new(RollbackSeatReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Seat service

type SeatHandler interface {
	CountRemainingSeats(context.Context, *CountRemainingSeatsRequest, *CountRemainingSeatsReply) error
	GetSeats(context.Context, *GetSeatsRequest, *GetSeatsReply) error
	RollbackSeat(context.Context, *RollbackSeatRequest, *RollbackSeatReply) error
}

func RegisterSeatHandler(s server.Server, hdlr SeatHandler, opts ...server.HandlerOption) error {
	type seat interface {
		CountRemainingSeats(ctx context.Context, in *CountRemainingSeatsRequest, out *CountRemainingSeatsReply) error
		GetSeats(ctx context.Context, in *GetSeatsRequest, out *GetSeatsReply) error
		RollbackSeat(ctx context.Context, in *RollbackSeatRequest, out *RollbackSeatReply) error
	}
	type Seat struct {
		seat
	}
	h := &seatHandler{hdlr}
	return s.Handle(s.NewHandler(&Seat{h}, opts...))
}

type seatHandler struct {
	SeatHandler
}

func (h *seatHandler) CountRemainingSeats(ctx context.Context, in *CountRemainingSeatsRequest, out *CountRemainingSeatsReply) error {
	return h.SeatHandler.CountRemainingSeats(ctx, in, out)
}

func (h *seatHandler) GetSeats(ctx context.Context, in *GetSeatsRequest, out *GetSeatsReply) error {
	return h.SeatHandler.GetSeats(ctx, in, out)
}

func (h *seatHandler) RollbackSeat(ctx context.Context, in *RollbackSeatRequest, out *RollbackSeatReply) error {
	return h.SeatHandler.RollbackSeat(ctx, in, out)
}
