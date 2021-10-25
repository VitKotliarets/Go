package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Setup game.
	fmt.Print("The Number of players: ")
	var number int
	fmt.Scanf("%d", &number)
	players := make([]string, number)
	for i := range players {
		fmt.Printf("The name of %d player: ", (i + 1))
		fmt.Scanf("%s", &players[i])
	}

	// Run game.
	fmt.Println("--------------------")
	stages := [3]string{"Ready", "Steady", "Go!!!"}
	for _, stage := range stages {
		fmt.Println(stage)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("--------------------")


	var iteration = 1
	for ; len(players) > 1; {
		var gamers []string
		for _, player := range players {
			var userInput string
			fmt.Println(player)
			fmt.Scanf("%s", &userInput)
			// Add to the next round player who answered correctly.
			if fizzbuzz(userInput, iteration) {
				fmt.Println("Correct")
				gamers = append(gamers, player)
			} else {
				fmt.Println("Incorrect")
			}
			iteration++
			fmt.Println("--------------------")
		}
		players = gamers
	}

	if len(players) == 0 {
		fmt.Println("You all are losers. Loooozers!!!")
	} else  {
		fmt.Printf("Winer: %s", players[0])
	}
}

func fizzbuzz(userInput string, number int) bool {
	var correctInput string
	if number % 3 == 0 && number % 5 == 0 {
		correctInput = "FizzBuzz"
	} else if number % 3 == 0 {
		correctInput = "Fizz"
	} else if number % 5 == 0 {
		correctInput = "Buzz"
	} else {
		correctInput = strconv.Itoa(number)
	}
	return userInput == correctInput
}