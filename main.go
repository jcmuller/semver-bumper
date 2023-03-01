package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/pflag"
	gosemver "golang.org/x/mod/semver"
)

func main() {
	setupFlags()

	if showVersion {
		printProgramVersion()
		os.Exit(0)
	}

	version, err := getVersion(pflag.Args())
	if err != nil {
		err = fmt.Errorf("error getting version: %w", err)
		fail(err)
	}

	ver, err := semver.NewVersion(version)
	if err != nil {
		err = fmt.Errorf("error parsing version: %w", err)
		fail(err)
	}

	if !show {
		err := bumpVersion(ver)
		if err != nil {
			fail(err)
		}
	}

	fmt.Printf("v%s\n", ver.String())
}

func fail(i any) {
	fmt.Fprintln(os.Stderr, i)
	os.Exit(1)
}

func getVersion(versions []string) (string, error) {
	var err error
	var readFromStdin bool

	if len(versions) == 0 || versions[0] == "-" {
		readFromStdin = true
	}

	if readFromStdin {
		versions, err = readInput()
		if err != nil {
			err = fmt.Errorf("error reading input: %w", err)
			return "", err
		}

	}

	gosemver.Sort(versions)

	return versions[len(versions)-1], nil
}

func bumpVersion(ver *semver.Version) error {
	switch increment {
	case "major":
		*ver = ver.IncMajor()
	case "minor":
		*ver = ver.IncMinor()
	case "patch":
		*ver = ver.IncPatch()
	}

	var err error

	*ver, err = ver.SetMetadata(metadata)
	if err != nil {
		err = fmt.Errorf("error setting metadata: %w", err)
		return err
	}
	*ver, err = ver.SetPrerelease(prerelease)
	if err != nil {
		err = fmt.Errorf("error setting pre-release: %w", err)
		return err
	}

	return nil
}

func readInput() ([]string, error) {
	allVersions, err := io.ReadAll(os.Stdin)
	if err != nil {
		err = fmt.Errorf("error reading from stdin: %w", err)
		return nil, err
	}

	return strings.Split(string(allVersions), "\n"), nil
}
