package sys

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/spf13/viper"
)

func GetLocalNode() *proto.NodeConfig {

	node := &proto.NodeConfig{
		Guid: viper.GetString("vinculum-guid"),
	}

	return node
}
