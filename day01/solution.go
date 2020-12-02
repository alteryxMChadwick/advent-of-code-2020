package code

func GetPair(nums []int) (int, int) {
	if 0 == len(nums) {
		return 0, 0
	}
	for _, first := range nums {
		for _, inner := range nums {
			if 2020 == first+inner {
				return first, inner
			}
		}
	}

	return 0, 0
}

func GetProduct(left int, right int) int {
	return left * right
}
