package tools

import (
	"github.com/Masterminds/semver/v3"
	"testing"
)

func TestIncPrerelease01(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3")
	actual := IncPrerelease(*ver)
	expected, _ := semver.NewVersion("v1.2.3")

	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestIncPrerelease02(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3-beta.1+build34")
	actual := IncPrerelease(*ver)
	expected, _ := semver.NewVersion("v1.2.3-beta.2+build34")

	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestIncPrerelease03(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3-beta+build34")
	actual := IncPrerelease(*ver)
	expected, _ := semver.NewVersion("v1.2.3-beta.1+build34")

	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestIncPrerelease04(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3-rc.9+build34")
	actual := IncPrerelease(*ver)
	expected, _ := semver.NewVersion("v1.2.3-rc.10+build34")

	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestNextPhase01(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3")
	actual := NextPhase(*ver)
	expected, _ := semver.NewVersion("v1.2.4-alpha.0")
	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestNextPhase02(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3-alpha.13")
	actual := NextPhase(*ver)
	expected, _ := semver.NewVersion("v1.2.3-beta.0")
	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestNextPhase03(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3-beta.10")
	actual := NextPhase(*ver)
	expected, _ := semver.NewVersion("v1.2.3-rc.0")
	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestNextPhase04(t *testing.T) {
	ver, _ := semver.NewVersion("v1.2.3-rc.3")
	actual := NextPhase(*ver)
	expected, _ := semver.NewVersion("v1.2.3")
	if !actual.Equal(expected) {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}
