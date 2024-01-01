package utils

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func ConvertAmount(amount float64, pow int) *big.Int {
	// 0.01을 18자리 소수점 이하로 이동
	ethValue := new(big.Float).SetFloat64(amount)
	s := fmt.Sprintf("1e%d", pow)
	f, _ := strconv.ParseFloat(s, 10)
	ethValue.Mul(ethValue, big.NewFloat(f)) // 1e18은 10^18을 나타냅니다.

	// 소수점 이하를 정수로 변환
	ethInt := new(big.Int)
	ethValue.Int(ethInt)

	return ethInt
}

// 소수점 자릿수 파악하기
func DecimalCount(price string) (bool, int) {
	value, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("부동소수점 변환 오류:", err)
		return false, -1
	}

	var decimalPlaces int
	decimalPlaces = -1

	// 부동소수점 값을 문자열로 변환하여 소수점 이하 자릿수 계산
	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(strValue, ".")
	if len(parts) == 2 {
		decimalPlaces = len(parts[1])
		fmt.Println("소수점 이하 자릿수:", decimalPlaces)
	} else {
		if len(parts[0]) == 1 {
			i, _ := strconv.Atoi(parts[0])
			if i < 10 {
				return true, 0
			}
		}
		fmt.Println("부동소수점 형식이 아닙니다.")
		return false, decimalPlaces
	}

	return true, decimalPlaces
}

// 맨끝에 있는 0 지우기
func LastWordDel(str, key string) string {
	return strings.TrimRight(str, key)
}

// 자릿수 반올림해서 자르기
func RoundFloat(num float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	rounded := math.Round(num*shift) / shift
	return rounded
}
