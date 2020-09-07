/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/AlexsJones/vinculum/pkg/config"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/proto/impl"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"

	"github.com/spf13/cobra"
)

func listen() {

	lis, err := net.Listen("tcp", config.DefaultGRPCConnectListeningAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	commsServer := impl.ConnectionServerImpl{}
	grpcServer := grpc.NewServer()

	proto.RegisterNodeServer(grpcServer, commsServer)

	color.Blue(fmt.Sprintf("Starting GRPC server for registering on %s", config.DefaultGRPCConnectListeningAddr))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)

	}
	defer grpcServer.Stop()

}

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start the listener server",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := rand.Intn(5)
		go listen()

		// Listening server started, the main command loop will send commands
		for {
			if tracker.Instance().Count() < 1 {

				time.Sleep(time.Duration(r) * time.Second)
				continue
			}

			for _, node := range tracker.Instance().Nodes {
				serverAdd := fmt.Sprintf("%s%s",node.IpAddr,config.DefaultGRPCommandListeningAddr)
				color.Yellow("Sending command to %s@%s",node.Guid,serverAdd)
				impl.ConnectCommand(tls,caFile,serverAdd,serverHostOverride)
				color.Blue("Sent")
			}

			time.Sleep(time.Duration(r) * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
