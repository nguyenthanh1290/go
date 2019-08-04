package main

import (
	"fmt"
	"strings"
)

const (
	space = " "
	dash  = "-"
)

func format(phone string) string {
	phone = removeNonDigits(phone)

	if len(phone) <= 3 {
		return phone
	}

	var result strings.Builder
	i := 0
	for i < len(phone)-4 {
		result.WriteString(phone[i : i+3])
		result.WriteString(dash)

		i += 3
	}

	remainingDigits := len(phone) - i
	switch remainingDigits {
	case 2, 3:
		result.WriteString(phone[i:])
		break
	case 4:
		result.WriteString(phone[i : i+2])
		result.WriteString(dash)
		result.WriteString(phone[i+2:])
		break
	}

	return result.String()
}

func removeNonDigits(s string) string {
	s = strings.ReplaceAll(s, space, "")
	s = strings.ReplaceAll(s, dash, "")
	return s
}

func main() {
	fmt.Println(format(""))                    //
	fmt.Println(format("-"))                   //
	fmt.Println(format("0-"))                  // 0
	fmt.Println(format("00"))                  // 00
	fmt.Println(format("0-0"))                 // 00
	fmt.Println(format("00-"))                 // 00
	fmt.Println(format("014-"))                // 014
	fmt.Println(format("01-4-"))               // 014
	fmt.Println(format("01-4-1"))              // 01-41
	fmt.Println(format("00-44  48 5555 8361")) // 004-448-555-583-61
	fmt.Println(format("0 - 22 1985-324"))     // 022-198-53-24
	fmt.Println(format("555372654"))           // 555-372-654
}
