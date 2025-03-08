package adapter

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	core "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/service"
	pb "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/proto"
	"google.golang.org/grpc"
)

type GrpcController struct {
	calendarService core.CalendarService
	pb.UnimplementedCalendarManagerServer
}

func NewGrpcController(calendarService core.CalendarService) *GrpcController {
	return &GrpcController{calendarService: calendarService}
}

func (c *GrpcController) Run() error {
	port := os.Getenv("PORT_GRPC_CALENDAR_MANAGER")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterCalendarManagerServer(s, c)

	fmt.Println("Starting GRPC server on port " + port)
	return s.Serve(listener)
}

func (c *GrpcController) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.EventResponse, error) {
	input := core.CreateEventInput{
		Name:        req.Name,
		Description: req.Description,
		Time:        req.Time,
		Duration:    req.Duration,
		Attendees:   req.Attendees,
	}
	event, err := c.calendarService.CreateEvent(input)
	if err != nil {
		return &pb.EventResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.EventResponse{
		Success:     true,
		Message:     "Event created successfully",
		Id:          uint32(event.Id),
		Name:        event.Name,
		Description: event.Description,
		Time:        event.Time.Format(time.RFC3339),
		Duration:    event.Duration,
		Attendees:   event.Attendees,
		CreatedAt:   event.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   event.UpdatedAt.Format(time.RFC3339),
		DeletedAt:   event.DeletedAt.Format(time.RFC3339),
	}, nil
}

func (c *GrpcController) ReadEvent(ctx context.Context, req *pb.ReadEventRequest) (*pb.EventResponse, error) {
	input := core.ReadEventInput{Id: uint32(req.Id)}
	event, err := c.calendarService.ReadEvent(input)
	if err != nil {
		return &pb.EventResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.EventResponse{
		Success:     true,
		Message:     "Event read successfully",
		Id:          uint32(event.Id),
		Name:        event.Name,
		Description: event.Description,
		Time:        event.Time.Format(time.RFC3339),
		Duration:    event.Duration,
		Attendees:   event.Attendees,
		CreatedAt:   event.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   event.UpdatedAt.Format(time.RFC3339),
		DeletedAt:   event.DeletedAt.Format(time.RFC3339),
	}, nil
}

func (c *GrpcController) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.EventResponse, error) {
	input := core.UpdateEventInput{
		Id:          uint32(req.Id),
		Name:        req.Name,
		Description: req.Description,
		Time:        req.Time,
		Duration:    req.Duration,
		Attendees:   req.Attendees,
	}
	event, err := c.calendarService.UpdateEvent(input)
	if err != nil {
		return &pb.EventResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.EventResponse{
		Success:     true,
		Message:     "Event updated successfully",
		Id:          uint32(event.Id),
		Name:        event.Name,
		Description: event.Description,
		Time:        event.Time.Format(time.RFC3339),
		Duration:    event.Duration,
		Attendees:   event.Attendees,
		CreatedAt:   event.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   event.UpdatedAt.Format(time.RFC3339),
		DeletedAt:   event.DeletedAt.Format(time.RFC3339),
	}, nil
}

func (c *GrpcController) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	input := core.DeleteEventInput{Id: uint32(req.Id)}
	err := c.calendarService.DeleteEvent(input)
	if err != nil {
		return &pb.DeleteEventResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DeleteEventResponse{
		Success: true,
		Message: "Event deleted successfully",
	}, nil
}
