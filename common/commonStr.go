package common

// 1、go的字符串翻转
func Reverse (str string) string {
	var reverseStr string
	strLen := len(str)
	for i:=strLen - 1; i>=0; i-- {
		reverseStr += string(str[i])
	}
	return reverseStr
}