package leveldumpener

func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsSafe(levelsArr []int) bool {
	ascending := 0
	descending := 0
	ignoreCounter := 1
	diff := 0
	for i := 1; i < len(levelsArr); i++ {
		diff = levelsArr[i-1] - levelsArr[i]
		if AbsoluteValue(diff) >= 1 && AbsoluteValue(diff) <= 3 {
			if diff < 0 {
				ascending++
			} else {
				descending++
			}
		} else {
			if ignoreCounter < 1 {
				return false
			}
			ignoreCounter--
			i++
			if i < len(levelsArr) {
				diff = levelsArr[i-2] - levelsArr[i]
				if !(AbsoluteValue(diff) >= 1 && AbsoluteValue(diff) <= 3) {
					if i == 2 {
						diff = levelsArr[i-1] - levelsArr[i]
						if AbsoluteValue(diff) >= 1 && AbsoluteValue(diff) <= 3 {
							continue
						}
					}
					return false
				}
			} else {
				continue
			}
		}
	}
	if ascending == 0 || descending == 0 {
		return true
	}
	if ignoreCounter == 1 {
		if ascending == 1 || descending == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}
