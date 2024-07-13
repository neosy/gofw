package nmongo

import "errors"

// Records errors
var ErrRecordAlreadyExists = errors.New("record already exists")
var ErrRecordNotFound = errors.New("data not found")
var ErrRecordCannotDecodeData = errors.New("cannot decode cursor data")
var ErrRecordCannotReadData = errors.New("cannot read data. Cursor error")
var ErrRecordSearching = errors.New("error searching for a record")
var ErrRecordUpdating = errors.New("error updating the record")
var ErrRecordInserting = errors.New("error inserting the record")

// DB Mongo Errors
var ErrCannotConnectToMongoDb = errors.New("cannot connect to mongodb")
var ErrMongoDbUnavailable = errors.New("unable to check connection to mongodb")
var ErrDBNameOrCollectionNotFound = errors.New("dbName or MongoCollectionName not defined")
var ErrUnableToCreateCollection = errors.New("cannot create MongoCollection name")
