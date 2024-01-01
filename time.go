package utils

import (
	"fmt"
	"time"
)

// 다음 정각시간까지 남은 시간 알려주기
func NextHour() time.Duration {
	// 현재 시간을 얻어옵니다.
	now := time.Now()

	// 다음 정각을 계산합니다.
	nextHour := now.Add(time.Hour)
	nextHour = time.Date(nextHour.Year(), nextHour.Month(), nextHour.Day(), nextHour.Hour(), 0, 0, 0, nextHour.Location())

	return nextHour.Sub(now)
}

// string -> time 으로 변경  layout 조심.
//
//	layout := "2006-01-02 15:04:05"
func StringToTime(input string, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func TimeToString(input time.Time, layout string) string {
	return input.Format(layout)
}

func UnixToDate(unixTime int64) (string, error) {
	times, err := unixTimestampToTime(unixTime)
	if err != nil {
		return "", err
	}

	send := TimeTostring(times)
	return send, nil
}

func unixTimestampToTime(unixTime int64) (time.Time, error) {
	// UNIX 시간 스탬프의 자릿수에 따라 처리
	switch {
	case unixTime >= 1e18: // 19 자리
		return time.Unix(unixTime/1e9, unixTime%1e9), nil
	case unixTime >= 1e15: // 16 자리
		return time.Unix(unixTime/1e6, unixTime%1e6*1e3), nil
	case unixTime >= 1e12: // 16 자리
		return time.Unix(unixTime/1e3, unixTime%1e3*1e3), nil
	case unixTime >= 1e9: // 10 자리
		return time.Unix(unixTime, 0), nil
	default:
		return time.Time{}, fmt.Errorf("Invalid UNIX timestamp")
	}
}

func TimeTostring(t time.Time) string {
	formatted := t.Format("2006-01-02")
	return formatted
}

// 블랙처리하기 -
func IsBlockTime(start, end int, location string) bool {
	// loc, _ := time.LoadLocation("Asia/Seoul")
	loc, _ := time.LoadLocation(location)
	koreaTime := time.Now().In(loc)

	// 현재 시간이 01:00부터 06:00 사이인지 확인합니다.
	if koreaTime.Hour() > end && koreaTime.Hour() < start {
		// 특정 시간대에는 데이터를 받지 않습니다.
		// fmt.Printf("현재 한국 시간: %s, 데이터를 받지 않습니다.\n", koreaTime.Format("15:04:05"))
		return false
	} else {
		// 특정 시간대가 아니라면 데이터를 받아서 처리합니다.
		// fmt.Printf("현재 한국 시간: %s, 데이터를 받아서 처리합니다.\n", koreaTime.Format("15:04:05"))
		return true
		// 여기에서 데이터를 받아서 처리하는 코드를 추가합니다.
		// 이 부분에 원하는 로직을 구현하세요.
	}
}
