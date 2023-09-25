package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmtPrint("Input score:")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString("\n")
	fmt.Println(inputScore)
}
