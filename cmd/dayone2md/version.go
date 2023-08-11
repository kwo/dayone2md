package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func version(extended bool) string {
	revision, lastCommit := versionInfo()
	if !extended {
		return fmt.Sprintf("version %s", revision)
	}
	return fmt.Sprintf("version %s %s %s %s/%s", revision, lastCommit, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// versionInfo returns commit hash (with dirty flag) and the last commit date.
func versionInfo() (string, string) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", ""
	}
	var revision, lastCommit, dirtyBuild string
	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			revision = kv.Value
		case "vcs.time":
			lastCommit = kv.Value
		case "vcs.modified":
			dirtyBuild = kv.Value
		}
	}
	if len(revision) > 8 {
		revision = revision[0:8]
	}
	if dirtyBuild == "true" {
		revision = revision + "*"
	}
	return revision, lastCommit
}
