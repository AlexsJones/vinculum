package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
)


type CommandServerImpl struct{}

func (CommandServerImpl) Send(ctx context.Context, syn *proto.CommandSyn) (*proto.CommandAck, error) {

	log.Debugf("Received incoming command %s",syn.CommandName.String())


	return &proto.CommandAck{
		Error: nil,
		CommandName: syn.CommandName,
	}, nil
}
