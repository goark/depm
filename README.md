# [depm] -- Visualize depndency packages and modules

[![check vulns](https://github.com/spiegel-im-spiegel/depm/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/depm/actions)
[![lint status](https://github.com/spiegel-im-spiegel/depm/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/depm/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/depm/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/depm.svg)](https://github.com/spiegel-im-spiegel/depm/releases/latest)

## Download and Build

```
$ go get github.com/spiegel-im-spiegel/depm@latest
```

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/depm/releases/latest).

## Usage

```
$ depm -h
Visualize depndency packages and modules.

Usage:
  depm [flags]
  depm [command]

Available Commands:
  help        Help about any command
  list        list modules
  module      analyze depndency modules
  package     analyze depndency packages
  version     print the version number

Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable
  -h, --help                 help for depm

Use "depm [command] --help" for more information about a command.
```

### Analyze Depndency Packages

```
$ depm package -h
analyze depndency packages.

Usage:
  depm package [flags] [package import path]

Aliases:
  package, pkg, p

Flags:
      --dot                 output by DOT language
      --dot-config string   config file for DOT language
  -h, --help                help for package
  -i, --include-internal    include internal packages
  -s, --include-standard    include standard Go library

Global Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable

$ depm package "github.com/spiegel-im-spiegel/depm" | jq .
[
  {
    "Package": {
      "ImportPath": "github.com/spf13/cobra",
      "Module": {
        "Path": "github.com/spf13/cobra",
        "Version": "v1.1.1"
      }
    },
    "Deps": [
      {
        "Package": {
          "ImportPath": "github.com/spf13/pflag",
          "Module": {
            "Path": "github.com/spf13/pflag",
            "Version": "v1.0.5"
          }
        }
      }
    ]
  },
...
```

### Analyze Depndency Modules

```
$ depm module -h
analyze depndency modules.

Usage:
  depm module [flags] [package import path]

Aliases:
  module, mod, m

Flags:
  -u, --check-update        check updating module
      --dot                 output by DOT language
      --dot-config string   config file for DOT language
  -h, --help                help for module

Global Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable

$ depm module "github.com/spiegel-im-spiegel/depm" | jq .
[
  {
    "Module": {
      "Path": "github.com/spf13/cobra@v1.1.1",
      "Packages": [
        "github.com/spf13/cobra"
      ]
    },
    "Deps": [
      {
        "Path": "github.com/spf13/pflag@v1.0.5",
        "Packages": [
          "github.com/spf13/pflag"
        ]
      }
    ]
  },
  ...
```

### Output DOT Language Format with Configuration File

```toml
[node]
  fontname = "Inconsolata"
[edge]
  color = "red"
```

```
$ depm module --dot --dot-config dotenc/sample.toml "github.com/spiegel-im-spiegel/depm" | dot -Tpng -o output.png
```

[![output.png](./output.png)](./output.png)

### Output List of Modules (compatible 'go list -m' command)

```
$ depm list -h
list modules, compatible 'go list -m' command

Usage:
  depm list [flags] [package import path]

Aliases:
  list, lst, l

Flags:
  -u, --check-update   check updating module
  -h, --help           help for list
  -j, --json           output by JSON format

Global Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable

$ depm list -u "github.com/spiegel-im-spiegel/depm"
github.com/BurntSushi/toml v0.3.1
github.com/emicklei/dot v0.15.0
github.com/spf13/cobra v1.1.1
github.com/spf13/pflag v1.0.5
github.com/spiegel-im-spiegel/depm
github.com/spiegel-im-spiegel/errs v1.0.2
github.com/spiegel-im-spiegel/gocli v0.10.3
golang.org/x/mod v0.3.0
golang.org/x/net v0.0.0-20201021035429-f5854403a974 [v0.0.0-20201031054903-ff519b6c9102]
golang.org/x/tools v0.0.0-20201105220310-78b158585360
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
```

[depm]: https://github.com/spiegel-im-spiegel/depm "spiegel-im-spiegel/depm: Visualize depndency packages and modules"
