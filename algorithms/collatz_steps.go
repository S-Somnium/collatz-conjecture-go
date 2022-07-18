package algorithms

func CollatzStepsCalc(initial int) (steps int) {
	previus := initial
	steps = 0
	for previus != 1 {
		previus = calculateNextValue(previus)
		steps++
	}
	return steps
}

func calculateNextValue(previous int) int {
	isOdd := previous % 2
	if isOdd == 1 {
		return calculateOddValue(previous)
	} else {
		return calculateEvenValue(previous)
	}
}

func calculateOddValue(previous int) int {
	return previous*3 + 1
}

func calculateEvenValue(previous int) int {
	return previous / 2
}
