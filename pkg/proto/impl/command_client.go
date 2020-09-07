package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	log "github.com/sirupsen/logrus"
)

func ConnectCommand(tls bool, caFile string, serverAddr string, serverHostOverride string) {

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

	if err := stream.Send(&proto.CommandSyn{
		KeepAlive: 1,
	}); err != nil {
		log.Fatalf("Failed to send UpdateSyn: %v", err)
	}

	stream.CloseSend()
	<-waitc
}