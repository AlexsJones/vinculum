package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)


type SyncServerImpl struct{}
//Command will run incoming commands from the leader
func (i SyncServerImpl) Command(ctx context.Context, input *proto.CommandInput) (*proto.CommandOutput, error) {

	log.Debug("Ran a commmand")

	return &proto.CommandOutput{},nil
}

func (SyncServerImpl) HealthCheck(ctx context.Context, syn *proto.HealthCheckSyn)(*proto.HealthCheckAck, error) {

	peer,ok := peer.FromContext(ctx)
	if ok {
		log.Debugf("Received incoming health check from %s",peer.Addr.String())
	}else {
		log.Debugf("Received incoming health check from an unknown leader")
	}
	return &proto.HealthCheckAck{
		Response: "OK",
		Error:    "",
	},nil
}
