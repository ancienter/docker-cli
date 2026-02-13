package image

import (
	"context"

	"github.com/docker/cli/cli/command"
)

type PushOptions struct {
	All       bool
	Remote    string
	Untrusted bool
	Quiet     bool
	Platform  string
}

func (o *PushOptions) toInternal() pushOptions {
	return pushOptions{
		all:       o.All,
		remote:    o.Remote,
		untrusted: o.Untrusted,
		quiet:     o.Quiet,
		platform:  o.Platform,
	}
}

func RunPush(ctx context.Context, dockerCli command.Cli, opts PushOptions) error {
	return runPush(ctx, dockerCli, opts.toInternal())
}
