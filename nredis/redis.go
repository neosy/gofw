package nredis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Главный класс для Mongo
type NRedis struct {
	address  string
	port     int
	password string
	dbIdx    int

	client *redis.Client
}

func New(addr string, port int, password string, dbIdx int) *NRedis {
	nredis := &NRedis{
		address:  addr,
		port:     port,
		password: password,
		dbIdx:    dbIdx,
	}

	return nredis
}

// Соединение с Redis
func (nredis *NRedis) Connect(ctx context.Context) error {
	redisAddress := fmt.Sprintf("%v:%v", nredis.address, nredis.port)

	nredis.client = redis.NewClient(&redis.Options{
		Addr:     redisAddress,    // Адрес Redis-сервера
		Password: nredis.password, // Пароль (если есть)
		DB:       nredis.dbIdx,    // Индекс базы данных
	})

	// Проверка состояния соединения
	err := nredis.client.Ping(ctx).Err()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}

	return err
}
