package service

import (
	"encoding/json"
	"github.com/influxdata/influxdb1-client/models"
	. "github.com/smartystreets/goconvey/convey"
	"hiveon-api/utils"
	"testing"
	"time"
)

func TestGetIndex(t *testing.T) {

	Convey("Test GetIndex pool service method", t, func() {

		curTime := time.Now()
		min30 := curTime.Add(time.Duration(-30) * time.Minute)
		min60 := curTime.Add(time.Duration(-60) * time.Minute)
		min120 := curTime.Add(time.Duration(-120) * time.Minute)
		hashrateConfig := utils.GetConfig().GetFloat64("app.config.pool.hashrate.hashrateCul") /
			utils.GetConfig().GetFloat64("app.config.pool.hashrate.hashrateCulDivider")

		hashrate := make([][]interface{}, 1)
		hashrate[0] = make([]interface{}, 2)
		hashrate[0][0] = min30.Format(time.RFC3339)
		hashrate[0][1] = json.Number("42")
		hashrateRow := models.Row{Values: hashrate}

		miner := make([][]interface{}, 1)
		miner[0] = make([]interface{}, 2)
		miner[0][0] = min60.Format(time.RFC3339)
		miner[0][1] = json.Number("777")
		minerRow := models.Row{Values: miner}

		worker := make([][]interface{}, 1)
		worker[0] = make([]interface{}, 2)
		worker[0][0] = min120.Format(time.RFC3339)
		worker[0][1] = json.Number("2049")
		workerRow := models.Row{Values: worker}

		Convey("When mocking repo", func() {
			poolService := NewPoolServiceWithRepo(mockMinerdashRepo)
			mockMinerdashRepo.On("GetPoolLatestShare").Return(hashrateRow)
			mockMinerdashRepo.On("GetPoolMiner").Return(minerRow)
			mockMinerdashRepo.On("GetPoolWorker").Return(workerRow)
			poolData := poolService.GetIndex()
			Convey("Then expect values:", func() {
				So(poolData.Code, ShouldEqual, 200)
				So(poolData.Data.Hashrate.Time, ShouldEqual, min30.Format(time.RFC3339))
				So(poolData.Data.Hashrate.ValidShares, ShouldEqual, 42)
				So(poolData.Data.Hashrate.Hashrate, ShouldEqual, hashrateConfig * 42)
				So(poolData.Data.Miner.Time, ShouldEqual, min60.Format(time.RFC3339))
				So(poolData.Data.Miner.Count, ShouldEqual, 0.8)
				So(poolData.Data.Worker.Time, ShouldEqual, min120.Format(time.RFC3339))
				So(poolData.Data.Worker.Count, ShouldEqual, 2.0)
			})
		})
	})
}
