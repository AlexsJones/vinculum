package follower

import "github.com/AlexsJones/vinculum/pkg/proto"

func Runtime(syn *proto.SyncSyn) (*proto.SyncAck, error) {

	var err error
	returnPayload := &proto.SyncAck{
		Error:       "",
		Response:    "",
		CommandName: syn.CommandName,
	}

	switch syn.CommandName {
	case proto.CheckName_HealthCheck:
		returnPayload.Response = "OK"
		break
	}

	return returnPayload, err
}
