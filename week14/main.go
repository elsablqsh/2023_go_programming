package main

import "fmt"

func status(name string) {
	balls := map[string]int{"elsa": 27, "su": 0}
	//ball := balls[name]
	ball, exists := balls[name]
	if !exists {
		fmt.Println(name, "is not game participant")
	} else if ball < 1 {
		fmt.Println(name, "is", balls[name], "is eliminated")
	} else {
		fmt.Println(name, "wins the game!!")
	}
	//fmt.Println(ball, exists)
}

func main() {
	status("elsa")
	status("su")
	status("kamal")
}
