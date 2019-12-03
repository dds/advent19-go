package aoc19

// Run ...
func Run(data []int) []int {
	for i := 0; ; {
		switch data[i] {
		case 99:
			// halt
			return data
		case 1:
			// add
			data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
			i += 4
		case 2:
			// multiply
			data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
			i += 4
		}
	}
}
