// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"bufio"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write git info to files",
	Long: `Write git info to files of project:

for Go Project, write to autotag.go
for Node.js Project, write to autotag.js
`,
	Run: func(cmd *cobra.Command, args []string) {
		latest := tools.GitLatestTag()
		_, err := semver.NewVersion(latest)
		if err != nil {
			log.Fatal(err)
		}
		rev := tools.GitRevOfTag(latest)
		_, err = os.Stat("go.mod")
		if err == nil { // Go Project
			file, err := os.Create("./autotag.go")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			writer := bufio.NewWriter(file)
			linesToWrite := []string{"package main\n", fmt.Sprintf("const AppVersion = \"%s\"", latest), fmt.Sprintf("const AppRevision = \"%s\"", rev)}
			for _, line := range linesToWrite {
				_, err := writer.WriteString(line + "\n")
				if err != nil {
					log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
				}
			}
			writer.Flush()
		}
		_, err = os.Stat("package.json")
		if err == nil { // Node.js Project
			file, err := os.Create("./autotag.js")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			writer := bufio.NewWriter(file)
			linesToWrite := []string{fmt.Sprintf("const AppVersion = \"%s\"", latest), fmt.Sprintf("const AppRevision = \"%s\"", rev)}
			for _, line := range linesToWrite {
				_, err := writer.WriteString(line + "\n")
				if err != nil {
					log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
				}
			}
			writer.Flush()

		}

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
