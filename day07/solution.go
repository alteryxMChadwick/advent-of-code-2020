package day07

import (
	"errors"
	"strconv"
	"strings"
)

func BagFits(rules map[string][]BagCount, containedBags []BagCount, bag string) bool {
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

func BuildRules(rulesBlob string) (map[string][]BagCount, error) {
	rules := strings.Split(rulesBlob, "\n")
	returnVal := make(map[string][]BagCount)
	for _, rule := range rules {
		rulePair := strings.Split(rule, " contain ")
		if 2 != len(rulePair) {
			return nil, errors.New("invalid rule")
		}
		containedBags := make([]BagCount, 0, 4)
		for _, bag := range strings.Split(rulePair[1], ", ") {
			elements := strings.Split(bag, " ")
			if 4 != len(elements) && 3 != len(elements) {
				return nil, errors.New("invalid rule")
			}
			if "no" != elements[0] {
				count, err := strconv.ParseInt(elements[0], 10, 32)
				if nil != err {
					return nil, err
				}
				containedBags = append(containedBags, BagCount{int(count),
					strings.Join(elements[1:3], " ")})
			}
		}

		bagNameParts := strings.Split(rulePair[0], " ")
		if 3 != len(bagNameParts) {
			return nil, errors.New("invalid rule")
		}

		bagName := strings.Join(bagNameParts[0:2], " ")

		returnVal[bagName] = containedBags
	}
	return returnVal, nil
}

func BagContains(rules map[string][]BagCount, containingBag string, contained bool) int {
	count := 0
	for _, containedBag := range rules[containingBag] {
		count += containedBag.Count * BagContains(rules, containedBag.Bag, true)
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
