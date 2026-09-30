package main

import (
	"fmt"
	"slices"
)

func main() {
	var passed bool = isAnagram("hannah", "hannah")
	fmt.Println(passed)

	passedWithMap := isAnagramWithMap("racecar", "racecar")
	fmt.Println(passedWithMap)
}

func isAnagram(s string, t string) bool {
	count := make([]rune, 26)

	for _, char := range s {
		count[rune(char)-'a']++
	}

	for _, char := range t {
		count[rune(char)-'a']--
	}

	return slices.Equal(count, make([]rune, 26))
}


// If we wanted to do via using maps and any unicode character
func isAnagramWithMap(s string, t string) bool {
	if len(s) != len(t) {
			return false
	}
	
	count := make(map[rune]int)
	
	for _, char := range s {
			count[char]++
	}
	
	for _, char := range t {
			count[char]--
	}
	
	for _, v := range count {
			if v != 0 {
					return false
			}
	}
	
	return true
}
