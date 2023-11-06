package main

import "fmt"

func main() {
	//var primes [3]int
	//primes[0] = 2
	//primes[2] = 5

	primes := [3]int{2, 3, 5} // initialized by array literal

	primes[2] = 6

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) // zero value
	fmt.Println(test)
	// fmt.Println(test[5]) invalid argument : index 5 out of bounds [0:5]

	fmt.Printf("%#v\n", primes)
	fmt.Printf("%#v\n", test)

	i := 0
	for i < 6 { // while
		fmt.Println(test[i])
		i++
	}
}
