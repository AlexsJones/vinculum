package sys

import (
	"encoding/json"
	"github.com/AlexsJones/vinculum/pkg/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zcalusic/sysinfo"
)

func GetLocalNode() *proto.NodeConfig {

	var si sysinfo.SysInfo

	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	sysi := proto.SystemInfo{}
	if err := json.Unmarshal(data, &sysi); err != nil {
		log.Fatal(err)
	}

	node := &proto.NodeConfig{
		Guid:       viper.GetString("vinculum-guid"),
		SystemInfo: &sysi,
	}

	log.Debug(string(data))

	return node
}
