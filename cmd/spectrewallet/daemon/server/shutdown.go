package server

import (
	"context"

	"github.com/spectre-project/spectred/cmd/spectrewallet/daemon/pb"
)

func (s *server) Shutdown(ctx context.Context, request *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	close(s.shutdown)
	return &pb.ShutdownResponse{}, nil
}
