package strings

import "reflect"

func createLettersCountMap(s string) map[int32]int {
	var lettersCountMap = make(map[int32]int)
	for _, char := range s {
		lettersCountMap[char] = lettersCountMap[char] + 1
	}
	return lettersCountMap
}

func isAnagrams(s1, s2 string) bool {
	return reflect.DeepEqual(createLettersCountMap(s1), createLettersCountMap(s2))
}
