// Package romannumerals solves the Roman Numerals exercise on Exercism's Go track.
package romannumerals

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var arabicToRoman = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// ToRomanNumeral takes in an integer and converts to roman numbers.
// It returns an error if there is a problem with the input.
func ToRomanNumeral(arabic int) (string, error) {
	var sb strings.Builder

	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("can only convert Arabic numbers between 1 and 3000")
	}

	arabicStr := strconv.Itoa(arabic)

	for i, r := range arabicStr {
		if r != '0' {
			var arabicBaseStr = string(r)

			// pad r with 0s
			zeroCount := len(arabicStr) - i
			if zeroCount > 0 {
				arabicBaseStr = RightPad2Len(arabicBaseStr, "0", zeroCount)
			}

			arabicBaseInt, err := strconv.Atoi(arabicBaseStr)
			if err != nil {
				return "", fmt.Errorf("can not convert %s to integer", arabicBaseStr)
			}

			romanBaseStr := getRomanNumeral(arabicBaseInt)

			sb.WriteString(romanBaseStr)
		}
	}

	return sb.String(), nil
}

func getRomanNumeral(arabic int) string {

	keys := getArabicKeysReverse()

	for _, i := range keys {
		if arabic >= i {
			var sb strings.Builder

			quotient := arabic / i
			modulus := arabic % i

			sb.WriteString(strings.Repeat(arabicToRoman[i], quotient))

			if modulus > 0 {
				sb.WriteString(getRomanNumeral(modulus))
			}

			return sb.String()
		}
	}

	return ""
}

func getArabicKeysReverse() []int {
	var keys []int
	for k := range arabicToRoman {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	return keys
}

// RightPad2Len https://github.com/DaddyOh/golang-samples/blob/master/pad.go
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
