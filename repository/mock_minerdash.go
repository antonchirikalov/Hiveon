package repository

import (
	"github.com/influxdata/influxdb1-client"
	. "github.com/influxdata/influxdb1-client/models"
	"github.com/stretchr/testify/mock"
)

type MockMinerdashRepository struct {
	mock.Mock
}

func (m *MockMinerdashRepository) GetHashrate20m(walletId string, workerId string) client.Result {
	panic("implement me")
}

func (m *MockMinerdashRepository) GetHashrate1d(walletId string) client.Result {
	panic("implement me")
}

func (m *MockMinerdashRepository) GetCountHistory(walletId string) Row {
	panic("implement me")
}

func (m *MockMinerdashRepository) GetBalance(walletId string) int {
	panic("implement me")
}

func (m *MockMinerdashRepository) GetHashrate(walletId string, workerId string) RepoHashrate {
	panic("implement me")
}

func (m *MockMinerdashRepository) GetPoolLatestShare() Row {
	args := m.Mock.Called()
	return args.Get(0).(Row)
}

func (m *MockMinerdashRepository) GetPoolWorker() Row {
	args := m.Mock.Called()
	return args.Get(0).(Row)
}

func (m *MockMinerdashRepository) GetPoolMiner() Row {
	args := m.Mock.Called()
	return args.Get(0).(Row)
}

func (m *MockMinerdashRepository) GetETHHashrate() IncomeCurrency {
	args := m.Mock.Called()
	income := args.Get(0).(IncomeCurrency)
	return income

}

func (m *MockMinerdashRepository) GetShares(walletId string, workerId string) Row {
	args := m.Mock.Called()
	return args.Get(0).(Row)
}

func (m *MockMinerdashRepository) GetLocalHashrateResult(walletId string, workerId string) Row {
	args := m.Mock.Called()
	return args.Get(0).(Row)
}

func (m *MockMinerdashRepository) GetAvgHashrate1d(walletId string, workerId string) client.Result {
	panic("implement me!")
}
