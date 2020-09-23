package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AlexsJones/vinculum/pkg/config"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	log "github.com/sirupsen/logrus"
)

type CommandServerImpl struct{}

func (c CommandServerImpl) Send(input *proto.Input, server proto.CTL_SendServer) error {

	if input.Command == "" {
		return errors.New("no command specified")
	}
	switch input.CommandType {
	case proto.CommandType_Shell:
		log.Printf("Leader requested to federate remote shell request: %v", input.Command)
		for _, node := range tracker.Instance().Nodes {

			if node.State == proto.NodeConfig_Unresponsive {
				log.Debugf("Ignoring unresponsive node %s@%s", node.Guid, node.IpAddr)
				continue
			}

			serverAdd := fmt.Sprintf("%s%s", node.IpAddr, config.DefaultGRPSyncListeningAddr)
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(time.Second*2)))

			defer cancel()

			update := proto.Update{}

			// Using the local node as server address rather than any default settings
			// ServerHostOverride might be a problem later
			ack, err := SendCommand(config.Tls, config.CaFile,
				serverAdd, config.ServerHostOverride,
				ctx, &proto.CommandInput{
					Command:     input.Command,
					CommandType: input.CommandType,
				})

			if err != nil {
				// Monitor error codes from connections
				if err.Error() == "context deadline exceeded" {
					node.State = proto.NodeConfig_Unresponsive
				} else {
					log.Warn(err)
					update.Error = err.Error()
					continue
				}
			}
			update.Error = ack.Error
			update.Response = ack.Response
			server.Send(&update)
		}
		break

	}
	return nil
}
