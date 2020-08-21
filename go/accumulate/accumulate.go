package accumulate

//Accumulate - given a collection perform an operation on each element of the collection
func Accumulate(collection []string, operation func(string) string) []string {
	for i, s := range collection {
		collection[i] = operation(s)
	}
	return collection
}
