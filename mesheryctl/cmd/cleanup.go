// Copyright 2019 The Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up Meshery",
	Long:  `Clean up Meshery configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Cleaning Meshery config...")
		if err := downloadFile(dockerComposeFile, fileURL); err != nil {
			log.Fatal("cleanup cmd: ", err)
		}
		log.Info("Meshery config is now cleaned up.")

	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
