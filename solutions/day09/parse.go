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

type file struct {
	id    int
	size  int
	start int
	moved bool
}

func parseFileList(
	input string,
) []file {
	files := make([]file, 0, len(input))

	start := 0
	for i, ch := range input {
		size := int(ch - '0')
		var id int
		if i%2 == 0 {
			id = i / 2
		} else {
			id = -1
		}

		f := file{
			id:    id,
			size:  size,
			start: start,
			moved: false,
		}
		start += size
		files = append(files, f)
	}

	return files
}
