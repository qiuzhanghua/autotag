// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"fmt"
	"github.com/qiuzhanghua/autotag/tools"
	"os"

	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write git info to files",
	Long: `Write git info to files of project:

for Go Project, write to autotag.go
for NodeJS Project, write to autotag.js
`,
	Run: func(cmd *cobra.Command, args []string) {
		currentHash, err := tools.GitCurrentHash()
		if err != nil {
			fmt.Println("can't get current hash")
			os.Exit(1)
		}
		fmt.Println("current hash", currentHash)
		fmt.Println(tools.GitCurrenTag())
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
