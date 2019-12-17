package gohelper

import "encoding/json"

// JSONUnmarshalByte - target SHOULD be a pointer
func JSONUnmarshalByte(byteValue []byte, target interface{}) error {
	return json.Unmarshal(byteValue, target)
}

// JSONUnmarshalString - target SHOULD be a pointer
func JSONUnmarshalString(stringValue string, target interface{}) error {
	byteValue := []byte(stringValue)
	return JSONUnmarshalByte(byteValue, target)
}

// JSONUnmarshalInterface - target SHOULD be a pointer
func JSONUnmarshalInterface(interfaceValue interface{}, target interface{}) error {
	byteValue, err := json.Marshal(interfaceValue)
	if err != nil {
		return err
	}
	return JSONUnmarshalByte(byteValue, target)
}
