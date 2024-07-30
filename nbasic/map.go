package nbasic

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Получения среза ключей из Map
func MapToKeys(dataMap map[string]interface{}) []string {
	keys := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}
	return keys
}

// Получения среза значений из Map
func MapToValues(dataMap map[string]interface{}) []interface{} {
	values := make([]interface{}, 0, len(dataMap))
	for _, v := range dataMap {
		values = append(values, v)
	}
	return values
}

// Преобразование Map в структуру
func MapStringToJSON(dataMap map[string]string) ([]byte, error) {
	dataJSON, err := json.Marshal(dataMap)

	if err != nil {
		fmt.Println(ErrConvertToJSON.Error())
	}

	return dataJSON, err
}

// Преобразование Map в структуру
func MapStringToStruct(dataMap map[string]string, data interface{}) error {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}

	for i := 0; i < dataType.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		fieldValue, ok := dataMap[fieldName]
		if ok {
			field := dataValue.Field(i)
			if field.CanSet() {
				value, err := StringToType(fieldValue, dataType.Field(i).Type)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(value))
			}
		}
	}

	return nil
}

// Преобразование Map в структуру
func MapStringToStructV2(dataMap map[string]string, data interface{}) error {
	dataJSON, err := MapStringToJSON(dataMap)

	if err != nil {
		return err
	}

	err = json.Unmarshal(dataJSON, data)
	if err != nil {
		fmt.Println(ErrConvertJSONToStruct.Error())
	}

	return err
}
