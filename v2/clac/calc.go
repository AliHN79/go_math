package clac

func Add(args ...int) int {
	s := 0
	for _, arg := range args {
		s += arg
	}
	return s
}

func Sub(a, b int) int {
	return a - b
}
