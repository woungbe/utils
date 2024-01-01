package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// GenerateInsertQuery 함수는 tableName으로 지정된 테이블에 data로 지정된 struct의 데이터를 INSERT하는 SQL 쿼리 문자열을 생성합니다.
func GenerateInsertQuery(tableName string, data interface{}) string {
	valueType := reflect.TypeOf(data)
	value := reflect.ValueOf(data)

	if valueType.Kind() != reflect.Struct {
		return ""
	}

	var columns []string
	var values []string

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldName := field.Name
		fieldValue := value.Field(i).Interface()
		// 값이 문자열이면 작은따옴표로 감싸줍니다.
		if strVal, ok := fieldValue.(string); ok {
			fieldValue = "'" + strVal + "'"
		}
		columns = append(columns, fieldName)
		values = append(values, fmt.Sprintf("%v", fieldValue))
	}

	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(values, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columnsStr, valuesStr)

	return query
}

func GenerateInsertOrUpdateQuery(tableName string, data interface{}, updateColumns []string) string {
	valueType := reflect.TypeOf(data)
	// value := reflect.ValueOf(data)

	if valueType.Kind() != reflect.Struct {
		return ""
	}

	var columns []string

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldName := field.Name
		columns = append(columns, fieldName)
	}

	columnsStr := strings.Join(columns, ", ")

	// 업데이트할 열과 값을 정확하게 지정합니다.
	var updateValues []string
	for _, col := range updateColumns {
		updateValues = append(updateValues, fmt.Sprintf("%s = '%s'", col, getColumnValue(data, col)))
	}
	updateValuesStr := strings.Join(updateValues, ", ")

	// 단일 SQL 문으로 INSERT 또는 UPDATE 쿼리 생성
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s", tableName, columnsStr, getValuesString(data), updateValuesStr)

	return query
}

// getColumnValue 함수는 struct의 특정 열의 값을 반환합니다.
func getColumnValue(data interface{}, columnName string) string {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Struct {
		field := value.FieldByName(columnName)
		if field.IsValid() {
			return fmt.Sprintf("%v", field.Interface())
		}
	}
	return ""
}

// getValuesString 함수는 struct의 값들을 쉼표로 구분된 문자열로 반환합니다.
func getValuesString(data interface{}) string {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Struct {
		var values []string
		for i := 0; i < value.NumField(); i++ {
			values = append(values, getColumnValue(data, value.Type().Field(i).Name))
		}
		return strings.Join(values, ", ")
	}
	return ""
}
