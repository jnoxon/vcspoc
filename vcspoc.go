package main

import (
	"log"
	"runtime"
	"runtime/debug"
)

func main() {
	vcsrevision := "unknown"
	vcsdirty := ""
	vcstime := "unknown"

	// requires go1.18
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, v := range info.Settings {
			switch v.Key {
			case "vcs.revision":
				vcsrevision = v.Value
			case "vcs.modified":
				if v.Value == "true" {
					vcsdirty = " (dirty)"
				}
			case "vcs.time":
				vcstime = v.Value
			}
		}
	}

	log.Printf(" git_commit=%s%s; commit_date=%s; go_version=%s", vcsrevision, vcsdirty, vcstime, runtime.Version())
}
