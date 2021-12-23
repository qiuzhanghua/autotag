// Package tools
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GitLastHash() (string, error) {
	// git rev-list --tags --max-count=1
	s, err := RunAndReturn("git", "rev-list", "--tags", "--max-count=1")
	return strings.TrimSpace(s), err
}

func GitCurrentHash() (string, error) {
	// git rev-parse HEAD
	s, err := RunAndReturn("git", "rev-parse", "HEAD")
	return strings.TrimSpace(s), err
}

func GitCurrentBranch() (string, error) {
	// git rev-parse --abbrev-ref f0c45331af2e8386f1e4ad5ef1946cba321d408f
	s, err := RunAndReturn("git", "rev-parse", "--abbrev-ref", "HEAD")
	return strings.TrimSpace(s), err
}

func GitCurrenTag() (string, error) {
	// git describe --tags f0c45331af2e8386f1e4ad5ef1946cba321d408f
	hash, err := GitCurrentHash()
	if err != nil {
		return "", err
	}
	s, err := RunAndReturn("git", "describe", "--tags", hash)
	return strings.TrimSpace(s), err
}

func GitGetConfig(args ...string) string {
	args = append([]string{"config", "--get"}, args...)
	s, err := RunAndReturn("git", args...)
	if err == nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func GitGetConfigBool(args ...string) bool {
	args = append([]string{"--bool"}, args...)
	return GitGetConfig(args...) == "true"
}

func GitClosestVersion() string {
	// git describe --tags --abbrev=0
	s, err := RunAndReturn("git", "describe", "--abbrev=0")
	if err != nil {
		return ""
	}
	arr := strings.Split(strings.TrimSpace(s), "/")
	tag := ""
	for _, item := range arr {
		if strings.HasPrefix(item, "v") {
			tag = item
		}
	}
	return tag
}

func bumpVersion(ver string, part int) string {
	prefix, parts := versionParts(ver)
	parts[part]++
	for i := part + 1; i < len(parts); i++ {
		parts[i] = 0
	}
	return versionString(prefix, parts)
}

func versionString(prefix string, parts []int) string {
	ver := fmt.Sprintf("%s%d", prefix, parts[0])
	for _, part := range parts[1:] {
		ver = fmt.Sprintf("%s.%d", ver, part)
	}
	return ver
}

// versionParts matches a px.y.z version, for non-digit values of p and digits
// x, y, and z.
func versionParts(s string) (prefix string, parts []int) {
	exp := regexp.MustCompile(`^([^\d]*)(\d+)\.(\d+)\.(\d+)$`)
	match := exp.FindStringSubmatch(s)
	if len(match) > 1 {
		prefix = match[1]
		parts = make([]int, len(match)-2)
		for i := range parts {
			parts[i], _ = strconv.Atoi(match[i+2])
		}
	}
	return
}
