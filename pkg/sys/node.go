package sys

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/spf13/viper"
)

func GetLocalNode() *proto.NodeConfig {

	info := getInfo()
	node := &proto.NodeConfig{
		Guid: viper.GetString("vinculum-guid"),
		Hostname: info.Hostname,
		Platform: info.Platform,
		Cpu: info.CPU,
		Ram: info.RAM,
		Disk: info.Disk,
	}
	return node
}