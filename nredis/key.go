package nredis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	KeySeparatorDef  string        = ":"
	KeyExpirationDef time.Duration = 0
	KeyLogEnabledDef bool          = true
)

type Key struct {
	value      string
	separator  string
	expiration time.Duration
	logEnabled bool

	client *redis.Client
}

// Создание объетка Key
func NewKey(separator string, part ...string) *Key {
	key := &Key{
		value:      "",
		separator:  separator,
		expiration: KeyExpirationDef,
		logEnabled: KeyLogEnabledDef,
	}

	key.value = KeyGen(separator, part...)

	return key
}

// Создание текстового ключа из состовляющих
func KeyGen(separator string, part ...string) string {
	parts := make([]string, len(part))
	key := ""

	for i := range parts {
		key += parts[i]
		if i+1 != len(parts) {
			key += separator
		}
	}

	return key
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
	exists_value, err := nKey.client.Exists(ctx, nKey.value).Result()

	if exists_value == 1 {
		exists = true
	}

	return exists, err
}

// Вставка как Ключ -> Значение
func (nKey *Key) Set(ctx context.Context, value interface{}) error {
	err := nKey.client.Set(ctx, nKey.value, value, nKey.expiration).Err()
	if nKey.logEnabled && err != nil {
		log.Println(ErrRecordInserting.Error())
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

// Чтение структуры по ключу
func (nKey *Key) GetStruct(ctx context.Context, key string, data interface{}) error {
	value, err := nKey.client.Get(ctx, key).Result()

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

// Вставка структуры как Ключ -> Значение
func (nKey *Key) SetStruct(ctx context.Context, value interface{}) error {
	valueJSON, err := json.Marshal(value)
	if err != nil {
		if nKey.logEnabled {
			log.Println(ErrCannotConvertToJSON.Error(), err)
		}
		return err
	}

	err = nKey.client.Set(ctx, nKey.value, valueJSON, nKey.expiration).Err()
	if err != nil {

		if nKey.logEnabled {
			log.Println(ErrRecordInserting.Error())
		}
		return err
	}

	return err
}
