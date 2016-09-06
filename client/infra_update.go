package client

import (
	"golang.org/x/net/context"

	"github.com/docker/engine-api/types/infra"
)

// InfraUpdate updates the infrastructure
func (cli *Client) InfraUpdate(ctx context.Context, config infra.Config) error {
	resp, err := cli.post(ctx, "/infra/update", nil, config, nil)
	ensureReaderClosed(resp)
	return err
}
