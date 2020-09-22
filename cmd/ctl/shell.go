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
package ctl

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var (
	inputCommand string
)
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Run remote shell commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Running command %s", inputCommand)

	},
}

func init() {

	shellCmd.Flags().StringVarP(&inputCommand, "command", "c", "", "Command to send to followers")
	if err := shellCmd.MarkFlagRequired("command"); err != nil {
		log.Fatal(err)
	}
	CtlCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
