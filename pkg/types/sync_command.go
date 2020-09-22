package types

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
)

type SyncCommand struct {
	CommandType proto.CheckName
	Args        string
	Context context.Context
}
