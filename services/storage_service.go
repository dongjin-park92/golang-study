package services

import (
	"context"
	"fmt"
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

func (s StorageService) GetObjectExist(ctx context.Context, name string) error {
	metricName := fmt.Sprintf("check_file_total_call_%s", name)
	// ---
	// 비즈니스 로직 실행 후 메트릭 증가
	// ---
	err := generateMetrics(metricName)
	if err != nil {
		return err
	}

	return nil
}
