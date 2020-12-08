package day07

import (
	"errors"
	"strconv"
	"strings"
)

func BagFits(
	rules map[string][]BagCount,
	containedBags []BagCount,
	bag string,
) bool {
	for _, containedBag := range containedBags {
		if bag == containedBag.Bag {
			return true
		}
		if BagFits(rules, rules[containedBag.Bag], bag) {
			return true
		}
	}

	return false
}

func CountBagsThatFit(rulesBlob string, bag string) (int, error) {
	rules, err := BuildRules(rulesBlob)
	if nil != err {
		return -1, err
	}
	count := 0
	for _, bags := range rules {
		if BagFits(rules, bags, bag) {
			count++
		}
	}

	return count, nil
}

type BagCount struct {
	Count int
	Bag   string
}

func GetContainedBags(bagList string) ([]BagCount, error) {
	containedBags := make([]BagCount, 0, 8)
	for _, bag := range strings.Split(bagList, ", ") {
		bag, err := NewBag(bag)
		if nil != err {
			return nil, err
		}
		if nil != bag {
			containedBags = append(containedBags, *bag)
		}
	}
	return containedBags, nil
}

func NewBag(bag string) (*BagCount, error) {
	firstIdx := strings.Index(bag, " ")
	lastIdx := strings.LastIndex(bag, " ")
	if firstIdx < 0 || len(bag) <= firstIdx || lastIdx < 0 || len(bag) <= lastIdx {
		return nil, errors.New("invalid rule")
	}
	elements := strings.Split(bag, " ")
	if 4 != len(elements) && 3 != len(elements) {
		return nil, errors.New("invalid rule")
	}
	countStr := bag[0:firstIdx]
	if "no" != countStr {
		count, err := strconv.ParseInt(countStr, 10, 32)
		if nil != err {
			return nil, err
		}
		return &BagCount{int(count), bag[firstIdx+1:lastIdx]}, nil
	}
	return nil, nil
}

func NewBagName(bagName string) (string, error) {
	idx := strings.LastIndex(bagName, " ")
	if idx < 0 || len(bagName) <= idx {
		return "", errors.New("invalid rule")
	}
	return bagName[0:idx], nil
}

func BuildRules(rulesBlob string) (map[string][]BagCount, error) {
	rules := strings.Split(rulesBlob, "\n")
	returnVal := make(map[string][]BagCount)
	for _, rule := range rules {
		rulePair := strings.Split(rule, " contain ")
		if 2 != len(rulePair) {
			return nil, errors.New("invalid rule")
		}
		containedBags, err := GetContainedBags(rulePair[1])
		if nil != err {
			return nil, err
		}

		bagName, err := NewBagName(rulePair[0])
		if nil != err {
			return nil, err
		}

		returnVal[bagName] = containedBags
	}
	return returnVal, nil
}

func BagContains(
	rules map[string][]BagCount,
	containingBag string,
	contained bool,
) int {
	count := 0
	for _, containedBag := range rules[containingBag] {
		count += containedBag.Count * BagContains(
			rules,
			containedBag.Bag,
			true)
	}

	if contained {
		return count + 1
	} else {
		return count
	}
}

func CountBagsContained(rulesBlob string, bag string) (int, error) {
	rules, err := BuildRules(rulesBlob)
	if nil != err {
		return -1, err
	}

	return BagContains(rules, bag, false), nil
}
