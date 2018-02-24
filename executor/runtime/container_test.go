package runtime

import (
	"testing"

	"github.com/Netflix/titus-executor/api/netflix/titus"
	"github.com/Netflix/titus-executor/config"
	runtimeTypes "github.com/Netflix/titus-executor/executor/runtime/types"
	protobuf "github.com/golang/protobuf/proto"
)

func TestImageNameWithTag(t *testing.T) {
	// TODO(fabio): no tight coupling with the (global) config package
	config.Load("../mock/config.json")

	expected := "docker.io/titusoss/alpine:latest"
	c := &runtimeTypes.Container{
		TitusInfo: &titus.ContainerInfo{
			ImageName: protobuf.String("titusoss/alpine"),
			Version:   protobuf.String("latest"),
		},
	}
	if got := c.QualifiedImageName(); got != expected {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestImageTagLatestByDefault(t *testing.T) {
	// TODO(fabio): no tight coupling with the (global) config package
	config.Load("../mock/config.json")

	expected := "docker.io/titusoss/alpine:latest"
	c := &runtimeTypes.Container{
		TitusInfo: &titus.ContainerInfo{
			ImageName: protobuf.String("titusoss/alpine"),
		},
	}
	if got := c.QualifiedImageName(); got != expected {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestImageByDigest(t *testing.T) {
	// TODO(fabio): no tight coupling with the (global) config package
	config.Load("../mock/config.json")

	expected := "docker.io/" +
		"titusoss/alpine@sha256:58e1a1bb75db1b5a24a462dd5e2915277ea06438c3f105138f97eb53149673c4"
	c := &runtimeTypes.Container{
		TitusInfo: &titus.ContainerInfo{
			ImageName:   protobuf.String("titusoss/alpine"),
			ImageDigest: protobuf.String("sha256:58e1a1bb75db1b5a24a462dd5e2915277ea06438c3f105138f97eb53149673c4"),
		},
	}
	if got := c.QualifiedImageName(); got != expected {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestImageByDigestIgnoresTag(t *testing.T) {
	// TODO(fabio): no tight coupling with the (global) config package
	config.Load("../mock/config.json")

	expected := "docker.io/" +
		"titusoss/alpine@sha256:58e1a1bb75db1b5a24a462dd5e2915277ea06438c3f105138f97eb53149673c4"
	c := &runtimeTypes.Container{
		TitusInfo: &titus.ContainerInfo{
			ImageName:   protobuf.String("titusoss/alpine"),
			Version:     protobuf.String("latest"),
			ImageDigest: protobuf.String("sha256:58e1a1bb75db1b5a24a462dd5e2915277ea06438c3f105138f97eb53149673c4"),
		},
	}
	if got := c.QualifiedImageName(); got != expected {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}
