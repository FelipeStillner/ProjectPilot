package adapter

import (
	"context"
	"fmt"
	"net"
	"os"

	core "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/service"
	pb "github.com/FelipeStillner/ProjectPilot/services/task-manager/proto"
	"google.golang.org/grpc"
)

type GrpcController struct {
	taskService core.TaskService
	pb.UnimplementedTaskManagerServer
}

func NewGrpcController(taskService core.TaskService) *GrpcController {
	return &GrpcController{taskService: taskService}
}

func (c *GrpcController) Run() error {
	port := os.Getenv("PORT_GRPC_TASK_MANAGER")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterTaskManagerServer(s, c)

	fmt.Println("Starting GRPC server on port " + port)
	return s.Serve(listener)
}

func (c *GrpcController) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	input := core.CreateTaskInput{
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Assignee:    req.Assignee,
		Status:      req.Status,
	}
	task, err := c.taskService.CreateTask(input)
	if err != nil {
		return nil, err
	}

	return &pb.TaskResponse{
		Id:          uint32(task.Id),
		Name:        task.Name,
		Description: task.Description,
		Priority:    task.Priority,
		Assignee:    task.Assignee,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (c *GrpcController) ReadTask(ctx context.Context, req *pb.ReadTaskRequest) (*pb.TaskResponse, error) {
	input := core.ReadTaskInput{Id: uint32(req.Id)}
	task, err := c.taskService.ReadTask(input)
	if err != nil {
		return nil, err
	}

	return &pb.TaskResponse{
		Id:          uint32(task.Id),
		Name:        task.Name,
		Description: task.Description,
		Priority:    task.Priority,
		Assignee:    task.Assignee,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (c *GrpcController) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	input := core.UpdateTaskInput{
		Id:          uint32(req.Id),
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Assignee:    req.Assignee,
		Status:      req.Status,
	}
	task, err := c.taskService.UpdateTask(input)
	if err != nil {
		return nil, err
	}

	return &pb.TaskResponse{
		Id:          uint32(task.Id),
		Name:        task.Name,
		Description: task.Description,
		Priority:    task.Priority,
		Assignee:    task.Assignee,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (c *GrpcController) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	input := core.DeleteTaskInput{Id: uint32(req.Id)}
	err := c.taskService.DeleteTask(input)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTaskResponse{Status: "task deleted"}, nil
}
