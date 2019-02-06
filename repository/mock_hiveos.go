package repository

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type MockHiveosRepository struct {
	mock.Mock
}

func (m *MockHiveosRepository) GetBalance(walletId string) float64 {
	panic("implement me")
}

func (m *MockHiveosRepository) GetBlock24NotUnckle() int {
	args := m.Mock.Called()
	return args.Get(0).(int)
}

func (m *MockHiveosRepository) GetBlock24Uncle() int {
	args := m.Mock.Called()
	return args.Get(0).(int)
}

func (m *MockHiveosRepository) GetBill(walletId string) *sql.Rows {
	args := m.Mock.Called()
	return args.Get(0).(*sql.Rows)
}

func (m *MockHiveosRepository) GetBillInfo(walletId string) RepoBillInfo {
	args := m.Mock.Called()
	return args.Get(0).(RepoBillInfo)
}

