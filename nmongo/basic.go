package nmongo

import "go.mongodb.org/mongo-driver/bson"

// Преобразование JSON в bson.M
func JSONTobsonM(json []byte) (bson.M, error) {
	var bsonData bson.M

	err := bson.UnmarshalExtJSON([]byte(json), true, &bsonData)

	return bsonData, err
}
