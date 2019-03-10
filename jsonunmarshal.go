package gohelper

import "encoding/json"

// JSONUnmarshalByte - target SHOULD be a pointer
func JSONUnmarshalByte(byteValue []byte, target *interface{}) error {
	return json.Unmarshal(byteValue, target)
}

// JSONUnmarshalString - target SHOULD be a pointer
func JSONUnmarshalString(stringValue string, target *interface{}) error {
	byteValue := []byte(stringValue)
	return JSONUnmarshalByte(byteValue, target)
}

// JSONUnmarshalInterface - target SHOULD be a pointer
func JSONUnmarshalInterface(interfaceValue interface{}, target *interface{}) error {
	byteValue, err := json.Marshal(interfaceValue)
	if err != nil {
		return err
	}
	return JSONUnmarshalByte(byteValue, target)
}

// JSONUnmarshalByteAndValidate - target SHOULD be a pointer
//func JSONUnmarshalByteAndValidate(byteValue []byte, target *interface{}) error {
//	err := JSONUnmarshalByte(byteValue, target)
//	if err != nil {
//		return err
//	}
//	return ValidateStruct(target)
//}

// JSONUnmarshalStringAndValidate - target SHOULD be a pointer
//func JSONUnmarshalStringAndValidate(jsonString string, target *interface{}) error {
//	err := JSONUnmarshalString(jsonString, target)
//	if err != nil {
//		return err
//	}
//	return ValidateStruct(target)
//}

// JSONUnmarshalInterfaceAndValidate - target SHOULD be a pointer
//func JSONUnmarshalInterfaceAndValidate(interfaceValue interface{}, target *interface{}) error {
//	err := JSONUnmarshalInterface(interfaceValue, target)
//	if err != nil {
//		return err
//	}
//	return ValidateStruct(target)
//}
