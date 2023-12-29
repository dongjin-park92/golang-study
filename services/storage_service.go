package services

import (
	"context"
	"golang-go/gateways"
)

type StorageService struct {
	storageGateway *gateways.AWSS3Gateway
}

func NewStorageService(storageGateway *gateways.AWSS3Gateway) *StorageService {
	return &StorageService{
		storageGateway: storageGateway,
	}
}

func (s StorageService) GetObjectExist(ctx context.Context) error {
	return nil
}
