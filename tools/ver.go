// Package tools
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package tools

import (
	"github.com/Masterminds/semver/v3"
	"strconv"
	"strings"
)

//IncPrerelease
// a.b.c-alpha.x => a.b.c-alpha.(x+1)
// a.b.c-beta.x => a.b.c-beta.(x+1)
// a.b.c => a.b.c    // unchanged
func IncPrerelease(v semver.Version) (vNext semver.Version) {
	vNext = v
	pre := v.Prerelease()
	//fmt.Println(pre)
	if len(pre) < 2 {
		return
	}
	arr := strings.Split(pre, ".")
	//fmt.Println(arr)
	n := len(arr)
	if n == 1 {
		vNext, _ = vNext.SetPrerelease(pre + ".1")
		//fmt.Println(vNext)
		return
	}
	if n == 2 {
		x, err := strconv.ParseInt(arr[1], 10, 32)
		if err != nil {
			x = 0
		}
		x += 1
		//fmt.Println(x)
		vNext, _ = vNext.SetPrerelease(arr[0] + "." + strconv.Itoa(int(x)))
		//fmt.Println(arr[0] + strconv.Itoa(int(x)))
		//fmt.Println(vNext)
		return
	}
	if n > 2 {
		vNext, _ = vNext.SetPrerelease("")
	}
	return
}

//NextPhase
//  a.b.c => a.b.(c+1)-alpha.0
//  a.b.c-alpha.x => a.b.(c+1)-beta.0
//  a.b.c-beta.x => a.b.c
func NextPhase(v semver.Version) (vNext semver.Version) {
	vNext = v
	pre := v.Prerelease()
	if len(pre) < 4 {
		vNext = vNext.IncPatch()
		vNext, _ = vNext.SetPrerelease("alpha.0")
		return
	}
	if strings.HasPrefix(pre, "alpha") {
		vNext, _ = vNext.SetPrerelease("beta.0")
		return
	}
	if strings.HasPrefix(pre, "beta") {
		vNext, _ = vNext.SetPrerelease("rc.0")
		return
	}
	if strings.HasPrefix(pre, "rc") {
		vNext = vNext.IncPatch()
		return
	}
	return
}
