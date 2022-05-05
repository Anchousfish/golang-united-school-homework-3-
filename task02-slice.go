package homework

func reverse(input []int64) (result []int64) {

	my_slice := make([]int64, 0, len(input))

	for i := len(input) - 1; i >= 0; i-- {
		my_slice = append(my_slice, input[i])
	}

	return my_slice
}
