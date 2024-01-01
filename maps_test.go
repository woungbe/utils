package utils

import (
	"fmt"
	"testing"
)

func TestMapToStruct(t *testing.T) {

	type MyStruct struct {
		Field1 string  `json:"Field1"`
		Field2 int64   `json:"Field2"`
		Field3 float64 `json:"Field3"`
	}

	mapData := map[string]interface{}{
		"Field1": "Hello",
		"Field2": "42", // 문자열로 오는 데이터
		"Field3": 3.14159,
	}

	var myStruct MyStruct

	// MapToStruct 함수를 사용하여 맵을 구조체로 변환
	if err := MapToStruct("json", mapData, &myStruct); err != nil {
		fmt.Println("오류:", err)
		return
	}

	fmt.Println("Field1:", myStruct.Field1)
	fmt.Println("Field2:", myStruct.Field2)
	fmt.Println("Field3:", myStruct.Field3)

}
