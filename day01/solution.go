package day01

//finds the two values in the given set that sum up to 2020
func GetPair(nums []int) (int, int) {
	if len(nums) < 2 {
		return 0, 0
	}

	for _, first := range nums {
		for _, second := range nums {
			if 2020 == first+second {
				if first < second {
					return second, first
				} else {
					return first, second
				}
			}
		}
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

	for _, first := range nums {
		for _, second := range nums {
			for _, third := range nums {
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
	}

	return 0, 0, 0
}

//produces the product of all the inputs
func GetTripletProduct(left int, middle int, right int) int {
	return left * middle * right
}
