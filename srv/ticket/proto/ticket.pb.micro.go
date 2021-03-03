// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ticket/ticket.proto

package ticket

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

// Api Endpoints for Ticket service

func NewTicketEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Ticket service

type TicketService interface {
	BookTickets(ctx context.Context, in *BookTicketsRequest, opts ...client.CallOption) (*BookTicketsReply, error)
	RefundTicket(ctx context.Context, in *RefundTicketRequest, opts ...client.CallOption) (*RefundTicketReply, error)
}

type ticketService struct {
	c    client.Client
	name string
}

func NewTicketService(name string, c client.Client) TicketService {
	return &ticketService{
		c:    c,
		name: name,
	}
}

func (c *ticketService) BookTickets(ctx context.Context, in *BookTicketsRequest, opts ...client.CallOption) (*BookTicketsReply, error) {
	req := c.c.NewRequest(c.name, "Ticket.BookTickets", in)
	out := new(BookTicketsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketService) RefundTicket(ctx context.Context, in *RefundTicketRequest, opts ...client.CallOption) (*RefundTicketReply, error) {
	req := c.c.NewRequest(c.name, "Ticket.RefundTicket", in)
	out := new(RefundTicketReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ticket service

type TicketHandler interface {
	BookTickets(context.Context, *BookTicketsRequest, *BookTicketsReply) error
	RefundTicket(context.Context, *RefundTicketRequest, *RefundTicketReply) error
}

func RegisterTicketHandler(s server.Server, hdlr TicketHandler, opts ...server.HandlerOption) error {
	type ticket interface {
		BookTickets(ctx context.Context, in *BookTicketsRequest, out *BookTicketsReply) error
		RefundTicket(ctx context.Context, in *RefundTicketRequest, out *RefundTicketReply) error
	}
	type Ticket struct {
		ticket
	}
	h := &ticketHandler{hdlr}
	return s.Handle(s.NewHandler(&Ticket{h}, opts...))
}

type ticketHandler struct {
	TicketHandler
}

func (h *ticketHandler) BookTickets(ctx context.Context, in *BookTicketsRequest, out *BookTicketsReply) error {
	return h.TicketHandler.BookTickets(ctx, in, out)
}

func (h *ticketHandler) RefundTicket(ctx context.Context, in *RefundTicketRequest, out *RefundTicketReply) error {
	return h.TicketHandler.RefundTicket(ctx, in, out)
}
