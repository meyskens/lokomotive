// Copyright 2020 The Lokomotive Authors
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
	"github.com/spf13/viper"

	"github.com/kinvolk/lokomotive/cli/cmd/cluster"
)

var componentApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Deploy or update a component",
	Long: `Deploy or update a component.
Deploys a component if not yet present, otherwise updates it.
When run with no arguments, all components listed in the configuration are applied.`,
	Run: runApply,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		return cluster.AvailableComponents(), cobra.ShellCompDirectiveNoFileComp
	},
}

var debug bool

//nolint:gochecknoinits
func init() {
	componentCmd.AddCommand(componentApplyCmd)
	pf := componentApplyCmd.PersistentFlags()
	addKubeconfigFileFlag(pf)
	pf.BoolVarP(&debug, "debug", "", false, "Print debug messages")
}

func runApply(cmd *cobra.Command, args []string) {
	contextLogger := log.WithFields(log.Fields{
		"command": "lokoctl component apply",
		"args":    args,
	})

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	options := cluster.ComponentApplyOptions{
		KubeconfigPath: kubeconfigFlag,
		ConfigPath:     viper.GetString("lokocfg"),
		ValuesPath:     viper.GetString("lokocfg-vars"),
	}

	if err := cluster.ComponentApply(contextLogger, args, options); err != nil {
		contextLogger.Fatalf("Applying components failed: %v", err)
	}
}
