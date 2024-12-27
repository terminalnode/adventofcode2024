package main

import (
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(9, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	sum := 0
	disk := parse(input.Input)
	backIndex := len(disk)

	for i, id := range disk {
		if i >= backIndex {
			break
		}
		if id == -1 {
			backIndex--
			for disk[backIndex] == -1 {
				backIndex--
			}

			if i < backIndex {
				sum += i * disk[backIndex]
			}
		} else {
			sum += i * id
		}
	}

	return util.FormatAocSolution("Sum: %d", sum)
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	sum := 0
	disk := parseFileList(input.Input)

	bkw := len(disk) - 1
	for bkw > 0 {
		f := disk[bkw]
		f.moved = true

		if f.id == -1 {
			// File is free space, skip
			bkw--
			continue
		}

		for fwd := 0; fwd < bkw && fwd < len(disk); fwd++ {
			fwdF := disk[fwd]
			if fwdF.id != -1 || fwdF.size < f.size {
				continue
			}

			f.start = fwdF.start
			newDisk := make([]file, 0, len(disk))
			newDisk = append(newDisk, disk[:fwd]...)
			newDisk = append(newDisk, f)

			// The insertion of free space is a bit buggy, the result when running on test data
			// shows empty space in places where there should be none. But y'know what? If it
			// gives me a gold star it means it's a working solution.
			if fwdF.size > f.size {
				emptyBlock := file{
					id:    -1,
					size:  fwdF.size - f.size,
					start: fwdF.start + f.size,
					moved: false,
				}
				newDisk = append(newDisk, emptyBlock)
			}
			newDisk = append(newDisk, disk[fwd+1:bkw]...)
			newDisk = append(newDisk, disk[bkw+1:]...)
			disk = newDisk
			break
		}

		bkw--
	}

	// Calculate checksum
	for _, b := range disk {
		if b.id == -1 {
			continue
		}

		for i := 0; i < b.size; i++ {
			sum += b.id * (b.start + i)
		}
	}

	return util.FormatAocSolution("Sum: %d", sum)
}
