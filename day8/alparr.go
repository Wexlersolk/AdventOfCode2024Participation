package main

func alparr() []string {
	var result []string

	for ch := 'a'; ch <= 'z'; ch++ {
		result = append(result, string(ch))
	}

	for ch := 'A'; ch <= 'Z'; ch++ {
		result = append(result, string(ch))
	}

	for ch := '0'; ch <= '9'; ch++ {
		result = append(result, string(ch))
	}
	return result
}

