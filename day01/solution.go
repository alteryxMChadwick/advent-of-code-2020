package code

func GetPair(nums []int) (int, int) {
	if 0 == len(nums) {
		return 0, 0
	}

	uniqueNums := map[int]bool{}
	for _, num := range nums {
		uniqueNums[num] = false
	}

	for first, _ := range uniqueNums {
		for inner, _ := range uniqueNums {
			if 2020 == first+inner {
				if first < inner {
					return inner, first
				} else {
					return first, inner
				}
			}
		}
		delete(uniqueNums, first)
	}

	return 0, 0
}

func GetProduct(left int, right int) int {
	return left * right
}
