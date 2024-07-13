package nbasic

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
