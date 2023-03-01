# semver-bumper

A semantic version bumper

## Installation

```
go install github.com/jcmuller/semver-bumper@latest
```

## Usage

```
$ semver-bumper v1.2.3
v1.2.4
```

```
# latest tag is v2.1.4
$ git tag --list 'v*' --sort=version:refname --merged | tail -n1 | semver-bumper --increment minor
v2.2.0
```

```
$ semver-bumper --help
Usage of semver-bumper:
  -i, --increment string     Increment [major|minor|patch]. Default level is patch (default "patch")
  -m, --metadata string      Set metadata version
  -p, --pre-release string   Set pre-release version
  -s, --show                 Show passed in version and exit
  -v, --version              Show semver-bumper's version and exit
