package main

import (
	"fmt"
	"runtime/debug"
)

func printProgramVersion() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("could not read build info")
	}

	fmt.Printf("semver-bumper %s\n", info.Main.Version)
}
