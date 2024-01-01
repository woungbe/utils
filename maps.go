package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// MapToStruct 함수는 맵을 구조체로 변환하는 함수입니다.
// mapData: 변환할 맵 데이터
// result: 변환된 구조체 결과
// 테스트: TestMapToStruct
func MapToStruct(tag string, mapData map[string]interface{}, result interface{}) error {
	// 입력된 결과 인터페이스의 유효성을 검사
	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("result 인자는 구조체 포인터여야 합니다")
	}

	// 결과 구조체의 타입 정보를 가져옴
	resultType := resultValue.Elem().Type()

	// 맵 데이터를 구조체로 복사
	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		fieldName := field.Tag.Get(tag)
		mapValue, exists := mapData[fieldName]
		if !exists {
			continue // 맵에 필드 이름이 없는 경우 스킵
		}

		fieldValue := resultValue.Elem().Field(i)

		if fieldValue.Type() == reflect.TypeOf(mapValue) {
			mapValueReflect := reflect.ValueOf(mapValue)
			fieldValue.Set(mapValueReflect.Convert(fieldValue.Type()))
			continue
		}

		// 맵 값을 구조체 필드로 할당 (타입 변환)
		if err := assignValue(mapValue, &fieldValue); err != nil {
			return fmt.Errorf("필드 %s: %v", fieldName, err)
		}
	}

	return nil
}

// assignValue 함수는 맵 값을 구조체 필드로 할당합니다. - 직접사용하지는 않고,  MapToStruct 에 붙여서 사용함
func assignValue(mapValue interface{}, fieldValue *reflect.Value) error {
	// 구조체 필드 타입과 맵 값 타입 확인
	fieldType := fieldValue.Type()
	mapValueType := reflect.TypeOf(mapValue)

	if mapValueType == nil {
		if fieldType.String() == "string" {
			fieldValue.SetString("")
			return nil
		}

		// 문자열(string)을 int64로 변환해서 할당하는 로직 추가
		if fieldType.String() == "int64" {
			fieldValue.SetInt(0)
			return nil
		}

		// bool
		if fieldType.String() == "bool" {
			fieldValue.SetBool(false)
			return nil
		}

		// float64
		if fieldType.String() == "float64" {
			fieldValue.SetFloat(0)
			return nil
		}

	} else {
		// 맵 값 타입을 구조체 필드 타입으로 변환 가능한지 확인
		if mapValueType.ConvertibleTo(fieldType) {
			*fieldValue = reflect.ValueOf(mapValue).Convert(fieldType)
			return nil
		}

		if fieldType.Kind() == reflect.Int {
			strValue := mapValue.(string)
			val, err := Int64(strValue)
			if err != nil {
				return err
			}
			fieldValue.SetInt(val)
			return nil
		}

		// 문자열(string)을 int64로 변환해서 할당하는 로직 추가
		if fieldType.Kind() == reflect.Int64 {
			strValue := mapValue.(string)
			intValue, err := strconv.ParseInt(strValue, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(intValue)
			return nil
		}

		// bool
		if fieldType.Kind() == reflect.Bool {
			var send bool
			boolValue := mapValue.(string)

			tmp, err := Int(boolValue)
			if err != nil {
				return err
			}

			if tmp == 0 {
				send = false
			} else {
				send = true
			}
			fieldValue.SetBool(send)
			return nil
		}

		// float64
		if fieldType.Kind() == reflect.Float64 {
			strValue := mapValue.(string)
			floatval, err := Float64(strValue)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(floatval)
			return nil
		}
	}

	// 나머지 타입 변환 로직 추가

	return fmt.Errorf("타입 변환 불가능: %v → %v", mapValueType, fieldType)
}

// 맵을 string map으로 변경 -- 더 좋은거 있어서 잘 안씀
func MapToStringMap(input map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for key, value := range input {
		strValue := fmt.Sprintf("%v", value) // 값을 문자열로 변환
		result[key] = strValue
	}
	return result
}

// 더 좋은거있어서 잘 안씀
func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	if objType.Kind() != reflect.Struct {
		panic("Input must be a struct")
	}

	result := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i).Interface()
		result[field.Name] = value
	}

	return result
}
