package image

import (
	"context"

	"github.com/docker/cli/cli/command"
)

type TagOptions struct {
	Image string
	Name  string
}

func (o *TagOptions) toInternal() tagOptions {
	return tagOptions{
		image: o.Image,
		name:  o.Name,
	}
}

func RunTag(ctx context.Context, dockerCli command.Cli, opts TagOptions) error {
	return dockerCli.Client().ImageTag(ctx, opts.Image, opts.Name)
}
