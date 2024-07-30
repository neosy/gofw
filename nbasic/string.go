package nbasic

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToLowerFirst(s string) (ret string) {
	ret = s

	if len(ret) > 0 {
		ret = strings.ToLower(ret[:1]) + ret[1:]
	}

	return
}

// Преобразование строки в заданный тип данных
func StringToType(value string, toType reflect.Type) (interface{}, error) {
	switch toType.Kind() {
	case reflect.String:
		return value, nil
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return intValue, nil
	case reflect.Int8:
		int8Value, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8Value, nil
	case reflect.Int16:
		int16Value, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16Value, nil
	case reflect.Int32:
		int32Value, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32Value, nil
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return floatValue, nil
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		return boolValue, nil
	// Добавьте обработку других типов
	default:
		err := fmt.Errorf("unsupported target type: %v", toType)
		fmt.Println(err.Error())
		return nil, err
	}
}
