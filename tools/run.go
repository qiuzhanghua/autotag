// Package tools
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package tools

import (
	"os"
	"os/exec"
)

func RunAndReturn(command string, args ...string) (string, error) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil

}

func RunAndForget(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunAndDetach
// tested under darwin
func RunAndDetach(command string, args ...string) error {
	args = append(args, "--detached")
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Process.Release()
}
