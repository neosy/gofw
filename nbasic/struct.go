package nbasic

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// Преобразование структуры в JSON
func StructToJSON(data interface{}) (dataJ []byte, err error) {
	dataJ, err = json.Marshal(data)

	if err != nil {
		log.Printf("%v: %v", ErrConvertToJSON.Error(), err)
		return
	}

	return
}

// Преобразование структуры в Map, состоящую из строк
func StructToMapString(data interface{}) (map[string]string, error) {
	dataMap := make(map[string]string, 0)
	val := reflect.ValueOf(data)
	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %v", typ)
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		value := val.Field(i)

		if value.Kind() == reflect.Struct {
			log.Println(
				"wrong body paloyad structure. Do not support nested",
			)
			return nil, ErrInternalError
		}

		tag := field.Tag.Get("json")
		key := ""
		if tag != "" {
			key = tag
		} else {
			key = field.Name
		}
		dataMap[key] = fmt.Sprint(value)
	}

	return dataMap, nil
}

// Преобразование структуры в Map
func StructToMap(data interface{}) map[string]interface{} {
	dataMap := make(map[string]interface{})

	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		fieldValue := dataValue.Field(i).Interface()
		dataMap[fieldName] = fieldValue
	}

	return dataMap
}

// Преобразование структуры в Map
func StructToMapV2(data interface{}) (dataMap map[string]interface{}, err error) {
	dataJ, err := StructToJSON(data)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataJ, &dataMap) // Convert to a map
	return
}
