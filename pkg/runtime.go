package pkg

import (
	"fmt"
	"github.com/AlexsJones/vinculum/pkg/config"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/proto/impl"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	"github.com/AlexsJones/vinculum/pkg/types"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

func RuntimeStart(tls bool, caFile string, serverHostOverride string) error {
	r := rand.Intn(5)
	// Listening server started, the main command loop will send commands
	for {
		if tracker.Instance().Count() < 1 {

			time.Sleep(time.Duration(r) * time.Second)
			continue
		}
		// Node heartbeat
		for _, node := range tracker.Instance().Nodes {
			serverAdd := fmt.Sprintf("%s%s", node.IpAddr, config.DefaultGRPCommandListeningAddr)
			color.Yellow("Sending command to %s@%s", node.Guid, serverAdd)

			impl.ConnectCommand(tls, caFile, serverAdd, serverHostOverride, []types.Command{
				{CommandType: proto.CommandName_HealthCheck,
					Args: ""},
			})
			color.Blue("Sent")
		}

		time.Sleep(time.Duration(r) * time.Second)
	}
}
