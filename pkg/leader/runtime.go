package leader

import (
	"context"
	"fmt"
	"github.com/AlexsJones/vinculum/pkg/config"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/proto/impl"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	"github.com/AlexsJones/vinculum/pkg/types"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func Run(tls bool, caFile string, serverHostOverride string) error {
	r := rand.Intn(5)
	// Listening leader started, the main command loop will send commands
	for {
		if tracker.Instance().Count() < 1 {

			time.Sleep(time.Duration(r) * time.Second)
			continue
		}
		// Node heartbeat ---------------------------------------------------------------------------------------------
		for _, node := range tracker.Instance().Nodes {

			if node.State == proto.NodeConfig_Unresponsive {
				log.Debugf("Ignoring unresponsive node %s@%s",node.Guid, node.IpAddr)
				continue
			}

			serverAdd := fmt.Sprintf("%s%s", node.IpAddr, config.DefaultGRPCommandListeningAddr)
			color.Yellow("Sending %s command to %s@%s", proto.CommandName_HealthCheck.String(), node.Guid, serverAdd)

			ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(time.Second * 2)))
			err := impl.SendCommand(tls,
				caFile, serverAdd,
				serverHostOverride,
				types.Command{
				CommandType: proto.CommandName_HealthCheck,
					Args: "",
					Context: ctx,
			})

			if err != nil {
				// Monitor error codes from connections
				if err.Error() == "context deadline exceeded" {
					node.State = proto.NodeConfig_Unresponsive
				}else {
					return err
				}
			}
		}
		// ------------------------------------------------------------------------------------------------------------

		time.Sleep(time.Duration(r) * time.Second)
	}
}
