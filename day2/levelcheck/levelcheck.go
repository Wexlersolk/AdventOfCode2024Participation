package levelcheck

func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsSafe(levelsArr []int) bool {
	ascending := false
	for i := 1; i < len(levelsArr); i++ {
		diff := levelsArr[i-1] - levelsArr[i]
		if diff < 0 && i == 1 {
			ascending = true
		}
		if AbsoluteValue(diff) >= 1 && AbsoluteValue(diff) <= 3 {
			if diff < 0 {
				if ascending != true {
					return false
				}
				ascending = true
			} else {
				if ascending != false {
					return false
				}
				ascending = false
			}

		} else {
			return false
		}

	}
	return true
}
