package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`ri`)

	fmt.Println(regex.MatchString("rifqi"))
}
