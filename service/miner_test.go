package service

import (
	"encoding/json"
	"github.com/influxdata/influxdb1-client/models"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"hiveon-api/repository"
	"hiveon-api/utils"
	"math"
	"testing"
	"time"
)

func TestGetFutureIncome(t *testing.T) {

	Convey("Test GetFutureIncome miner service method", t, func() {
		Convey("When mocking hiveosRepository", func() {
			minerService := NewMinerServiceWithRepo(mockBlockRepo, mockMinerdashRepo, mockHiveosRepo, mockRedisRepo)
			mockMinerdashRepo.On("GetETHHashrate").Return(repository.IncomeCurrency{USD: 13.21, BTC: 8.13, CNY: 5.8}).Once()
			mockBlockRepo.On("GetIncomeResult").Return(321.84).Once()
			mockBlockRepo.On("GetIncome24h").Return(42.42).Once()
			data := minerService.GetFutureIncome()
			Convey("Then expect values:", func() {
				So(data.Code, ShouldEqual, 200)
				So(data.Data.CNY, ShouldEqual, 5.8)
				So(data.Data.BTC, ShouldEqual, 8.13)
				So(data.Data.USD, ShouldEqual, 13.21)
				So(data.Data.Income, ShouldEqual, 321)
				So(data.Data.Income1d, ShouldEqual, 42)

			})
		})
	})
}

func TestGetBillInfo(t *testing.T) {

	Convey("Test GetBillInfo miner service method", t, func() {
		repoBillInfo := repository.RepoBillInfo{
			Balance:   string("2024.62"),
			FirstPaid: float64(42.21),
			FirstTime: string("1914-07-28 00:00:00"),
			TotalPaid: float64(9999.99),
		}
		Convey("When mocking hiveosRepository", func() {
			minerService := NewMinerServiceWithRepo(mockBlockRepo, mockMinerdashRepo, mockHiveosRepo,mockRedisRepo)
			mockHiveosRepo.On("GetBillInfo", mock.AnythingOfType("string")).Return(repoBillInfo)
			billInfo := minerService.GetBillInfo("42")
			Convey("Then expect values:", func() {
				So(billInfo.Balance, ShouldEqual, "2024.62")
				So(billInfo.FirstPaid, ShouldEqual, float64(42.21))
				So(billInfo.FirstTime, ShouldEqual, "1914-07-28T00:00:00Z")
				So(billInfo.TotalPaid, ShouldEqual, float64(9999.99))
			})
		})
	})
}

func TestGetBill(t *testing.T) {

	Convey("Test GetBill miner service method", t, func() {

		db, mocksql, _ := sqlmock.New()
		defer db.Close()
		mockrows := sqlmock.NewRows([]string{"id", "paid", "status", "time", "tx_hash"}).
			AddRow(1, "firstpaid", 1, "2019-01-02 00:00:00", "firsthash").
			AddRow(2, "firstpaid", 32, "2019-01-02 00:00:00", "secondhash")
		mocksql.ExpectQuery("SELECT").WillReturnRows(mockrows)
		rows, _ := db.Query("SELECT")
		defer rows.Close()

		Convey("When mocking hiveosRepository", func() {
			minerService := NewMinerServiceWithRepo(mockBlockRepo, mockMinerdashRepo, mockHiveosRepo,mockRedisRepo)
			mockHiveosRepo.On("GetBill", mock.AnythingOfType("string")).Return(rows)
			bill := minerService.GetBill("42")
			Convey("Then expect values:", func() {
				So(bill.Code, ShouldEqual, 200)
				So(bill.Data[0].Id, ShouldEqual, 1)
				So(bill.Data[0].Paid, ShouldEqual, "firstpaid")
				So(bill.Data[0].Status, ShouldEqual, "1")
				So(bill.Data[0].TXHash, ShouldEqual, "firsthash")
				So(bill.Data[0].Time, ShouldEqual, "2019-01-02T00:00:00Z")
				So(bill.Data[1].Id, ShouldEqual, 2)
				So(bill.Data[1].Paid, ShouldEqual, "firstpaid")
				So(bill.Data[1].Status, ShouldEqual, "32")
				So(bill.Data[1].TXHash, ShouldEqual, "secondhash")
				So(bill.Data[1].Time, ShouldEqual, "2019-01-02T00:00:00Z")
			})
		})
	})
}

func TestGetShares(t *testing.T) {

	Convey("Test GetShares miner service method", t, func() {

		curTime := time.Now()
		min60 := curTime.Add(time.Duration(-60) * time.Minute)
		min15 := curTime.Add(time.Duration(-15) * time.Minute)
		min5 := curTime.Add(time.Duration(-5) * time.Minute)
		hashrateConf := utils.GetConfig().GetFloat64("app.config.pool.hashrate.hashrateCul") /
			utils.GetConfig().GetFloat64("app.config.pool.hashrate.hashrateCulDivider")

		shares := make([][]interface{}, 4)
		for i := range shares {
			shares[i] = make([]interface{}, 4)
		}
		shares[0][0] = min60.Format(time.RFC3339)
		shares[0][1] = json.Number("12")
		shares[0][2] = json.Number("24")
		shares[0][3] = json.Number("48")
		shares[1][0] = min15.Format(time.RFC3339)
		shares[1][1] = json.Number("32")
		shares[1][2] = json.Number("64")
		shares[1][3] = json.Number("128")
		shares[2][0] = min5.Format(time.RFC3339)
		shares[2][1] = json.Number("999")
		shares[2][2] = json.Number("32")
		shares[2][3] = json.Number("86")
		shares[3][0] = curTime.Format(time.RFC3339)
		shares[3][1] = json.Number("53")
		shares[3][2] = json.Number("2354")
		shares[3][3] = json.Number("53332")
		sharesRow := models.Row{Values: shares}

		hashrate := make([][]interface{}, 3)
		for i := range hashrate {
			hashrate[i] = make([]interface{}, 2)
		}
		hashrate[0][0] = min60.Format(time.RFC3339)
		hashrate[0][1] = json.Number("65")
		hashrate[1][0] = min15.Format(time.RFC3339)
		hashrate[1][1] = json.Number("4")
		hashrate[2][0] = curTime.Format(time.RFC3339)
		hashrate[2][1] = json.Number("7")
		hashrateRow := models.Row{Values: hashrate}

		Convey("When mocking hiveosRepository", func() {
			minerService := NewMinerServiceWithRepo(mockBlockRepo, mockMinerdashRepo, mockHiveosRepo,mockRedisRepo)
			mockMinerdashRepo.On("GetShares", mock.AnythingOfType("string")).
				Return(sharesRow)
			mockMinerdashRepo.On("GetLocalHashrateResult", mock.AnythingOfType("string")).
				Return(hashrateRow)
			result := minerService.GetShares("42", "")
			Convey("Then expect values:", func() {
				So(result.Code, ShouldEqual, 200)
				So(result.Data, ShouldHaveLength, 2)
				So(result.Data[0].Time, ShouldEqual, min15.Format(time.RFC3339))
				So(result.Data[0].InvalidShares, ShouldEqual, 32)
				So(result.Data[0].ValidShares, ShouldEqual, 64)
				So(result.Data[0].StaleShares, ShouldEqual, 128)
				So(result.Data[0].LocalHashrate, ShouldEqual, 4)
				So(result.Data[0].Hashrate, ShouldEqual, math.Round(result.Data[0].ValidShares * hashrateConf))
				So(result.Data[0].MeanHashrate, ShouldEqual, math.Round(result.Data[0].Hashrate / 1 * 100) / 100)
				So(result.Data[1].Time, ShouldEqual, min5.Format(time.RFC3339))
				So(result.Data[1].InvalidShares, ShouldEqual, 999)
				So(result.Data[1].ValidShares, ShouldEqual, 32)
				So(result.Data[1].StaleShares, ShouldEqual, 86)
				So(result.Data[1].LocalHashrate, ShouldEqual, 0)
				So(result.Data[1].Hashrate, ShouldEqual, math.Round(result.Data[1].ValidShares * hashrateConf))
				So(result.Data[1].MeanHashrate, ShouldEqual, math.Round(
					(result.Data[0].Hashrate + result.Data[1].Hashrate) / 2 * 100) / 100)
				// set MeanHashrate precision to 2: DONE
			})
		})
	})
}