package gateways

import "context"

type AWSS3Gateway struct {
}

func NewS3Client(ctx context.Context) *AWSS3Gateway {
	return &AWSS3Gateway{}
}
