package types

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
)

type Command struct {
	CommandType proto.CommandName
	Args        string
	Context context.Context
}