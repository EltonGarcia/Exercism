package strain

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(predicate func(int) bool) (result Ints) {
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return
}

func (list Ints) Discard(predicate func(int) bool) Ints {
	return list.Keep(func(n int) bool { return !predicate(n) })
}

func (list Lists) Keep(predicate func([]int) bool) (result Lists) {
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return
}

func (list Strings) Keep(predicate func(string) bool) (result Strings) {
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return
}
