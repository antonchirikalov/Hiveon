package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	. "hiveon-api/model"
	. "hiveon-api/utils"
)

type IUserRepository interface {
	GetUserWallets(fid int) []UserWallet
	SaveUserWallet(user UserWallet)
}

type UserRepository struct {
	client *gorm.DB
}

func GetUserRepositoryClient() *gorm.DB {

	conn := fmt.Sprintf("host=%s  port=%s sslmode=disable user=%s dbname=%s password=%s",
		GetConfig().GetString("postgres.host"),
		GetConfig().GetString("postgres.port"),
		GetConfig().GetString("postgres.user"),
		GetConfig().GetString("postgres.dbName"),
		GetConfig().GetString("postgres.password"),
	)
	var err error
	client, err := gorm.Open("postgres", conn)
	client.LogMode(true)

	if err != nil {
		log.Error(err)
	}

	return client

}

func NewUserRepository() IUserRepository {
	return &UserRepository{client: GetUserRepositoryClient()}
}

func (g *UserRepository) GetUserWallets(fid int) []UserWallet {
	var userWallets []UserWallet
	g.client.Where("fid = ?", fid).Find(&userWallets)
	return userWallets
}

func (g *UserRepository) SaveUserWallet(user UserWallet) {
	g.client.Save(&user)
	return
}
