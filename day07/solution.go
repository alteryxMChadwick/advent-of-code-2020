package day07

import (
	"errors"
	"strings"
)

func BagFits(rules map[string][]string, containedBags []string, bag string) bool {
	for _, containedBag := range containedBags {
		if bag == containedBag {
			return true
		}
		if BagFits(rules, rules[containedBag], bag) {
			return true
		}
	}

	return false
}

func BuildRules(rulesBlob string) (map[string][]string, error) {
	rules := strings.Split(rulesBlob, "\n")
	returnVal := make(map[string][]string)
	for _, rule := range rules {
		rulePair := strings.Split(rule, " contain ")
		if 2 != len(rulePair) {
			return nil, errors.New("invalid rule")
		}
		containedBags := strings.Split(rulePair[1], ", ")
		for idx, bag := range containedBags {
			bag = strings.TrimLeft(bag, "0123456789 ")
			elements := strings.Split(bag, " ")
			if 3 != len(elements) && 2 != len(elements) {
				return nil, errors.New("invalid rule")
			}
			containedBags[idx] = strings.Join(elements[0:2], " ")
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
