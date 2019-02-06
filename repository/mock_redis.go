package repository

import (
	"github.com/stretchr/testify/mock"
)

type MockRedisRepository struct {
	mock.Mock
}

func (m *MockRedisRepository) GetLatestWorker(walletId string)  map[string]string {
	return make(map[string]string)
}
