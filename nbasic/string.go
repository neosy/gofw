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
func StringToType(value string, toType reflect.Kind) (interface{}, error) {
	switch toType {
	case reflect.String:
		return value, nil
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return intValue, nil
	case reflect.Uint:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return uint(intValue), nil
	case reflect.Int8:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return int8(intValue), nil
	case reflect.Uint8:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return uint8(intValue), nil
	case reflect.Int16:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return int16(intValue), nil
	case reflect.Uint16:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return uint16(intValue), nil
	case reflect.Int32:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return int32(intValue), nil
	case reflect.Uint32:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return uint32(intValue), nil
	case reflect.Int64:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return int64(intValue), nil
	case reflect.Uint64:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return uint64(intValue), nil
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
