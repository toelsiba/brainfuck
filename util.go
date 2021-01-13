package brainfuck

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}
