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
func StructToMapStringInterface(data interface{}) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})
	err := structToMapStringInterfaceOne("", reflect.ValueOf(data), dataMap)

	return dataMap, err
}

func structToMapStringInterfaceOne(prefix string, dataValue reflect.Value, dataMap map[string]interface{}) error {
	var err error

	switch dataValue.Kind() {
	case reflect.Ptr:
		err = structToMapStringInterfaceOne(prefix, dataValue.Elem(), dataMap)
	case reflect.Struct:
		dataType := dataValue.Type()
		for i := 0; i < dataValue.NumField(); i++ {
			fieldName := dataType.Field(i).Name
			fieldValue := dataValue.Field(i)
			if prefix != "" {
				fieldName = prefix + "." + fieldName
			}
			err = structToMapStringInterfaceOne(fieldName, fieldValue, dataMap)
			if err != nil {
				break
			}
		}
	default:
		dataType := dataValue.Type()
		baseType := ReflectKindToType(dataType.Kind())
		var fieldValue any
		if baseType != nil {
			baseValue := reflect.New(baseType).Elem()
			baseValue.Set(dataValue.Convert(baseType))
			fieldValue = baseValue.Interface()
		} else {
			//fieldValue = dataValue.Field(i).Elem().Interface()
			err = ErrConvertStructToMap
			fmt.Println(err.Error())
			return err
		}
		dataMap[MapNameCorrect(prefix)] = fieldValue
	}

	return err
}

// Преобразование структуры в Map через JSON
func StructToMapStringInterfaceV2(data interface{}) (dataMap map[string]interface{}, err error) {
	dataJ, err := StructToJSON(data)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataJ, &dataMap) // Convert to a map
	return
}
