// Package tools
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package tools

import (
	"github.com/Masterminds/semver/v3"
	"sort"
	"strings"
)

func GitHeadHash() string {
	// git rev-parse HEAD
	s, err := RunAndReturn("git", "rev-parse", "HEAD")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func GitInstalled() bool {
	_, err := RunAndReturn("git", "--version")
	if err != nil {
		return false
	}
	return true
}

func GitDirIsRepo(p string) bool {
	_, err := RunAndReturn("git", "-C", p, "rev-parse")
	if err != nil {
		return false
	}
	return true
}

func GitAddTag(tag string) (string, error) {
	s, err := RunAndReturn("git", "tag", tag)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

func GitRevOfTag(tag string) string {
	s, err := RunAndReturn("git", "rev-parse", "--short", tag)
	if err == nil {
		return strings.Trim(s, " \n")
	}
	return ""
}

func GitHashOfTag(tag string) string {
	s, err := RunAndReturn("git", "rev-parse", tag)
	if err == nil {
		return strings.Trim(s, " \n")
	}
	return ""
}

// GitDateOfHash
// see https://stackabuse.com/git-show-date-of-a-commit/
func GitDateOfHash(hash string) string {
	s, err := RunAndReturn("git",
		"show", "-s", "--format=%cd", "--date=format:%Y-%m-%d",
		hash)
	if err == nil {
		return strings.Trim(s, " \n")
	}
	return ""

}

func GitLatestTag() string {
	arr, err := GitAllTags()

	if err != nil {
		return "0.0.0"
	}
	n := len(arr)
	if n < 1 || n == 1 && len(arr[0]) == 0 {
		return "0.0.0"
	}
	var vs []semver.Version
	for _, s := range arr {
		v, err := semver.NewVersion(s)
		if err == nil {
			vs = append(vs, *v)
		}
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i].LessThan(&vs[j]) })
	if strings.HasPrefix(arr[0], "v") {
		return "v" + vs[n-1].String()
	}
	return vs[n-1].String()
}

func GitAllTags() ([]string, error) {
	s, err := RunAndReturn("git", "tag")
	if err != nil {
		return nil, err
	}
	s = strings.Trim(s, " \n")
	//fmt.Println("'", s, "'", len(s))
	return strings.Split(s, "\n"), nil
	//s, err := RunAndReturn("git", "rev-list", "--tags")
	//if err != nil {
	//	return nil, err
	//}
	//revs := strings.Split(s, "\n")
	//var tags []string
	//for _, rev := range revs {
	//	t, err := RunAndReturn("git", "describe", "--tags", rev)
	//	//fmt.Println(strings.TrimSpace(t), err)
	//	if err != nil {
	//		//			fmt.Println(err)
	//		continue
	//	}
	//	tags = append(tags, strings.TrimSpace(t))
	//}
	//return tags, nil
}
