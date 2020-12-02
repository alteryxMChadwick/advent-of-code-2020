package day02

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(input string) (int, int, byte, string, error) {
	parts := strings.Split(input, ":")

	if 2 != len(parts) {
		return -1, -1, ' ', "", errors.New("invalid argument")
	}

	reqs := parts[0]
	pwd := strings.TrimSpace(parts[1])

	parts = strings.Split(reqs, " ")

	if 2 != len(parts) {
		return -1, -1, ' ', "", errors.New("invalid argument")
	}

	countRange := parts[0]
	req := parts[1]

	parts = strings.Split(countRange, "-")

	if 2 != len(parts) {
		return -1, -1, ' ', "", errors.New("invalid argument")
	}

	minStr := parts[0]
	maxStr := parts[1]

	min, err := strconv.ParseInt(minStr, 10, 32)
	if err != nil {
		return -1, -1, ' ', "", err
	}
	max, err := strconv.ParseInt(maxStr, 10, 32)
	if err != nil {
		return -1, -1, ' ', "", err
	}

	return int(min), int(max), []byte(req)[0], pwd, nil
}

func ValidatePassword(input string) (bool, error) {
	min, max, req, pwd, err := Parse(input)
	if err != nil {
		return false, err
	}
	count := strings.Count(pwd, string(req))
	if min <= count && count <= max {
		return true, nil
	}

	return false, nil
}

func ValidatePasswords(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		valid, err := ValidatePassword(input)
		if err != nil {
			return -1, err
		}
		if valid {
			sum++
		}
	}

	return sum, nil
}

func ValidatePasswordPart2(input string) (bool, error) {
	first, second, req, pwd, err := Parse(input)
	if err != nil {
		return false, err
	}
	valid := false
	if first-1 <= len(pwd) {
		if req == []byte(pwd)[first-1] {
			valid = true
		}
	}
	if second-1 <= len(pwd) {
		if req == []byte(pwd)[second-1] {
			valid = !valid
		}
	}
	return valid, nil
}

func ValidatePasswordsPart2(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		valid, err := ValidatePasswordPart2(input)
		if err != nil {
			return -1, err
		}
		if valid {
			sum++
		}
	}

	return sum, nil
}
