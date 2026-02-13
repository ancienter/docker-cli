package image

import (
	"bytes"
	"context"
	"io"

	"github.com/containerd/platforms"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/inspect"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type InspectOptions struct {
	Format   string
	Refs     []string
	Platform string
}

func (o *InspectOptions) toInternal() inspectOptions {
	return inspectOptions{
		format:   o.Format,
		refs:     o.Refs,
		platform: o.Platform,
	}
}

func RunInspect(ctx context.Context, dockerCLI command.Cli, opts InspectOptions) error {
	return runInspect(ctx, dockerCLI, opts.toInternal())
}

func RunInspectRaw(ctx context.Context, dockerCLI command.Cli, opts InspectOptions) error {
	return runInspectRaw(ctx, dockerCLI, opts.toInternal())
}

func runInspectRaw(ctx context.Context, dockerCLI command.Cli, opts inspectOptions) error {
	var platform *ocispec.Platform
	if opts.platform != "" {
		p, err := platforms.Parse(opts.platform)
		if err != nil {
			return err
		}
		platform = &p
	}

	apiClient := dockerCLI.Client()
	return inspect.Inspect(io.Discard, opts.refs, opts.format, func(ref string) (interface{}, []byte, error) {
		var buf bytes.Buffer
		resp, err := apiClient.ImageInspect(ctx, ref,
			client.ImageInspectWithRawResponse(&buf),
			client.ImageInspectWithPlatform(platform),
		)
		if err != nil {
			return image.InspectResponse{}, nil, err
		}
		return resp, buf.Bytes(), err
	})
}
