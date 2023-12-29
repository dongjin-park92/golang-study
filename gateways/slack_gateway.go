package gateways

import "context"

type SlackGateway struct {
}

func NewSlackClient(ctx context.Context) *SlackGateway {
	return &SlackGateway{}
}
