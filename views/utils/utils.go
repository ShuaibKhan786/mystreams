package utils

import "strings"

func SanitizeStr(str string) string {
	strArr := strings.Split(str, "\n")
	firstStr := strings.Join(strArr, " ")

	strArr = strings.Split(firstStr, "\t")

	return strings.Join(strArr, " ")
}
