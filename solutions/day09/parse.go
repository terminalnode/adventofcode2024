package main

func parse(
	input string,
) []int {
	disk := make([]int, 0, len(input)*9)
	for i, ch := range input {
		size := int(ch - '0')
		var id int
		if i%2 == 0 {
			id = i / 2
		} else {
			id = -1
		}

		for chunkI := 0; chunkI < size; chunkI++ {
			disk = append(disk, id)
		}
	}

	return disk
}
