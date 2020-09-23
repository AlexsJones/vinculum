package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/sys"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NodeClient(tls bool, caFile string, serverAddr string,
	serverHostOverride string) (*proto.ConnectionAck, error) {

	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			return nil, errors.New("caFile is missing")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			return nil, fmt.Errorf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	color.Yellow(fmt.Sprintf("NodeClient: Connecting to %s", serverAddr))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, fmt.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewNodeClient(conn)

	return client.Register(context.Background(), &proto.ConnectionSyn{
		Node: sys.GetLocalNode(),
	})
}
