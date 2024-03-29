// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"path/filepath"
	"strings"

	"github.com/asciifaceman/drawdio/pkg/drawer"
	"github.com/asciifaceman/drawdio/pkg/intake"
	"github.com/asciifaceman/drawdio/pkg/logging"
	"github.com/spf13/cobra"
)

// pixelateCmd represents the pixelate command
var pixelateCmd = &cobra.Command{
	Use:   "pixelate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.Logger()
		i, err := intake.New(&intake.Config{
			Filename: filename,
			Logger:   logger,
		})
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		set, err := i.Sample()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		decomposed := set.Decompose()
		base := filepath.Base(filename)
		drawer.Sound(strings.TrimSuffix(base, filepath.Ext(base)), decomposed)
	},
}

func init() {
	rootCmd.AddCommand(pixelateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pixelateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pixelateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
