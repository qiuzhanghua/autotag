/*
Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
*/
package main

import (
	"fmt"
	"github.com/qiuzhanghua/autotag/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autotag",
	Short: "AutoTag application",
	Long:  `AutoTag application for tag add/(write to file)/(calculate next).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please run:\n    autotag --help")
	},
}

func main() {
	cmd.VersionCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Printf("autotag %s (%s %s)\n", AppVersion, AppRevision, AppBuildDate)
	}
	rootCmd.AddCommand(cmd.VersionCmd, cmd.ShowCmd, cmd.NextCmd, cmd.WriteCmd)
	rootCmd.Execute()
}
