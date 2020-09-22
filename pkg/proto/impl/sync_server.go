package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/follower"
	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)


type SyncServerImpl struct{}

func (SyncServerImpl) Send(ctx context.Context, syn *proto.SyncSyn) (*proto.SyncAck, error) {

	peer,ok := peer.FromContext(ctx)
	if ok {
		log.Debugf("Received incoming command %s from %s", syn.CommandName.String(),peer.Addr.String())
	}else {
		log.Debugf("Received incoming command %s", syn.CommandName.String())
	}
	return follower.Runtime(syn)
}
