package main

import "fmt"

func main() {
	//var games map[int]string
	// games = make(map[int]string)

	games := make(map[int]string)
	// append
	games[427] = "elsa"
	games[251] = "su"
	games[108] = "kamal"
	games[106] = "victor"
	games[408] = "yujin"

	for _, v := range games {
		fmt.Println(v)
	}
	games[312] = "tae" // update
	delete(games, 408) // delete

	for k, v := range games {
		fmt.Println(k, v)
	}
}
