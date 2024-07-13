package nmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Репозитарий для Mongo
type Repository struct {
	mongoDB *mongo.Database

	nmongo      *NMongo
	Name        string
	Collections RepositoryCollections
}

// Создание коллекций в Mongo
func (repo *Repository) createCollections(ctx context.Context) {
	for c := repo.Collections.first; c != nil; c = c.next {
		_ = repo.Collections.createCollection(ctx, c.Name)
	}
}
