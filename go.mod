module github.com/iwilltry42/confd

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/aws/aws-sdk-go v1.38.63
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/docker/cli v20.10.7+incompatible
	github.com/docker/docker v20.10.7+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/garyburd/redigo v1.6.2
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/vault/api v1.1.0
	github.com/kelseyhightower/memkv v0.1.1
	github.com/moby/term v0.0.0-20210610120745-9d4ed1856297 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/sirupsen/logrus v1.7.0
	github.com/xordataexchange/crypt v0.0.3-0.20170626215501-b2862e3d0a77
	go.etcd.io/etcd/client/v3 v3.5.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
