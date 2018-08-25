// Copyright © 2018 ilove91
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
	"github.com/ilove91/91dl/dl"
	"github.com/spf13/cobra"
)

var vlinks []string

// linksCmd represents the links command
var linksCmd = &cobra.Command{
	Use:   "links",
	Short: "Download videos by links",
	Long:  `Download videos by links`,
	Run: func(cmd *cobra.Command, args []string) {
		dl.Initialize()
		dl.LinksDl(vlinks)
	},
}

func init() {
	rootCmd.AddCommand(linksCmd)

	linksCmd.Flags().StringSliceVarP(&vlinks, "videos", "v", nil, "videos links")
}
