package conv

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func ToStr(n int) string {
	return strconv.Itoa(n)
}
