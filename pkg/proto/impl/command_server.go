package impl

import (
	"errors"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	log "github.com/sirupsen/logrus"
)

type CommandServerImpl struct{}

func (c CommandServerImpl) Send(input *proto.Input, server proto.Command_SendServer) error {

	if input.Command == "" {
		return errors.New("no command specified")
	}
	switch input.CommandType {
	case proto.CommandType_CtlShell:
		log.Printf("Recieved incoming command: %v", input.Command)
		for _, node := range tracker.Instance().Nodes {

			if node.State == proto.NodeConfig_Unresponsive {
				log.Debugf("Ignoring unresponsive node %s@%s", node.Guid, node.IpAddr)
				continue
			}

		//	serverAdd := fmt.Sprintf("%s%s", node.IpAddr, config.DefaultGRPSyncListeningAddr)
			//ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(time.Second*2)))

			update := proto.Update{}

			//err := impl.SendSyncCommand(ctx)
			//
			//if err != nil {
			//	// Monitor error codes from connections
			//	if err.Error() == "context deadline exceeded" {
			//		node.State = proto.NodeConfig_Unresponsive
			//	} else {
			//		log.Warn(err)
			//		update.Error = err.Error()
			//		continue
			//	}
			//}

			server.Send(&update)
		}

		break

	}
	return nil
}
