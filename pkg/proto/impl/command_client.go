package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
)

func ConnectCommand(tls bool, caFile string, serverAddr string,
	serverHostOverride string, commands []types.Command) {

	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			log.Fatal("caFile is missing")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewCommandClient(conn)

	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			_, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive an UpdateSyn : %v", err)
			}

		}
	}()
	for _, cmd := range commands {
		if err := stream.Send(&proto.CommandSyn{
			CommandName: cmd.CommandType,
			CommandArgs: cmd.Args,
		}); err != nil {
			log.Fatalf("Failed to send UpdateSyn: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
