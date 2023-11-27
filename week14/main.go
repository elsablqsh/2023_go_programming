package main

import "fmt"

func main() {
	//balls := make(map[string]int)
	var balls map[string]int
	fmt.Printf("%#v\n", balls)
	balls["elsa"] = 27
	fmt.Println(balls["elsa"])
	fmt.Println(balls["su"])
}
