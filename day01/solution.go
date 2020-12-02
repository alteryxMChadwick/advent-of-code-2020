package day01

//finds the two values in the given set that sum up to 2020
func GetPair(nums []int) (int, int) {
	if len(nums) < 2 {
		return 0, 0
	}

	uniqueNums := map[int]bool{}
	for _, num := range nums {
		uniqueNums[num] = false
	}

	for first := range uniqueNums {
		for second := range uniqueNums {
			if 2020 == first+second {
				if first < second {
					return second, first
				} else {
					return first, second
				}
			}
		}
		//this number isn't part of any combination that makes a solution, remove it from the set before future iterations
		delete(uniqueNums, first)
	}

	return 0, 0
}

//produces the product of all the inputs
func GetProduct(left int, right int) int {
	return left * right
}

//finds the three values in the given set that sum up to 2020
func GetTriplet(nums []int) (int, int, int) {
	if len(nums) < 3 {
		return 0, 0, 0
	}

	uniqueNums := map[int]int{}
	for _, num := range nums {
		uniqueNums[num]++
	}

	for first := range uniqueNums {
		for second, count := range uniqueNums {
			if first == second && count < 2 {
				continue
			}
			for third, count := range uniqueNums {
				if first == second && second == third && count < 3 {
					continue
				}
				if second == third || first == third && count < 2 {
					continue
				}
				if 2020 == first+second+third {
					if first <= second && second <= third {
						return third, second, first
					} else if second <= first && first <= third {
						return third, first, second
					} else if third <= second && second <= first {
						return first, second, third
					} else if second <= third && third <= first {
						return first, third, second
					} else if third <= first && first <= second {
						return second, first, third
					} else if first <= third && third <= second {
						return second, third, first
					} else {
						return first, second, third
					}
				}
			}
		}
		//this number isn't part of any combination that makes a solution, remove it from the set before future iterations
		delete(uniqueNums, first)
	}

	return 0, 0, 0
}

//produces the product of all the inputs
func GetTripletProduct(left int, middle int, right int) int {
	return left * middle * right
}
