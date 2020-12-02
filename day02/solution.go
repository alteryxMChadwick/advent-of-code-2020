package day02

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(input string) (int, int, string, string, error) {
	parts := strings.Split(input, ":")

	if 2 != len(parts) {
		return -1, -1, "", "", errors.New("invalid argument")
	}

	reqs := parts[0]
	pwd := strings.TrimSpace(parts[1])

	parts = strings.Split(reqs, " ")

	if 2 != len(parts) {
		return -1, -1, "", "", errors.New("invalid argument")
	}

	countRange := parts[0]
	req := parts[1]

	parts = strings.Split(countRange, "-")

	if 2 != len(parts) {
		return -1, -1, "", "", errors.New("invalid argument")
	}

	minStr := parts[0]
	maxStr := parts[1]

	min, err := strconv.ParseInt(minStr, 10, 32)
	if err != nil {
		return -1, -1, "", "", err
	}
	max, err := strconv.ParseInt(maxStr, 10, 32)
	if err != nil {
		return -1, -1, "", "", err
	}

	return int(min), int(max), req, pwd, nil
}

func ValidatePasswords(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		min, max, req, pwd, err := Parse(input)
		if err != nil {
			return -1, err
		}
		count := strings.Count(pwd, req)
		if min <= count && count <= max {
			sum++
		}
	}

	return sum, nil
}