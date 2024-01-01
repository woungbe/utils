package utils

import "runtime"

// 해당 func의 이름을 알려주는 함수로
// : 로그형태로만 사용할 것 ..
// ** 이걸 통해서 로직을 만들지 말것 - 리로스 많이 먹음 , 정확하지 않음
func CurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(0)
	functionName := runtime.FuncForPC(pc).Name()
	return functionName
}
