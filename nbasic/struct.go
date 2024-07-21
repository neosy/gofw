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

// Преобразование структуры в Map
func StructToMap(data interface{}) (dataMap map[string]interface{}, err error) {
	dataJ, err := StructToJSON(data)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataJ, &dataMap) // Convert to a map
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
