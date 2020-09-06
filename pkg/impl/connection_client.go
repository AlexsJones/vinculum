package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	log "github.com/sirupsen/logrus"
)

func Connect(tls bool, caFile string, serverAddr string, serverHostOverride string) {

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

	client := proto.NewConnectionClient(conn)

	stream, err := client.ConnectionStream(context.Background())

	waitc := make(chan struct{})
	go func() {
		for {
			updateAck, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive an UpdateSyn : %v", err)
			}
			log.Debugf("Received UpdateSyn from %s",updateAck.ResponderGuid)

		}
	}()

	if err := stream.Send(&proto.ConnectionSyn{
		SenderGuid: viper.GetString("vinculum-guid"),
	}); err != nil {
		log.Fatalf("Failed to send UpdateSyn: %v", err)
	}

	stream.CloseSend()
	<-waitc
}