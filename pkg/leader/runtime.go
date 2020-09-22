package leader

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/AlexsJones/vinculum/pkg/config"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/proto/impl"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

func runHealthCheck(tls bool, caFile string, serverHostOverride string,
	parentContext context.Context) error {
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
				log.Debugf("Ignoring unresponsive node %s@%s", node.Guid, node.IpAddr)
				continue
			}

			serverAdd := fmt.Sprintf("%s%s", node.IpAddr, config.DefaultGRPSyncListeningAddr)
			color.Yellow("Sending health check sync to %s@%s", node.Guid, serverAdd)

			ctx, _ := context.WithDeadline(parentContext, time.Now().Add(time.Duration(time.Second*2)))
			err := impl.SendSyncCommand(tls,
				caFile, serverAdd,
				serverHostOverride,
				ctx)

			if err != nil {
				// Monitor error codes from connections
				if err.Error() == "context deadline exceeded" {
					node.State = proto.NodeConfig_Unresponsive
				} else {
					return err
				}
			}
		}
		// ------------------------------------------------------------------------------------------------------------

		time.Sleep(time.Duration(r) * time.Second)
	}
}

//Run ...
func Run(tls bool, caFile string, serverHostOverride string, context context.Context) error {

	return runHealthCheck(tls, caFile, serverHostOverride, context)
}
