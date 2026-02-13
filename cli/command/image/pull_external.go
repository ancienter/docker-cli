package image

import (
	"context"

	"github.com/docker/cli/cli/command"
)

type PullOptions struct {
	Remote    string
	All       bool
	Platform  string
	Quiet     bool
	Untrusted bool
}

func (o *PullOptions) toInternal() pullOptions {
	return pullOptions{
		remote:    o.Remote,
		all:       o.All,
		platform:  o.Platform,
		quiet:     o.Quiet,
		untrusted: o.Untrusted,
	}
}

func RunPull(ctx context.Context, dockerCLI command.Cli, opts PullOptions) error {
	return runPull(ctx, dockerCLI, opts.toInternal())
}
