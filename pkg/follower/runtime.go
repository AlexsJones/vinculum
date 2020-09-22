package follower

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
)

func Runtime(syn *proto.SyncSyn) (*proto.SyncAck, error) {

	returnPayload := &proto.SyncAck{
		Error:       "",
		Response:    "OK",
	}

	return returnPayload, nil
}
