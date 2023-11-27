package main

import "fmt"

func status(name string) {
	balls := map[string]int{"elsa": 27, "su": 0}
	//ball := balls[name]
	ball, exists := balls[name]
	fmt.Println(ball, exists)
}

func main() {
	status("elsa")
	status("su")
	status("kamal")
}
