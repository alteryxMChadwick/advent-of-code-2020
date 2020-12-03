package day03

func CountTreesPart1(input []string) int {
	return CountTreesPart2(input, 3, 1)
}

func CountTreesPart2(input []string, left int, down int) int {
	rowCount := len(input)
	if rowCount < 1 {
		return -1
	}
	rowWidth := len(input[0])
	if rowWidth < 1 {
		return -1
	}
	xPos := 0
	yPos := 0
	treesHit := 0
	treeChar := byte('#')

	for {
		if treeChar == input[yPos][xPos] {
			treesHit++
		}
		xPos += left
		yPos += down
		if rowCount <= yPos {
			break
		}
		if rowWidth <= xPos {
			xPos = xPos % rowWidth
		}
	}

	return treesHit
}

func TotalTreesPart2(input []string) (int, int, int, int, int) {
	one := CountTreesPart2(input, 1, 1)
	two := CountTreesPart2(input, 3, 1)
	three := CountTreesPart2(input, 5, 1)
	four := CountTreesPart2(input, 7, 1)
	five := CountTreesPart2(input, 1, 2)

	return one, two, three, four, five
}
