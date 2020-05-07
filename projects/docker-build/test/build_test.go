package test

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	dockerClient "github.com/docker/docker/client"
)

const dockerfileA = "FROM alpine\n"

func TestDockerBuild(t *testing.T) {
	cli, err := dockerClient.NewEnvClient()
	if err != nil {
		panic(err)
	}

	dockerfileAReader := strings.NewReader(dockerfileA)
	response, err := cli.ImageBuild(context.Background(), dockerfileAReader, types.ImageBuildOptions{
		Tags: []string{"test"},
	})
	if err != nil {
		panic(err)
	}

	respBuffer, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		panic(err)
	}

	t.Log(string(respBuffer))
}
