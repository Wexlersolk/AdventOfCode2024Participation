package main

func calculate2(disk []string) []string {
	for i := len(disk) - 1; i >= 0; {
		if disk[i] == "." {
			i--
			continue
		}
		number := disk[i]
		numbercount := 0

		for m := i; m >= 0 && disk[m] == number; m-- {
			numbercount++
		}

		disk = dotnumberswap(disk, i-numbercount+1, i+1)

		i -= numbercount
	}

	return disk
}

func dotnumberswap(disk []string, k, n int) []string {
	row := disk[k:n]
	glen := len(row)

	for j := 0; j < len(disk); j++ {
		if disk[j] == row[0] {
			break
		}
		if disk[j] == "." {
			for m := j; m < len(disk); m++ {
				if disk[m] != "." {
					if m-j >= glen {
						copy(disk[j:j+glen], row)
						for v := range row {
							row[v] = "."
						}
					}
					break
				}
			}
		}
	}

	return disk
}
