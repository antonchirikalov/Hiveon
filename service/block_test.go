package service

import (
	"github.com/labstack/gommon/log"
	. "github.com/smartystreets/goconvey/convey"
	"hiveon-api/repository"
	"os"
	"testing"
)


var (
	mockBlockRepo     *repository.MockBlockRepository
	mockMinerdashRepo *repository.MockMinerdashRepository
	mockHiveosRepo    *repository.MockHiveosRepository
	mockRedisRepo     *repository.MockRedisRepository
	testBlockService BlockService
	testMinerService MinerService
	testPoolService PoolService

)

func TestMain(m *testing.M) {
	log.Info("Setup mocks")
	mockMinerdashRepo = &repository.MockMinerdashRepository{}
	mockBlockRepo = &repository.MockBlockRepository{}
	mockHiveosRepo = &repository.MockHiveosRepository{}
	mockRedisRepo = &repository.MockRedisRepository{}

	testBlockService =  NewBlockServiceWithRepo(mockHiveosRepo)
	testMinerService = NewMinerServiceWithRepo(mockBlockRepo, mockMinerdashRepo, mockHiveosRepo,mockRedisRepo)
	testPoolService = NewPoolServiceWithRepo(mockMinerdashRepo)
	os.Exit(m.Run())
}

func TestGetBlockCount(t *testing.T){
	Convey("Test GetBlockCount block service method", t, func() {
		Convey("When mocking hiveosRepository", func() {
			mockHiveosRepo.On("GetBlock24NotUnckle").Return(48)
			mockHiveosRepo.On("GetBlock24Uncle").Return(2)
			blockData := testBlockService.GetBlockCount()
			Convey("Then expect values:", func() {
				So(blockData.Data.Blocks, ShouldEqual, 48)
				So(blockData.Data.Uncles, ShouldEqual, 2)
			})
		})
	})
}
