package isbn

import (
	"regexp"
	"strconv"
)

func ValidISBN(isbn string) bool {
	re := regexp.MustCompile(`[^0-9]`)
	str := re.ReplaceAll([]byte(isbn), []byte(""))
	digits := make([]int, len(str))

	for k, v := range str {
		i, _ := strconv.Atoi(string(v))
		if k%2 != 0 {
			digits[k] = i * 3
		} else {
			digits[k] = i
		}
	}

	var count int
	for _, v := range digits {
		count += v
	}

	return count%10 == 0
}
