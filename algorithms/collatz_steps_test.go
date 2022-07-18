package algorithms

import (
	"testing"
)

type data struct {
	value    int
	expected int
}

var sample = []data{
	{value: 1, expected: 0},
	{value: 2, expected: 1},
	{value: 11, expected: 14},
	{value: 13, expected: 9},
	{value: 131, expected: 28},
	{value: 132, expected: 28},
	{value: 54, expected: 112},
	{value: 543, expected: 136},
}

func TestCollatzSetps(t *testing.T) {
	for _, data := range sample {
		steps := CollatzStepsCalc(data.value)
		if steps != data.expected {
			t.Errorf("Collatz Steps FAILED. Expected %d, got %d. Parameter: %d", data.expected, steps, data.value)
			return
		}
	}
	t.Logf("Collatz Steps PASSED.")
}
