package grpc

import (
	"context"

	pb "github.com/10xdev4u-alt/VaultMesh/pkg/api/grpc/proto"
	"google.golang.org/grpc"
)

// Server implements the gRPC service.
type Server struct {
	pb.UnimplementedVaultMeshServiceServer
}

// NewServer creates a new gRPC server.
func NewServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterVaultMeshServiceServer(s, &Server{})
	return s
}

// GetNodeInfo returns current node status.
func (s *Server) GetNodeInfo(ctx context.Context, req *pb.NodeInfoRequest) (*pb.NodeInfoResponse, error) {
	return &pb.NodeInfoResponse{
		PeerId:         "local-node",
		ConnectedPeers: 0,
	}, nil
}
