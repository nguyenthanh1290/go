package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid determines whether or not a number is valid per the Luhn formula.
func Valid(id string) bool {
	id = strings.Replace(id, " ", "", -1)
	re := regexp.MustCompile(`^\d{2,}$`)
	if !re.MatchString(id) {
		return false
	}

	arr := strings.Split(id, "")
	sum := 0
	alternate := false

	for i := len(arr) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(arr[i])
		// if err != nil {
		// 	return false
		// }
		if alternate {
			n <<= 1
			if n > 9 {
				n -= 9
			}
		}
		alternate = !alternate
		sum += n
	}

	return (sum%10 == 0)
}
