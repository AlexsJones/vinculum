package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

func SendCommand(tls bool, caFile string, serverAddr string,
	serverHostOverride string, commands types.Command) {

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

	clientDeadline := time.Now().Add(time.Duration(time.Second * 5))
	ctx, _ := context.WithDeadline(context.Background(), clientDeadline)

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.DialContext(ctx,serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewCommandClient(conn)

	commandAck, err := client.Send(ctx,&proto.CommandSyn{
		CommandName: commands.CommandType,
		CommandArgs: commands.Args,
	})
	if err != nil {
		log.Fatal(err)
	}
	if commandAck.Error != "" {
		log.Warnf("%s sent error %s for command %s",serverAddr,commandAck.Error,commandAck.CommandName)
	}

}
