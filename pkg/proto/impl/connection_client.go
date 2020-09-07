package impl

import (
	"context"
	"fmt"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/sys"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

)

func ConnectClient(tls bool, caFile string, serverAddr string, serverHostOverride string) (*proto.ConnectionAck,error) {

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
	color.Yellow(fmt.Sprintf("Connecting to %s",serverAddr))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewNodeClient(conn)

	return client.Register(context.Background(),&proto.ConnectionSyn{
		Node: sys.GetLocalNode(),

	})
}