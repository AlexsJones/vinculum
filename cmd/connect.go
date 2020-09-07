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
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"os"
)

var (
	tls                bool
	caFile             string
	serverAddr         string
	serverHostOverride string
)

func commandListener(c chan bool) {
	lis, err := net.Listen("tcp", config.DefaultGRPCommandListeningAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	commsServer := impl.CommandServerImpl{}
	grpcServer := grpc.NewServer()

	proto.RegisterCommandServer(grpcServer, commsServer)

	color.Blue(fmt.Sprintf("Starting GRPC server for commands on %s", config.DefaultGRPCommandListeningAddr))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)

	}

	c <- true

	defer grpcServer.Stop()
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to another vinculum",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := make(chan bool)

		color.Blue("Starting command listener")

		go commandListener(ctx)

		color.Blue("Starting connection client")

		if _, err := impl.ConnectClient(tls, caFile, serverAddr, serverHostOverride); err != nil {
			log.Fatal(err)
		}

		for
		{
			select {
			case msg := <-ctx:
				if msg {
					os.Exit(0)
				}
			}
		}
	},
}

func init() {

	connectCmd.Flags().BoolVarP(&tls, "tls", "t", false, "Connection uses TLS if true, else plain TCP")
	connectCmd.Flags().StringVarP(&caFile, "cafile", "c", "", "The file containing the CA cert file")
	connectCmd.Flags().StringVarP(&serverAddr, "serverAddr", "s", "", "The server address in the format of host:port")
	connectCmd.Flags().StringVarP(&serverHostOverride, "serverHostOverride", "o", "", "The server name used to verify the hostname returned by the TLS handshake")
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
