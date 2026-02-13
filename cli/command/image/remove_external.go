package image

import (
	"context"

	"github.com/docker/cli/cli/command"
)

type RemoveOptions struct {
	Force     bool
	NoPrune   bool
	Platforms []string
}

func (o *RemoveOptions) toInternal() removeOptions {
	return removeOptions{
		force:     o.Force,
		noPrune:   o.NoPrune,
		platforms: o.Platforms,
	}
}

func RunRemove(ctx context.Context, dockerCLI command.Cli, opts RemoveOptions, images []string) error {
	return runRemove(ctx, dockerCLI, opts.toInternal(), images)
}
