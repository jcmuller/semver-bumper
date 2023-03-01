package main

import "github.com/spf13/pflag"

var (
	increment   = "patch"
	prerelease  string
	metadata    string
	show        bool
	showVersion bool
)

func setupFlags() {
	pflag.StringVarP(&increment, "increment", "i", increment, "Increment [major|minor|patch]. Default level is patch")
	pflag.StringVarP(&prerelease, "pre-release", "p", "", "Set pre-release version")
	pflag.StringVarP(&metadata, "metadata", "m", "", "Set metadata version")
	pflag.BoolVarP(&show, "show", "s", false, "Show passed in version and exit")
	pflag.BoolVarP(&showVersion, "version", "v", false, "Show semver's version and exit")

	pflag.Parse()
}
