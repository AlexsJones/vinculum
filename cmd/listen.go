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
	"github.com/AlexsJones/vinculum/pkg/impl"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	defaultGRPCListeningAddr = ":9671"
)
// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start vinculum in a passive listening mode",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		if addr := viper.GetString("defaultGRPCListeningAddr"); addr != "" {
			defaultGRPCListeningAddr = addr
		}

		lis, err := net.Listen("tcp", defaultGRPCListeningAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		handShakeServer := impl.HandShakeImpl{}
		grpcServer := grpc.NewServer()

		proto.RegisterHandShakeServer(grpcServer,handShakeServer)

		color.Blue(fmt.Sprintf("Starting GRPC server on %s",defaultGRPCListeningAddr))

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
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
