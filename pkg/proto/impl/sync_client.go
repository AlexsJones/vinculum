package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func SendSyncCommand(tls bool, caFile string, serverAddr string,
	serverHostOverride string, ctx context.Context) error {

	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			return errors.New("caFile is missing")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to create TLS credentials %v", err))
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewSyncClient(conn)

	commandAck, err := client.Send(ctx, &proto.SyncSyn{
		CommandArgs: "",
	})
	if err != nil {
		return err
	}
	if commandAck.Error != "" {
		log.Warnf("%s sent error %s for health check", serverAddr, commandAck.Error)
	}
	return nil
}
