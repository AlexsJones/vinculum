package impl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//SendCTLCommand inputs a CLI based control command
func SendCTLCommand(tls bool, caFile string, serverAddr string,
	serverHostOverride string, command string, commandType proto.CommandType) error {
	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			return errors.New("caFile is missing")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			return fmt.Errorf("failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	color.Yellow(fmt.Sprintf("CTLClient: Connecting to %s", serverAddr))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return fmt.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewCTLClient(conn)

	sendClient, err := client.Send(context.Background(), &proto.Input{
		Command:     command,
		CommandType: commandType,
	})
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	for {
		response, err := sendClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.CommandClient(_) = _, %v", client, err)
		}
		log.Println(response)
	}

	return nil
}
