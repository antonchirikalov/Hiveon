package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	. "hiveon-api/utils"
)

type IBlockRepository interface {
	GetIncome24h() float64
	GetIncome7d() float64
	GetIncomeResult() float64
	GetIncomeHistory() *sql.Rows
}

type BlockRepository struct {
	client *sql.DB
}

func GetBlockRepositoryClient() *sql.DB {

	username := GetConfig().GetString("sequelize3.username")
	password := GetConfig().GetString("sequelize3.password")
	host := GetConfig().GetString("sequelize3.host")
	database := GetConfig().GetString("sequelize3.database")

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+")/"+database)

	if err != nil {
		log.Error(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
	}

	return db
}

func NewBlockRepository() IBlockRepository {
	return &BlockRepository{client: GetBlockRepositoryClient()}
}

func (m *BlockRepository) queryFloatSingle(query string) float64 {
	var result float64
	row := m.client.QueryRow(query)
	row.Scan(&result)
	return result
}

func (m *BlockRepository) GetIncome24h() float64 {
	sql := "SELECT avg(expected_earning) * 2160/100000000.0 as income24h" +
		" FROM expected_earning_result where end_time >= (UNIX_TIMESTAMP() -(24 * 3600))"
	return m.queryFloatSingle(sql)
}

func (m *BlockRepository) GetIncome7d() float64 {
	sql := "SELECT avg(expected_earning) * 2160/100000000.0 as income24h" +
		" FROM expected_earning_result where end_time >= (UNIX_TIMESTAMP() -(7 * 24 * 3600))"
	return m.queryFloatSingle(sql)
}

func (m *BlockRepository) GetIncomeResult() float64 {
	sql := "SELECT expected_earning * 2160/100000000.0 as income FROM expected_earning_result where end_time >= UNIX_TIMESTAMP() limit 1"
	return m.queryFloatSingle(sql)
}

func (m *BlockRepository) GetIncomeHistory() *sql.Rows {
	sql := "select start_time,expected_earning as income from expected_earning_result where" +
		" start_time >= (UNIX_TIMESTAMP() -(24 * 3600))"
	rows, err := m.client.Query(sql)
	if err != nil {
		log.Error(err)
	}
	return rows
}
