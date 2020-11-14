module github.com/spiegel-im-spiegel/depm

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/emicklei/dot v0.15.0
	github.com/spf13/cobra v1.1.1
	github.com/spiegel-im-spiegel/errs v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.10.3
	golang.org/x/tools v0.0.0-20201113202037-1643af1435f3
)

replace github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
