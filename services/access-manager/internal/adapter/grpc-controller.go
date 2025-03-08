package adapter

import (
	"context"
	"fmt"
	"net"
	"os"

	core "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/service"
	pb "github.com/FelipeStillner/ProjectPilot/services/access-manager/proto"
	"google.golang.org/grpc"
)

type GrpcController struct {
	accessService core.AccessService
	pb.UnimplementedAccessManagerServer
}

func NewGrpcController(accessService core.AccessService) *GrpcController {
	return &GrpcController{accessService: accessService}
}

func (c *GrpcController) Run() error {
	port := os.Getenv("PORT_GRPC_ACCESS_MANAGER")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterAccessManagerServer(s, c)

	fmt.Println("Starting GRPC server on port " + port)
	return s.Serve(listener)
}

func (c *GrpcController) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.TeamResponse, error) {
	input := core.CreateTeamInput{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
	output, err := c.accessService.CreateTeam(input)
	if err != nil {
		return &pb.TeamResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.TeamResponse{
		Success: true,
		Message: "Team created successfully",
		Id:      uint32(output.Id),
		Name:    output.Name,
	}, nil
}

func (c *GrpcController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	input := core.CreateUserInput{
		Username: req.Username,
		Password: req.Password,
		TeamId:   req.TeamId,
	}
	output, err := c.accessService.CreateUser(input)
	if err != nil {
		return &pb.UserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UserResponse{
		Success:  true,
		Message:  "User created successfully",
		Id:       uint32(output.Id),
		Username: output.Username,
		TeamId:   uint32(output.TeamId),
	}, nil
}

func (c *GrpcController) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	input := core.LoginInput{
		Username: req.Username,
		Password: req.Password,
	}

	output, err := c.accessService.Login(input)
	if err != nil {
		return &pb.LoginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Success: true,
		Message: "Login successful",
		UserId:  *output,
	}, nil
}
