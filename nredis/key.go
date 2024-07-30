package nredis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/neosy/gofw/nbasic"
)

const (
	KeySeparatorDef  string        = ":"
	KeyExpirationDef time.Duration = 0
	KeyLogEnabledDef bool          = true
)

type Key struct {
	name       string
	separator  string
	expiration time.Duration
	logEnabled bool

	client *redis.Client
}

// Создание объетка Key
func NewKey(separator string, part ...string) *Key {
	key := &Key{
		name:       "",
		separator:  separator,
		expiration: KeyExpirationDef,
		logEnabled: KeyLogEnabledDef,
	}

	key.name = KeyGen(separator, part...)

	return key
}

// Создание текстового ключа из состовляющих
func KeyGen(separator string, part ...string) (key string) {
	for i := range part {
		key += part[i]
		if i+1 != len(part) {
			key += separator
		}
	}

	return
}

func (key *Key) LogEnable() {
	key.logEnabled = true
}

func (key *Key) LogDisable() {
	key.logEnabled = true
}

func (key *Key) ExpirationSet(expiration time.Duration) {
	key.expiration = expiration
}

// Инициализация redis.client
func (key *Key) ClientSet(client *redis.Client) {
	key.client = client
}

// Создание объетка Key
func CreateKey(part ...string) *Key {
	key := NewKey(KeySeparatorDef, part...)

	return key
}

// Проверка существования ключа
func (nKey *Key) Exists(ctx context.Context) (bool, error) {
	var exists bool
	exists_value, err := nKey.client.Exists(ctx, nKey.name).Result()

	if exists_value == 1 {
		exists = true
	}

	return exists, err
}

// Вставка как Ключ -> Значение
func (nKey *Key) Set(ctx context.Context, data interface{}) error {
	err := nKey.client.Set(ctx, nKey.name, data, nKey.expiration).Err()
	if nKey.logEnabled && err != nil {
		log.Println(ErrRecordInserting.Error())
	}

	return err
}

// Вставка как Ключ -> Значение
func (nKey *Key) HSet(ctx context.Context, data interface{}) error {
	err := nKey.client.HSet(ctx, nKey.name, data, nKey.expiration).Err()
	if nKey.logEnabled && err != nil {
		log.Println(ErrRecordInserting.Error(), err)
		return err
	}

	_, err = nKey.client.Expire(ctx, nKey.name, nKey.expiration).Result()
	if err != nil {
		log.Println("oшибка при установке срока действия", err)
	}

	return err
}

// Чтение значения по ключу
func (nKey *Key) Get(ctx context.Context, key string) (string, error) {
	value, err := nKey.client.Get(ctx, key).Result()

	if nKey.logEnabled && err != nil {
		log.Println(ErrRecordSearching.Error())
	}

	return value, err
}

// Чтение значения по ключу
func (nKey *Key) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	dataMap, err := nKey.client.HGetAll(ctx, key).Result()

	if nKey.logEnabled && err != nil {
		log.Println(ErrRecordSearching.Error())
		return nil, err
	}

	return dataMap, err
}

// Чтение структуры по ключу. Хранение в виде JSON
func (nKey *Key) GetStructJSON(ctx context.Context, data interface{}) error {
	value, err := nKey.client.Get(ctx, nKey.name).Result()

	if nKey.logEnabled && err != nil {
		log.Println(ErrRecordSearching.Error())
		return err
	}

	err = json.Unmarshal([]byte(value), data)

	if nKey.logEnabled && err != nil {
		log.Println(ErrCannotConvertFromJSON.Error())
		return err
	}

	return err
}

// Чтение структуры по ключу
func (nKey *Key) HGetStruct(ctx context.Context, data interface{}) error {
	dataMap, err := nKey.HGetAll(ctx, nKey.name)

	if err != nil {
		return err
	}

	err = nbasic.MapStringToStruct(dataMap, data)

	return err
}

// Вставка структуры как Ключ -> Значение. Хранение в виде JSON
func (nKey *Key) SetStructJSON(ctx context.Context, value interface{}) error {
	valueJSON, err := json.Marshal(value)
	if err != nil {
		if nKey.logEnabled {
			log.Println(ErrCannotConvertToJSON.Error(), err)
		}
		return err
	}

	err = nKey.client.Set(ctx, nKey.name, valueJSON, nKey.expiration).Err()
	if err != nil {

		if nKey.logEnabled {
			log.Println(ErrRecordInserting.Error())
		}
		return err
	}

	return err
}

// Вставка структуры как Ключ -> Значение
func (nKey *Key) HSetStruct(ctx context.Context, data interface{}) error {
	dataMap := nbasic.StructToMap(data)

	err := nKey.client.HSet(ctx, nKey.name, dataMap).Err()
	if err != nil {

		if nKey.logEnabled {
			log.Println(ErrRecordInserting.Error())
		}
		return err
	}

	return err
}
