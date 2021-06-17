package docker

import (
	"os"
	"strings"

	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/client"
	dockerclient "github.com/docker/docker/client"
	log "github.com/iwilltry42/confd/log"
	"github.com/pkg/errors"
)

/*
 * StoreClient Interface
 */

type DockerClient struct {
	client *dockerclient.Client
}

func NewDockerClient() (*DockerClient, error) {
	c, err := getDockerClient()
	if err != nil {
		return nil, err
	}

	return &DockerClient{
		client: c,
	}, nil
}

func GetValues(keys []string) (map[string]string, error) {
	log.Info("Test")
	return nil, nil
}

func WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	return 1, nil
}

/*
 * Helper Functions
 */

func getDockerClient() (*dockerclient.Client, error) {
	var err error
	var cli *client.Client

	dockerHost := os.Getenv("DOCKER_HOST")

	if strings.HasPrefix(dockerHost, "ssh://") {
		var helper *connhelper.ConnectionHelper

		helper, err = connhelper.GetConnectionHelper(dockerHost)
		if err != nil {
			return nil, err
		}
		cli, err = client.NewClientWithOpts(
			client.WithHost(helper.Host),
			client.WithDialContext(helper.Dialer),
			client.WithAPIVersionNegotiation(),
		)
	} else {
		cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cli, err
}
