package impl

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"

)

type ConnectionImpl struct{}

// ConnectionStream is a bidirectional listener/client stream
func (ConnectionImpl) ConnectionStream(server proto.Connection_ConnectionStreamServer) error {
	for {
		updateSyn, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Debugf("Received UpdateSyn from %s",updateSyn.SenderGuid)

		if err := server.Send(&proto.ConnectionAck{
			ResponderGuid: viper.GetString("vinculum-guid"),
		}); err != nil {
				return err
		}

	}
}

