package utils

import (
	"fmt"
	"testing"
)

func TestStringToTime(t *testing.T) {

	// input 형태에 맞춰서 layout 를 지정해주면 time으로 인식할 수 있음.
	input := "11-12-13 2024.01.01"
	layout := "15-04-05 2006.01.02"

	ts, er := StringToTime(input, layout)

	if er != nil {
		fmt.Println("er : ", er)
	}
	fmt.Println(ts)

}

func TestIsBlockTIme(t *testing.T) {
	b := IsBlockTime(01, 06, "Asia/Seoul")
	if b {
		fmt.Println("IsBlockTime : ", b)
	} else {
		fmt.Println("IsBlockTime : ", b)
	}

}
