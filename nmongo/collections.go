package nmongo

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

// Коллекции Mongo
type RepositoryCollections struct {
	number      int
	first, last *RepositoryCollection
	nmongo      *NMongo
}

// Колличество коллекций
func (objs *RepositoryCollections) Number() int {
	return objs.number
}

// Добавление коллекции в репозитарий
func (objs *RepositoryCollections) add(obj *RepositoryCollection) *RepositoryCollection {
	var ret *RepositoryCollection

	objs.number++

	if obj != nil {
		if objs.first == nil {
			obj.prev = nil
			objs.first = obj
		} else {
			obj.prev = objs.last
			objs.last.next = obj
		}
		obj.next = nil
		objs.last = obj
		ret = objs.last
	}

	// Инициализация ссылок
	ret.initLinks(objs.nmongo)

	return ret
}

// Добавление коллекции в репозитарий по имени
func (objs *RepositoryCollections) AddByName(name string) *RepositoryCollection {
	return objs.add(&RepositoryCollection{Name: name})
}

// Получение коллекции из репозитария по имени
func (objs *RepositoryCollections) Collection(name string) *RepositoryCollection {
	cur := objs.first

	for cur != nil {
		if cur.Name == name {
			break
		}
		cur = cur.next
	}

	return cur
}

// Получение коллекции из репозитария по номеру
func (objs *RepositoryCollections) CollectionByNum(num int) *RepositoryCollection {
	cur := objs.first

	for i := 0; cur != nil; i++ {
		if i == num {
			break
		}
		cur = cur.next
	}

	return cur
}

// Инициализация ссылок
func (c *RepositoryCollection) initLinks(nmongo *NMongo) {
	c.nmongo = nmongo
	c.repo = nmongo.Repo
}

func (repo *Repository) addCollections(names interface{}) {
	values := reflect.ValueOf(names)

	if values.Kind() == reflect.Ptr {
		values = values.Elem() // Получаем значение, на которое указывает указатель
	} else {
		log.Println("Input is not a pointer to a struct.")
		return
	}

	if values.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct.")
		return
	}

	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		if field.Kind() == reflect.String {
			repo.Collections.AddByName(field.String())
		}
	}
}

// Создание коллекции в Mongo
func (collections *RepositoryCollections) createCollection(ctx context.Context, collectionName string) error {
	mongoClient := collections.nmongo.client
	dbName := collections.nmongo.Repo.Name
	mongoDB := mongoClient.Database(dbName)

	err := mongoDB.CreateCollection(ctx, collectionName)
	if err != nil {
		if mongoErr, ok := err.(mongo.CommandError); ok {
			if mongoErr.Code != 48 {
				log.Fatalln(
					"cannot create MongoCollection",
					fmt.Sprintf("dbName: %v", dbName),
					fmt.Sprintf("dbCollectionName: %v", collectionName),
					fmt.Sprint(err),
				)

				return mongoErr
			} else {
				log.Println(
					"MongoCollection already exists",
					fmt.Sprintf("dbName: %v", dbName),
					fmt.Sprintf("dbCollectionName: %v", collectionName),
				)
			}
		}
	} else {
		collections.Collection(collectionName).mongoCollection = mongoDB.Collection(collectionName)
	}

	return nil
}
