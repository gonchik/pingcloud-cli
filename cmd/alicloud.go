/*
Copyright © 2024 Gonchik Tsymzhitov <gonchik.tsymzhitov@atlassiancommunity.com>

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
	"github.com/reoim/pingcloud-cli/ping"
	"github.com/spf13/cobra"
)

// aliCmd represents the AliCloud command
var aliCmd = &cobra.Command{
	Use:   "ali",
	Short: "Check latencies of AliCloud (Aliyun) regions.",
	Run: func(cmd *cobra.Command, args []string) {
		c := ping.CmdOption{
			Option:     "alicloud",
			OptionName: "Ali",
			ListFlg:    list,
			Args:       args,
		}

		c.StartCmd()
	},
}

func init() {
	rootCmd.AddCommand(aliCmd)
	aliCmd.Flags().BoolVarP(&list, "list", "l", false, "List all available regions")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
