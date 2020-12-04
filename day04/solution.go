package day04

import (
	"errors"
	"strconv"
	"strings"
)

func checkIntInRange(num, min, max int64) bool {
	return min <= num && num <= max
}

func checkStringInIntRange(str string, min, max int64) (bool, error) {
	num, err := strconv.ParseInt(str, 10, 32)
	if nil != err {
		return false, err
	}
	return checkIntInRange(num, min, max), nil
}

func CheckBirthYear(yr string) (bool, error) {
	return checkStringInIntRange(yr, 1920, 2002)
}

func CheckIssueYear(yr string) (bool, error) {
	return checkStringInIntRange(yr, 2010, 2020)
}

func CheckExpirationYear(yr string) (bool, error) {
	return checkStringInIntRange(yr, 2020, 2030)
}

func CheckHeight(height string) (bool, error) {
	l := len(height)
	if l < 4 || 5 < l {
		return false, errors.New("incorrect string length")
	}
	heightType := height[l-2 : l]
	if "in" == heightType {
		return checkStringInIntRange(height[:l-2], 59, 76)
	} else if "cm" == heightType {
		return checkStringInIntRange(height[:l-2], 150, 193)
	}
	return false, errors.New("missing height type")
}

func CheckHairColor(hairColor string) (bool, error) {
	if 7 != len(hairColor) {
		return false, errors.New("incorrect string length")
	}
	if '#' != hairColor[0] {
		return false, errors.New("missing '#'")
	}

	for _, char := range hairColor[1:] {
		if !checkIntInRange(int64(char), '0', '9') && !checkIntInRange(int64(char), 'a', 'f') {
			return false, nil
		}
	}
	return true, nil
}

var (
	validEyeColors = map[string]bool {
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
)

func CheckEyeColor(eyeColor string) (bool, error) {
	if 3 != len(eyeColor) {
		return false, errors.New("incorrect string length")
	}
	_, ok := validEyeColors[eyeColor]
	return ok, nil
}

func CheckPassportId(passportId string) (bool, error) {
	if 9 != len(passportId) {
		return false, errors.New("incorrect string length")
	}
	for _, char := range passportId {
		if !checkIntInRange(int64(char), '0', '9') {
			return false, nil
		}
	}
	return true, nil
}

func CheckCountryId(countryId string) (bool, error) {
	return true, nil
}

func CheckKeysPart2(passport string, requiredKeys map[string]func(string) (bool, error)) bool {
	components := strings.Split(passport, " ")

	foundKeys := map[string]string{}
	for _, pair := range components {
		keyVal := strings.Split(pair, ":")
		if 2 != len(keyVal) {
			return false
		}

		foundKeys[keyVal[0]] = keyVal[1]
	}

	for key, check := range requiredKeys {
		val := foundKeys[key]
		if result, err := check(val); !result || nil != err {
			return false
		}
	}

	return true
}

func CountValidIdsPart2(passportBlob string, keys map[string]func(string) (bool, error)) int {
	passports := SplitBlobToStrings(passportBlob)
	count := 0
	for _, passport := range passports {
		if CheckKeysPart2(passport, keys) {
			count++
		}
	}
	return count
}

func SplitBlobToStrings(passports string) []string {
	passports = strings.ReplaceAll(passports, "\n\n", "\r")
	passports = strings.ReplaceAll(passports, "\n", " ")
	return strings.Split(passports, "\r")
}

func CheckKeysPart1(passport string, requiredKeys map[string]bool) bool {
	components := strings.Split(passport, " ")

	foundKeys := map[string]string{}
	for _, pair := range components {
		keyVal := strings.Split(pair, ":")
		if 2 != len(keyVal) {
			return false
		}

		foundKeys[keyVal[0]] = keyVal[1]
	}

	for key, required := range requiredKeys {
		_, ok := foundKeys[key]
		if !ok && required {
			return false
		}
	}

	return true
}

func CountValidIdsPart1(passportBlob string, keys map[string]bool) int {
	passports := SplitBlobToStrings(passportBlob)
	count := 0
	for _, passport := range passports {
		if CheckKeysPart1(passport, keys) {
			count++
		}
	}
	return count
}
