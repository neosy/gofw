package nredis

import "errors"

// Records errors
var ErrRecordAlreadyExists = errors.New("record already exists")
var ErrRecordNotFound = errors.New("data not found")
var ErrRecordSearching = errors.New("error searching for a record")
var ErrRecordUpdating = errors.New("error updating the record")
var ErrRecordInserting = errors.New("error inserting the record")

// DB Mongo Errors
var ErrCannotConnectToMongoDb = errors.New("cannot connect to Redis")
