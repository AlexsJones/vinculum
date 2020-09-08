package follower

import "github.com/AlexsJones/vinculum/pkg/proto"

func Runtime(syn *proto.CommandSyn)(*proto.CommandAck, error ){

	var err error
	returnPayload := &proto.CommandAck{
		Error: "",
		Response: "",
		CommandName: syn.CommandName,
	}

	switch syn.CommandName {
	case proto.CommandName_HealthCheck:
		returnPayload.Response = "OK"
		break
	}



	return returnPayload, err
}
