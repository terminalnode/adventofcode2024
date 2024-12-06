package proto

import (
	"context"
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
	port string,
	day int,
	part1 util.Solution,
	part2 util.Solution,
) *grpc.Server {
	lis, err := net.Listen("tcp", port)
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
