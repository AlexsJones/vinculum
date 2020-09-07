package impl

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
	"io"
)

type CommandServerImpl struct{}

//Stream command is run on the vinculum client to process and acknowledge commands
func (CommandServerImpl) Stream(server proto.Command_StreamServer) error {

	for {
		inComingCmd, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Debugf("Received incoming command %s",inComingCmd.CommandName.String())


		if err := server.Send(&proto.CommandAck{

		}); err != nil {
			log.Warn(err)
		}

	}
}
