package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	semver "github.com/Masterminds/semver/v3"
	"github.com/spf13/pflag"
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

func getVersion(args []string) (string, error) {
	var version string
	var readFromStdin bool

	switch len(args) {
	case 0:
		readFromStdin = true
	case 1:
		version = args[0]
		readFromStdin = version == "-"
	default:
		return "", fmt.Errorf("invalid version supplied. Either pass it in as STDIN, or as the only argument to this program")
	}

	if readFromStdin {
		v, err := readInput(os.Stdin)
		if err != nil {
			err = fmt.Errorf("error reading stdin: %w", err)
			return "", err
		}

		version = string(v)
	}

	return version, nil
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

func readInput(reader io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if _, err := buf.Write(scanner.Bytes()); err != nil {
			err = fmt.Errorf("error writing to buffer: %w", err)
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("error scanning: %w", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
