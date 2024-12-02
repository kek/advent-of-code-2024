package historians

import (
	"sort"
	"strconv"
	"strings"
)

func ParseInput(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		l, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			continue
		}
		left = append(left, int(l))
		r, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			continue
		}
		right = append(right, int(r))
	}
	return left, right
}

func SolveDay1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i, l := range left {
		r := right[i]
		if l < r {
			total += r - l
		} else {
			total += l - r
		}
	}

	return total
}
