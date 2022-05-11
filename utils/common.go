package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func JSON(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func LikeQuery(column, str string) string {
	return fmt.Sprintf("%s like %%%s%%", column, str)
}

func InQuery(column string, s []string) string {
	inStr := ""
	for _, i := range s {
		inStr += fmt.Sprintf("%s,", i)
	}
	return fmt.Sprintf("%s in (%s)", column, strings.TrimSuffix(inStr, ","))
}

func NotInQuery(column string, s []string) string {
	inStr := ""
	for _, i := range s {
		inStr += fmt.Sprintf("%s,", i)
	}
	return fmt.Sprintf("%s not in (%s)", column, strings.TrimSuffix(inStr, ","))
}
