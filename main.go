package main

import (
	"fmt"
	"strings"
	"regexp"
)

func GetCount(str string) (count int) {
	splitted_string := strings.Split(str, "")
	cnt := 0
	validChar := regexp.MustCompile(`[aeiouAEIOU]`)

	for num := range splitted_string {
		if(validChar.MatchString(splitted_string[num])){
			cnt += 1
		}
	}
	return cnt
}

func main () {
	fmt.Print(GetCount("d dfj wsxA scbce"))
}
