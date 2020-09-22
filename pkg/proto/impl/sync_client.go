package impl

import (
	"errors"
	"fmt"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func SendCommand(tls bool, caFile string, serverAddr string,
	serverHostOverride string, command types.SyncCommand) error {

	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			return errors.New("caFile is missing")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to create TLS credentials %v",err))
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.DialContext(command.Context,serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewSyncClient(conn)

	commandAck, err := client.Send(command.Context,&proto.SyncSyn{
		CommandName: command.CommandType,
		CommandArgs: command.Args,
	})
	if err != nil {
		return err
	}
	if commandAck.Error != "" {
		log.Warnf("%s sent error %s for command %s",serverAddr,commandAck.Error,commandAck.CommandName)
	}
	return nil
}
