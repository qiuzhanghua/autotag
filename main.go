/*
Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
*/
package main

import (
	"fmt"
	"github.com/qiuzhanghua/autotag/cmd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.SetUsageFunc(func(command *cobra.Command) error {
		fmt.Print(`Usage:
    autotag help
    autotag version
    autotag show
    autotag write
    autotag next [pre|phase|patch|minor|major]

`)
		return nil
	})

}

var rootCmd = &cobra.Command{
	Use:   "autotag",
	Short: "AutoTag application",
	Long:  `AutoTag application for tag add/(write to file)/(calculate next).`,
}

func main() {
	cmd.VersionCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Printf("autotag %s (%s %s)\n", AppVersion, AppRevision, AppBuildDate)
	}
	rootCmd.AddCommand(cmd.VersionCmd, cmd.ShowCmd, cmd.NextCmd, cmd.WriteCmd)
	_ = rootCmd.Execute()
}
