package types

import "github.com/AlexsJones/vinculum/pkg/proto"

type Command struct {
	CommandType proto.CommandName
	Args        string
}
