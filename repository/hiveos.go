package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	. "hiveon-api/utils"
)

type IHiveosRepository interface {
	GetBlock24NotUnckle() int
	GetBlock24Uncle() int
	GetBillInfo(walletId string) RepoBillInfo
	GetBill(walletId string) *sql.Rows
	GetBalance(walletId string) float64
}

type HiveosRepository struct {
	hiveClient *sql.DB
}

type RepoBillInfo struct {
	Balance   string
	FirstPaid float64
	FirstTime string
	TotalPaid float64
}

func GetHaveonClient() *sql.DB {

	username := GetConfig().GetString("sequelize2.username")
	password := GetConfig().GetString("sequelize2.password")
	host := GetConfig().GetString("sequelize2.host")
	database := GetConfig().GetString("sequelize2.database")

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+")/"+database)

	//defer db.Close()

	if err != nil {
		log.Error(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
	}

	return db
}
func NewHiveosRepository() IHiveosRepository {
	return &HiveosRepository{hiveClient: GetHaveonClient()}
}

func (repo *HiveosRepository) queryIntSingle(query string) int {
	var result int
	row := repo.hiveClient.QueryRow(query)
	row.Scan(&result)
	return result
}

func (repo *HiveosRepository) GetBlock24NotUnckle() int {
	pgOneDay := GetConfig().GetString("app.config.pool.pgOneDay")
	sql := "select count(id) as count from blocks as b where is_uncle=0 and b.block_ts > UNIX_TIMESTAMP(DATE_SUB(now(), interval " +
		pgOneDay + ")) * 1000"
	return repo.queryIntSingle(sql)
}

func (repo *HiveosRepository) GetBlock24Uncle() int {
	pgOneDay := GetConfig().GetString("app.config.pool.pgOneDay")
	sql := "select count(id) as count from blocks as b where is_uncle=1 and b.block_ts > UNIX_TIMESTAMP(DATE_SUB(now(), interval " +
		pgOneDay + ")) * 1000"
	return repo.queryIntSingle(sql)
}

func (repo *HiveosRepository) GetBill(walletId string) *sql.Rows {
	rows, err := repo.hiveClient.Query("select p.id, pd.paid, p.status, p.create_ts, p.tx_hash  from payment_details pd " +
		"inner join payments p on p.id = pd.id where pd.miner_wallet = ? order by pd.id desc limit 30", walletId)
	if err != nil {
		log.Error(err)
	}

	return rows
}

func (repo *HiveosRepository) GetBillInfo(walletId string) RepoBillInfo {
	var totalPaid float64

	err := repo.hiveClient.QueryRow("select sum(paid) as totalPaid from payment_details where miner_wallet = ?", walletId).Scan(&totalPaid)
	if err != nil {
		log.Error(err)
	}

	var firstPaid, paymentId float64
	var firstTime, balance string
	err = repo.hiveClient.QueryRow("select paid,payment_id from payment_details where miner_wallet = ? order by id limit 1",
		walletId).Scan(&firstPaid, &paymentId)
	if err != nil {
		log.Error(err)
	}

	repo.hiveClient.QueryRow("select create_ts from payments where id=?", paymentId).Scan(&firstTime)
	if err != nil {
		log.Error(err)
	}

	repo.hiveClient.QueryRow("select balance from deposits where miner_wallet =?", walletId).Scan(&balance)
	if err != nil {
		log.Error(err)
	}

	return RepoBillInfo{Balance: balance, FirstPaid: firstPaid, FirstTime: firstTime, TotalPaid: totalPaid}

}

func (repo *HiveosRepository) GetBalance(walletId string) float64 {
	var res float64
	err := repo.hiveClient.QueryRow("select balance from deposits where miner_wallet = ?", walletId).Scan(&res)
	if err != nil {
		log.Error(err)
	}
	return res
}
