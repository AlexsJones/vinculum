package types

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
	"time"
)

type Command struct {
	CommandType proto.CommandName
	Args        string
	MaxDuration time.Duration
}
