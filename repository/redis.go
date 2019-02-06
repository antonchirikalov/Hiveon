package repository

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	. "hiveon-api/utils"
)

type IRedisRepository interface {
	GetLatestWorker(walletId string)  map[string]string
}

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository() IRedisRepository {
	return &RedisRepository{redisClient: GetRedisClient()}
}

func GetRedisClient() *redis.Client {
	host := GetConfig().GetString("redis.host")
	port := GetConfig().GetString("redis.port")
	password := GetConfig().GetString("redis.password")
	db := GetConfig().GetInt("redis.db")
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error(err)
	}

	return client
}

func (repo *RedisRepository) GetLatestWorker(walletId string) map[string]string {
	result, err := repo.redisClient.HGetAll("last-update:" + walletId).Result()
	if err != nil {
		log.Error(err)
	}
	return  result
}
