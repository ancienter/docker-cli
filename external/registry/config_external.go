package registry

import (
	"github.com/distribution/reference"
	r "github.com/docker/cli/internal/registry"
	"github.com/docker/docker/api/types/registry"
)

func NewIndexInfo(reposName reference.Named) *registry.IndexInfo {
	return r.NewIndexInfo(reposName)
}
