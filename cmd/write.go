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
	"time"
)

// WriteCmd represents the write command
var WriteCmd = &cobra.Command{
	Use:   "write",
	Short: "Write git info to files",
	Long: `Write git info to files of project:

for Go Project, write to autotag.go
for Node.js Project, write to autotag.js
`,
	Run: func(cmd *cobra.Command, args []string) {
		latest := tools.GitLatestTag()
		if len(latest) == 0 {
			log.Fatalf("No tag defined in this repository")
		}
		_, err := semver.NewVersion(latest)
		if err != nil {
			log.Fatal(err)
		}
		rev := tools.GitRevOfTag(latest)
		dataString := time.Now().Format("2006-01-02")
		_, err = os.Stat("go.mod")
		if err == nil { // Go Project
			file, err := os.Create("./autotag.go")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			writer := bufio.NewWriter(file)
			linesToWrite := []string{"package main\n",
				fmt.Sprintf("var AppVersion = \"%s\"", latest),
				fmt.Sprintf("var AppRevision = \"%s\"", rev),
				fmt.Sprintf("var AppBuildDate = \"%s\"", dataString),
			}
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
			linesToWrite := []string{fmt.Sprintf("const AppVersion = \"%s\"", latest),
				fmt.Sprintf("const AppRevision = \"%s\"", rev),
				fmt.Sprintf("const AppBuildDate = \"%s\"", dataString),
			}
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
