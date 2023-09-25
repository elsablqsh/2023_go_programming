package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1 //1~100

	fmt.Println("Guess Game (1~100) : ")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("Your chance :", 10-i)
		fmt.Print("Guess the number :")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber == answer {
			fmt.Printn("CORRECT! CONGRATULATIONS!")
		} else if inputNumber < answer {
			fmt.Println("Your guessed number is lower than the answer.") // answer should be higher
		} else if inputNumber > answer {
			fmt.Println("Your guessed number is higher than the answer") // answer is lower
		}

	}
}
