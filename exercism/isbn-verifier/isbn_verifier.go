package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN validates the input ISBN number.
func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}

	arr := strings.Split(isbn, "")
	sum := 0
	
	if arr[len(arr)-1] == "X" {
		sum += 10
		arr = arr[0 : len(arr)-1]
	}

	for i := 0; i < len(arr); i++ {
		v, err := strconv.Atoi(arr[i])
		if err != nil {
			return false
		}
		sum += v * (10 - i)
	}

	if sum%11 == 0 {
		return true
	}

	return false
}
