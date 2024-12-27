package proto

import (
	"context"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/env"
	"github.com/terminalnode/adventofcode2024/common/util"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	UnimplementedAdventServiceServer
	solvePart1 util.Solution
	solvePart2 util.Solution
}

func CreateGRPCServer(
	day int,
	part1 util.Solution,
	part2 util.Solution,
) *grpc.Server {
	port := env.GetStringOrDefault(env.GrpcPort, "50051")
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterAdventServiceServer(grpcServer, &server{solvePart1: part1, solvePart2: part2})
	log.Printf("gRPC server for day #%d starting on port %s", day, port)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Fatal gRPC server error: %v", err)
		}
	}()
	return grpcServer
}

// This part of the architecture is a bit #YOLO
// But at least I got to play around a little with gRPC
// Throughout the entire process I've only used the Kubernetes solution over HTTP

func (s *server) SolvePart1(
	ctx context.Context,
	req *InputData,
) (*InputResponse, error) {
	result, err := s.solvePart1(util.AocInput{Input: req.Input})
	if s := result.Solution; s != "" {
		return &InputResponse{Result: s}, nil
	} else {
		return &InputResponse{Result: err.Error()}, nil
	}
}

func (s *server) SolvePart2(
	ctx context.Context,
	req *InputData,
) (*InputResponse, error) {
	result, err := s.solvePart2(util.AocInput{Input: req.Input})
	if s := result.Solution; s != "" {
		return &InputResponse{Result: s}, nil
	} else {
		return &InputResponse{Result: err.Error()}, nil
	}
}
