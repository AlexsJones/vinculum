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
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	log "github.com/sirupsen/logrus"
	"io"
)

var (
	tls bool
	caFile string
	serverAddr string
	serverHostOverride string
)

func connect() {

	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			log.Fatal("caFile is missing")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewConnectionClient(conn)

	stream, err := client.ConnectionStream(context.Background())

	waitc := make(chan struct{})
	go func() {
		for {
			updateAck, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive an UpdateSyn : %v", err)
			}
			log.Debugf("Received UpdateSyn from %s",updateAck.ResponderGuid)

		}
	}()

	if err := stream.Send(&proto.ConnectionSyn{
		SenderGuid: viper.GetString("vinculum-guid"),
	}); err != nil {
		log.Fatalf("Failed to send UpdateSyn: %v", err)
	}

	stream.CloseSend()
	<-waitc
}
// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to another vinculum",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		connect()
	},
}

func init() {

	connectCmd.Flags().BoolVarP(&tls,"tls","t",false,"Connection uses TLS if true, else plain TCP")
	connectCmd.Flags().StringVarP(&caFile,"cafile","c","","The file containing the CA cert file")
	connectCmd.Flags().StringVarP(&serverAddr,"serverAddr","s","","The server address in the format of host:port")
	connectCmd.Flags().StringVarP(&serverHostOverride,"serverHostOverride","o","","The server name used to verify the hostname returned by the TLS handshake")
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
