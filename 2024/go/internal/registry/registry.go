package registry

type SolveFunc func(inputFile string) (int64, error)

var Registry = make(map[string]SolveFunc)
