package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	whiteSphereNumbers = 69
	redSphereNumbers   = 26
	whiteTicketNumbers = 5
	redTicketNumbers   = 1
)

func draw(numberOfBalls int) []int {
	sphere := rand.Perm(numberOfBalls)
	for i := 0; i < numberOfBalls; i++ {
		sphere[i]++
	}
	return sphere
}

func main() {
	fmt.Println("Welcome to Lottery Generator for Powerball game")
	fmt.Println(strings.Repeat("-", 42))

	fmt.Println("Press [Enter] to generate numbers for a Powerball ticket")
	fmt.Scanln()

	rand.Seed(time.Now().UTC().UnixNano())

	var whiteSphere []int
	whiteSphere = draw(whiteSphereNumbers)
	redSphere := draw(redSphereNumbers)

	whiteChoice := whiteSphere[:whiteTicketNumbers]
	redChoice := redSphere[:redTicketNumbers]
	fmt.Println("Numbers selected for your ticket:")
	fmt.Printf("- White: %v\r\n", whiteChoice)
	fmt.Printf("- Red: %v\r\n", redChoice)

	fmt.Println()
	fmt.Println("Smells like a jackpot!")
	fmt.Println("Wish you luck!")
}
