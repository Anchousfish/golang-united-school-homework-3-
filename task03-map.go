package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {

	my_slice := make([]string, 0, len(input))
	int_slice := make([]int, 0, len(input))
	for k := range input {
		int_slice = append(int_slice, k)
	}
	sort.Ints(int_slice)
	for i := range int_slice {
		my_slice = append(my_slice, input[int_slice[i]])
	}

	return my_slice
}
