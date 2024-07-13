package nmongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Главный класс для Mongo
type NMongo struct {
	address  string
	port     int
	user     string
	password string

	client *mongo.Client

	Repo *Repository
}

// Создание объекта NMongo
func New(addr string, port int, user string, password string, dbName string, collectionNames interface{}) *NMongo {
	if user == "" {
		log.Panic("Mongo user not defined")
	}

	nmongo := &NMongo{
		address:  addr,
		port:     port,
		user:     user,
		password: password,
	}

	nmongo.Repo = &Repository{}
	nmongo.Repo.Name = dbName
	nmongo.Repo.nmongo = nmongo
	nmongo.Repo.Collections.nmongo = nmongo

	nmongo.Repo.addCollections(collectionNames)

	return nmongo
}

// Соединение с Mongo
func (nmongo *NMongo) Connect(ctx context.Context) error {
	dbName := nmongo.Repo.Name

	mongoDBAddress := fmt.Sprintf("mongodb://%v:%v", nmongo.address, nmongo.port)
	clientOptions := options.Client().ApplyURI(mongoDBAddress).SetAuth(options.Credential{
		Username:   nmongo.user,
		Password:   nmongo.password,
		AuthSource: nmongo.Repo.Name,
	})

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(
			"Cannot connect to MongoDB",
			fmt.Sprint("mongoDBAddress", mongoDBAddress),
			fmt.Sprint("repoUser", nmongo.user),
			fmt.Sprint("repoName", nmongo.Repo.Name),
			fmt.Sprint(err),
		)

		return err
	}

	nmongo.client = client
	nmongo.Repo.mongoDB = client.Database(dbName)

	nmongo.Repo.createCollections(ctx)

	for c := nmongo.Repo.Collections.first; c != nil; c = c.next {
		c.mongoCollection = client.Database(dbName).Collection(c.Name)
	}

	return err
}
