package nmongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Преобразование JSON в bson.M
func JSONToBSON(json []byte) (bson.M, error) {
	var bsonData bson.M

	err := bson.UnmarshalExtJSON([]byte(json), true, &bsonData)

	return bsonData, err
}

// Преобразование Struc в bson.M
func StructToBSON(data interface{}) (dataBSON []byte, err error) {
	dataBSON, err = bson.Marshal(data)

	return dataBSON, err
}
