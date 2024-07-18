package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("%v\n", canConstruct("aa", "aab"))

}

// func canConstruct(ransomNote string, magazine string) bool {
// 	ml := strings.Split(magazine, "")

// 	for _, v := range ml {
// 		ransomNote = strings.Replace(ransomNote, v, "", 1)
// 	}
// 	if ransomNote == "" {
// 		return true
// 	}
// 	return false
// }

func canConstruct(ransomNote string, magazine string) bool {
	ml := map[string]int{}

	for _, v := range strings.Split(magazine, "") {
		ml[v] += 1
	}

	var exists bool
	var val int

	for _, v := range strings.Split(ransomNote, "") {
		val, exists = ml[v]
		if !exists {
			return false
		}
		if val == 0 {
			return false
		}

		ml[v] = ml[v] - 1
	}
	return true
}
