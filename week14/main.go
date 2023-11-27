package main

import "fmt"

func main() {
	//var games map[int]string
	// games = make(map[int]string)

	games := map[int]string{
		427: "elsa",
		251: "su",
		108: "kamal",
		106: "victor",
		408: "yujin"}
	// append

	for _, v := range games {
		fmt.Println(v)
	}
	games[312] = "tae" // update
	delete(games, 408) // delete

	for k, v := range games {
		fmt.Println(k, v)
	}
}
