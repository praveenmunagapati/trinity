package main

import (
	"strconv"

	"github.com/clownpriest/trinity/core/system"
)

const (
	MAJOR = 0
	MINOR = 0
	PATCH = 1
)

var major = strconv.Itoa(MAJOR)
var minor = strconv.Itoa(MINOR)
var patch = strconv.Itoa(PATCH)

var versionNumStr = major + "." + minor + "." + patch

func version() string {
	return versionNumStr + "\n" + system.CommitHead[:15]
}
