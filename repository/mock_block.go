package repository

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type MockBlockRepository struct {
	mock.Mock
}

func (m *MockBlockRepository) GetIncome7d() float64 {
	args := m.Mock.Called()
	return args.Get(0).(float64)
}

func (m *MockBlockRepository) GetIncome24h() float64 {
	args := m.Mock.Called()
	return args.Get(0).(float64)
}

func (m *MockBlockRepository) GetIncomeResult() float64 {
	args := m.Mock.Called()
	return args.Get(0).(float64)
}

func (m *MockBlockRepository) GetIncomeHistory() *sql.Rows{
	panic("implement me!")
}
