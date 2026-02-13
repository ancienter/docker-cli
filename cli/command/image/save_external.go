package image

import (
	"context"

	"github.com/docker/cli/cli/command"
)

type SaveOptions struct {
	Images    []string
	Output    string
	Platform  string
	Pull      bool
	Untrusted bool
}

func (o *SaveOptions) toInternal() saveOptions {
	return saveOptions{
		images:   o.Images,
		output:   o.Output,
		platform: o.Platform,
	}
}

func RunSave(ctx context.Context, dockerCLI command.Cli, opts SaveOptions) error {
	return runSave(ctx, dockerCLI, opts.toInternal())
}
