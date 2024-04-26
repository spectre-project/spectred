package client

import (
	"context"
	"time"

	"github.com/spectre-project/spectred/cmd/spectrewallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/spectre-project/spectred/cmd/spectrewallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the spectrewalletd server, and returns the client instance
func Connect(address string) (pb.SpectrewalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("spectrewallet daemon is not running, start it with `spectrewallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewSpectrewalletdClient(conn), func() {
		conn.Close()
	}, nil
}
