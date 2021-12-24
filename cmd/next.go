// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	NextCmd.SetUsageFunc(func(command *cobra.Command) error {

		fmt.Print(`Usage:
    autotag next [pre|phase|patch|minor|major]`)
		return nil
	})
}

// NextCmd represents the next command
var NextCmd = &cobra.Command{
	Use:   "next",
	Short: "Add tag for next",
	Long:  `Add next (pre/phase/patch/minor/major) tag.`,
	Args: func(cmd *cobra.Command, args []string) error {
		opts := []string{"pre", "phase", "patch", "minor", "major"}
		if len(args) == 1 {
			for _, s := range opts {
				if s == args[0] {
					return nil
				}
			}
		}
		return errors.New("argument should be one of (" + strings.Join(opts, " ") + ")")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !tools.GitInstalled() {
			fmt.Println("Git not installed.")
			return
		}
		if !tools.GitDirIsRepo(".") {
			fmt.Println("Current directory is not a git repository.")
			return
		}
		if len(tools.GitHeadHash()) == 0 {
			fmt.Println("Current git repository has not any commit.")
			return
		}
		if len(args) == 1 && (args[0] == "pre" || args[0] == "phase" || args[0] == "patch" || args[0] == "minor" || args[0] == "major") {
			latest := tools.GitLatestTag()
			if len(latest) < 1 {
				latest = "0.0.0"
			}
			v, err := semver.NewVersion(latest)
			if err != nil {
				fmt.Printf("%s is not a Semver\n", latest)
				return
			}
			var next string
			if args[0] == "pre" {
				next = tools.IncPrerelease(*v).String()
			} else if args[0] == "phase" {
				next = tools.NextPhase(*v).String()
			} else if args[0] == "patch" {
				next = (*v).IncPatch().String()
			} else if args[0] == "minor" {
				next = (*v).IncMinor().String()
			} else if args[0] == "major" {
				next = (*v).IncMajor().String()
			}
			if strings.HasPrefix(latest, "v") {
				next = "v" + next
			}
			if next == latest {
				fmt.Printf("Add tag %s is not needed.\n", next)
				return
			}
			ret, err := tools.GitAddTag(next)
			if err != nil {
				fmt.Printf("Add tag %s error: %s\n", next, err)
			} else {
				fmt.Printf("Add tag %s ok. %s\n", next, ret)
			}
		} else {
			fmt.Print(`to do next "git tag <next_tag_name>", please use:
autotag next pre
autotag next phase
autotag next patch
autotag next minor
autotag next major
`)
		}
	},
}
