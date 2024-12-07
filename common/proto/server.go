package proto

import (
	"context"
	"github.com/terminalnode/adventofcode2024/common"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	UnimplementedAdventServiceServer
	solvePart1 common.Solution
	solvePart2 common.Solution
}

func StartGRPCServer(
	port string,
	part1 common.Solution,
	part2 common.Solution,
) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterAdventServiceServer(grpcServer, &server{solvePart1: part1, solvePart2: part2})
	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SolvePart1(
	ctx context.Context,
	req *InputData,
) (*InputResponse, error) {
	result := s.solvePart1(req.Input)
	return &InputResponse{Result: result}, nil
}

func (s *server) SolvePart2(
	ctx context.Context,
	req *InputData,
) (*InputResponse, error) {
	result := s.solvePart2(req.Input)
	return &InputResponse{Result: result}, nil
}
