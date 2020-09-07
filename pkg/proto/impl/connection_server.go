package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/sys"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
	"strings"
)

type ConnectionServerImpl struct{}

func (ConnectionServerImpl) Register(ctx context.Context, syn *proto.ConnectionSyn) (*proto.ConnectionAck, error) {

	peer, bool := peer.FromContext(ctx)
	if !bool {
		return &proto.ConnectionAck{
			Node:  sys.GetLocalNode(),
			State: proto.ConnectionAck_Negotiating,
		}, nil
	}

	spl := strings.Split(peer.Addr.String(), ":")

	syn.Node.IpAddr = spl[0]
	syn.Node.Network = peer.Addr.Network()

	if tracker.Instance().Get(syn.Node.Guid) == nil {
		tracker.Instance().Add(syn.Node)
		log.Debugf("Added node %s: %v", syn.Node.Guid, syn.Node)
	} else {
		log.Debugf("Received UpdateSyn from %s", syn.Node.Guid)
	}

	return &proto.ConnectionAck{
		Node:  sys.GetLocalNode(),
		State: proto.ConnectionAck_Known,
	}, nil
}
