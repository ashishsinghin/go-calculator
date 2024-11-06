package evaluator

type Stack []float64

func (s *Stack) Push(value float64) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (float64, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, true
}
