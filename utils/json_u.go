package utils

import "encoding/json"

func AnyToMap(data any) (res map[string]interface{}, err error) {
	a, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(a, &res)
	return
}

func MapToAny[K any](m map[string]interface{}) (K, error) {
	data, err := json.Marshal(m)
	var result K
	if err == nil {
		err = json.Unmarshal(data, &result)
	}
	return result, err
}

func AnyToJson(data any) (result string, err error) {
	a, err := json.Marshal(data)
	if err != nil {
		return
	}
	result = string(a)
	return
}

func JsonToAny[K any](jsonStr string) (K, error) {
	var result K
	err := json.Unmarshal([]byte(jsonStr), &result)
	return result, err
}

func AnyToByte(data any) ([]byte, error) {
	return json.Marshal(data)
}

func ByteToAny[K any](data []byte) (K, error) {
	var result K
	err := json.Unmarshal(data, &result)
	return result, err
}

func LoadFromListMap[K any](m []map[string]interface{}) (result []K, err error) {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, &result)
	}

	return
}
