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

var p1 int
var p2 int
var t string

// pagesCmd represents the hot command
var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "按照页面下载",
	Long:  `按照页面下载`,
	Run: func(cmd *cobra.Command, args []string) {
		dl.Initialize()
		dl.PagesDl(p1, p2, t)
	},
}

func init() {
	rootCmd.AddCommand(pagesCmd)

	pagesCmd.Flags().StringVarP(&t, "type", "t", "hot", "类别: [ori new hot rp long md tf mf rf top lasttop hd]")
	pagesCmd.Flags().IntVar(&p1, "st", 1, "开始页")
	pagesCmd.Flags().IntVar(&p2, "ed", 1, "结束页")
}
