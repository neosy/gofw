package nmongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Коллекция Mongo
type RepositoryCollection struct {
	nmongo     *NMongo
	repo       *Repository
	prev, next *RepositoryCollection

	mongoCollection *mongo.Collection

	Name string
}

type Cursor struct {
	ctx         context.Context
	cancel      context.CancelFunc
	mongoCursor *mongo.Cursor
}

// Проверка существует ли запись по фильтру
func (collection *RepositoryCollection) Exists(ctx context.Context, filter interface{}) (ret bool, err error) {
	var result bson.M
	if err = collection.mongoCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		} else {
			log.Println(
				ErrRecordSearching.Error(),
				err,
			)
		}
		return false, err
	}
	return true, err
}

// Поиск записи по Id
// Пример:
// data = &<struct>{}
// err = Collections.Collection(<name>).FindId(recId, data)
// data - значение записи вернутся сюда
func (collection *RepositoryCollection) FindId(ctx context.Context, id string, data interface{}) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return collection.FindOne(ctx, bson.M{"_id": objectId}, data)
}

// Поиск одной записи по фильтру
func (collection *RepositoryCollection) FindOne(ctx context.Context, filter interface{}, data interface{}) (err error) {
	cursor := collection.mongoCollection.FindOne(ctx, filter)
	if err = cursor.Decode(data); err != nil {
		log.Println(
			ErrRecordCannotDecodeData.Error(),
			err,
		)
		return err
	}

	if err = cursor.Err(); err != nil {
		log.Println(
			ErrRecordCannotReadData.Error(),
			err,
		)
		return err
	}

	return err
}

// Следующая запись
func (cursor *Cursor) Next() bool {
	return cursor.mongoCursor.Next(cursor.ctx)
}

// Следующая запись
func (cursor *Cursor) Close() {
	defer cursor.cancel()
	defer cursor.mongoCursor.Close(cursor.ctx)
}

// Преобразование
func (cursor *Cursor) Decode(data interface{}) (err error) {
	err = cursor.mongoCursor.Decode(data)

	if err != nil {
		log.Println(
			ErrRecordCannotDecodeData.Error(),
			err,
		)
	}

	return
}

// Поиск записей по фильтру
//
//	  Пример:
//		 collection := nmongo.Repo.Collections.Collection(ms.collectionNames.Transactions)
//		 cursor, err := collection.Find(ctx, filter)
//		 if err != nil {
//	        return
//	   	 }
//		 defer cursor.Close()
//
//		 filter := bson.M{"txnId": data.TxnId}
//		 record := &entity.Transaction{}
//		 for cursor.Next() {
//		 	cursor.Decode(record)
//		 	fmt.Println(record)
//		 }
func (collection *RepositoryCollection) Find(ctx context.Context, filter interface{}) (cursor *Cursor, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(30)*time.Second)

	cursor = &Cursor{}

	cursor.ctx = ctx
	cursor.cancel = cancel
	cursor.mongoCursor, err = collection.mongoCollection.Find(ctx, filter)
	if err != nil {
		log.Println(
			ErrRecordNotFound.Error(),
			err,
		)
		return nil, err
	}

	return
}

// Выбор данных по фильтру и пребразование в структуру
//
//	 Пример:
//		filter := bson.M{}
//		filter["txnId"] = <id>
//		txnData := &entity.Transaction{}
//		datas, err := FindToStruct(ctx, &filter, txnData)
//
//		if len(datas) > 0 {
//			txnData = datas[0].(*entity.Transaction)
//		}
func (collection *RepositoryCollection) FindToStruct(ctx context.Context, filter interface{}, data interface{}) ([]interface{}, error) {
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	datas := make([]interface{}, 0)

	for cursor.Next() {
		if err = cursor.Decode(data); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	if err := cursor.mongoCursor.Err(); err != nil {
		log.Println(
			ErrRecordCannotReadData.Error(),
			err,
		)
		return nil, err
	}

	return datas, err
}

// Выбор данных по фильтру и пребразование в структуру
//
//	 Пример:
//		filter := bson.M{}
//		filter["txnId"] = <id>
//		txnData := &entity.Transaction{}
//		datas, err := FindToStruct2(ctx, &filter, txnData)
//
//		if len(datas) > 0 {
//			txnData = datas[0].(*entity.Transaction)
//		}
func (collection *RepositoryCollection) FindToStruct2(ctx context.Context, filter interface{}, data interface{}) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(30)*time.Second)
	defer cancel()

	cursor, err := collection.mongoCollection.Find(ctx, filter)

	if err != nil {
		log.Println(
			ErrRecordNotFound.Error(),
			err,
		)
		return nil, err
	}
	defer cursor.Close(ctx)

	datas := make([]interface{}, 0)
	for cursor.Next(ctx) {
		if err := cursor.Decode(data); err != nil {
			log.Println(
				ErrRecordCannotDecodeData.Error(),
				err,
			)
			return nil, err
		}

		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		log.Println(
			ErrRecordCannotReadData.Error(),
			err,
		)
		return nil, err
	}

	return datas, err
}

// Вставка записи в коллекцию
func (collection *RepositoryCollection) InsertOne(ctx context.Context, data interface{}) (recId string, err error) {
	result, err := collection.mongoCollection.InsertOne(ctx, data)
	if err != nil {
		log.Println(
			ErrRecordInserting.Error(),
			err,
		)
	}

	recId = result.InsertedID.(primitive.ObjectID).Hex()
	return recId, err
}

// Обновить одну запись в коллекции по фильтру
//
//	Пример:
//	Условие обновления (например, по полю "name")
//	filter := bson.M{"name": "John Doe"}
//	Новое значение поля "age"
//	data := bson.M{"age": 30}
//	updateResult, err := collection.UpdateOne(filter, data)
func (collection *RepositoryCollection) UpdateOne(ctx context.Context, filter interface{}, data interface{}) (*mongo.UpdateResult, error) {
	dataSetUpd := bson.M{"$set": data}

	// Выполняем обновление
	result, err := collection.mongoCollection.UpdateOne(ctx, filter, dataSetUpd)
	if err != nil {
		log.Println(
			ErrRecordUpdating.Error(),
			err,
		)
		return nil, err
	}

	return result, err
}

// Обновить несколько записей в коллекции по фильтру
//
//	Пример:
//	Условие обновления (например, по полю "name")
//	filter := bson.M{"name": "John Doe"}
//	Новое значение поля "age"
//	data := bson.M{"age": 30}
//	updateResult, err := collection.UpdateMany(filter, data)
func (collection *RepositoryCollection) UpdateMany(ctx context.Context, filter interface{}, data interface{}) (*mongo.UpdateResult, error) {
	dataSetUpd := bson.M{"$set": data}

	// Выполняем обновление
	result, err := collection.mongoCollection.UpdateMany(ctx, filter, dataSetUpd)
	if err != nil {
		log.Println(
			ErrRecordUpdating.Error(),
			err,
		)
		return nil, err
	}

	return result, err
}
