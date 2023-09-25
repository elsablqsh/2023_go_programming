package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "I love H#eningkai"
	replaceWords := strings.NewReplacer("#", "u")
	fixedWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWords)

}
