package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type cards struct {
	num  int
	face string
}

//Card faces
var SPADE = "Spade"
var DIAMOND = "Diamond"
var CLOVER = "Clover"
var CLUB = "Club"

//Non numerical cards
var ACE int = 1
var JACK int = 11
var QUEEN int = 12
var KING int = 13

//number of cards in a deck
var MAX_CARDS int = 52

//highest point value cpu will draw until
var MAX_CPU int = 15

//deck of cards
var cardDeck [52]cards

var gameEND int = 0
var gamePlaying int = 1

//total points
var playerTotal int = 0
var cpuTotal int = 0

//hands
var playerHand []cards
var CPUHand []cards

//Prints both players hands
func printHands(playerHandLen int, CPUHandLen int, gameOver int) {

	fmt.Print("Players hand: \n")

	//prints the player's hand
	for i := 0; i < playerHandLen; i++ {

		//card is an Ace
		if playerHand[i].num == ACE {
			fmt.Print("A ", playerHand[i].face, "\t")

			//card is a Jack
		} else if playerHand[i].num == JACK {
			fmt.Print("J ", playerHand[i].face, "\t")

			//card is a Queen
		} else if playerHand[i].num == QUEEN {
			fmt.Print("Q ", playerHand[i].face, "\t")

			//card is a King
		} else if playerHand[i].num == KING {
			fmt.Print("K ", playerHand[i].face, "\t")

			//card is a number
		} else {
			fmt.Print(playerHand[i].num, " ", playerHand[i].face, "\t")
		}
	}

	fmt.Print("\n\n")
	fmt.Print("CPU's hand: \n")

	//prints the CPU's hand
	for i := 0; i < CPUHandLen; i++ {

		//game isn't over so hide first card
		if i == 0 && gameOver == 1 {
			fmt.Print("Hidden\t")

			//card is an ace
		} else if CPUHand[i].num == ACE {
			fmt.Print("A ", CPUHand[i].face, "\t")

			//card is a Jack
		} else if CPUHand[i].num == JACK {
			fmt.Print("J ", CPUHand[i].face, "\t")

			//card is a !ueen
		} else if CPUHand[i].num == QUEEN {
			fmt.Print("Q ", CPUHand[i].face, "\t")

			//card is a King
		} else if CPUHand[i].num == KING {
			fmt.Print("K ", CPUHand[i].face, "\t")

			//card is a number
		} else {
			fmt.Print(CPUHand[i].num, " ", CPUHand[i].face, "\t")
		}
	}

	fmt.Print("\n")
	fmt.Print("\n")
}

//fills a deck with 52 cards
func createDeck() {

	var counter int = 1

	//loops through and makes deck
	for i := 0; i < MAX_CARDS; i += 4 {

		cardDeck[i].num = counter
		cardDeck[i].face = SPADE

		cardDeck[i+1].num = counter
		cardDeck[i+1].face = DIAMOND

		cardDeck[i+2].num = counter
		cardDeck[i+2].face = CLOVER

		cardDeck[i+3].num = counter
		cardDeck[i+3].face = CLUB

		counter++
	}
}

//Gives each player two starting cards
func startHands() {

	var playerHandLen int = 0
	var CPUHandLen int = 0
	var drawCard int = 0

	//Gets starting two cards for the players
	for playerHandLen != 2 || CPUHandLen != 2 {

		//generates random card
		rand.Seed(time.Now().UnixNano())
		minPlayer := 0
		maxPlayer := MAX_CARDS - 1
		drawCard = rand.Intn(maxPlayer-minPlayer+1) + minPlayer

		//card is invalid
		if cardDeck[drawCard].num == -1 {

			//card can be drawn
		} else if playerHandLen != 2 {

			playerHand = append(playerHand, cardDeck[drawCard])

			cardDeck[drawCard].num = -1

			playerHandLen++
		}

		//generates random card
		rand.Seed(time.Now().UnixNano())
		minCPU := 0
		maxCPU := MAX_CARDS - 1
		drawCard = rand.Intn(maxCPU-minCPU+1) + minCPU

		//card is invalid
		if cardDeck[drawCard].num == -1 {

			//card can be drawn
		} else if CPUHandLen != 2 {

			CPUHand = append(CPUHand, cardDeck[drawCard])

			cardDeck[drawCard].num = -1

			CPUHandLen++
		}
	}
}

//Displays game menu
func displayMenu() string {

	var userInput string

	fmt.Print("What would you like to do?: (1,2,3)\n")
	fmt.Print("1. Show hands\n")
	fmt.Print("2. Play turn\n")
	fmt.Print("3. Quit\n-> ")

	fmt.Scanln(&userInput)
	fmt.Print("\n")

	return userInput
}

//processes the player's turn
func playerTurn(playerHandLen int) int {

	var validDraw int = 0
	var drawCard int = 0

	for validDraw != 1 {

		//generates random card
		rand.Seed(time.Now().UnixNano())
		minPlayer := 0
		maxPlayer := MAX_CARDS - 1
		drawCard = rand.Intn(maxPlayer-minPlayer+1) + minPlayer

		//card is invalid
		if cardDeck[drawCard].num == -1 {

			//card can be drawn
		} else {

			playerHand = append(playerHand, cardDeck[drawCard])

			cardDeck[drawCard].num = -1

			playerHandLen++

			validDraw = 1
		}
	}

	return playerHandLen
}

//calculates the player's point value
func playerPoints(playerHandLen int) {

	var total int = 0

	for i := 0; i < playerHandLen; i++ {

		if i == 0 {
			//player has an ace and a 10, jack, queen, or king
			if ((playerHand[i].num == ACE && playerHand[i+1].num >= 10) || (playerHand[i].num >= 10 && playerHand[i+1].num == ACE)) && playerHandLen == 2 {

				total = 21
				break
			}
		}

		//user has a jack, queen, or king
		if playerHand[i].num == JACK || playerHand[i].num == QUEEN || playerHand[i].num == KING {

			total += 10

		} else {

			total += playerHand[i].num
		}
	}

	playerTotal = total
}

//processes the cpu's turn
func cpuTurn(CPUHandLen int) int {

	var validDraw int = 0
	var drawCard int = 0

	for validDraw != 1 {

		//generates random card
		rand.Seed(time.Now().UnixNano())
		minCPU := 0
		maxCPU := MAX_CARDS - 1
		drawCard = rand.Intn(maxCPU-minCPU+1) + minCPU

		//card is invalid
		if cardDeck[drawCard].num == -1 {

			//card can be drawn
		} else {

			CPUHand = append(CPUHand, cardDeck[drawCard])

			cpuTotal += cardDeck[drawCard].num

			cardDeck[drawCard].num = -1

			CPUHandLen++

			validDraw = 1
		}
	}

	return CPUHandLen
}

//calculates the cpu's point value
func cpuPoints(CPUHandLen int) {

	var total int = 0

	//counts CPU's hand
	for i := 0; i < CPUHandLen; i++ {

		if i == 0 {
			//cpu has an ace and a 10, jack, queen, or king
			if ((CPUHand[i].num == ACE && CPUHand[i+1].num >= 10) || (CPUHand[i].num >= 10 && CPUHand[i+1].num == ACE)) && CPUHandLen == 2 {

				total = 21
				break
			}
		}
		//cpu has a jack, queen, or king
		if CPUHand[i].num == JACK || CPUHand[i].num == QUEEN || CPUHand[i].num == KING {

			total += 10

		} else {

			total += CPUHand[i].num
		}
	}

	cpuTotal = total
}

//starts the game
func gameStart() {

	//creates a new deck of cards
	createDeck()

	//clears hands
	playerHand = nil
	CPUHand = nil

	var userInput string
	var gameOver int = 0
	var cpuStop int = 0
	var playerHandLen int = 2
	var CPUHandLen int = 2

	playerTotal = 0
	cpuTotal = 0

	//gets starting hands
	startHands()

	//Checks if anyone has 21
	playerPoints(playerHandLen)
	cpuPoints(CPUHandLen)

	//both players have 21
	if playerTotal == 21 && cpuTotal == 21 {

		fmt.Print("YOU TIE!\n")

		printHands(playerHandLen, CPUHandLen, gameEND)
		return

		//player has 21
	} else if playerTotal == 21 {
		fmt.Print("YOU WIN!\n")

		printHands(playerHandLen, CPUHandLen, gameEND)
		return

		//cpu has 21
	} else if cpuTotal == 21 {
		fmt.Print("YOU LOSE!\n")

		printHands(playerHandLen, CPUHandLen, gameEND)
		return
	}

	//while the game isn't over
	for gameOver != 1 {

		userInput = displayMenu()

		//user chooses to print hands
		if userInput == "1" {

			printHands(playerHandLen, CPUHandLen, gamePlaying)

			//user chooses to play turn
		} else if userInput == "2" {

			var validInput int = 0

			fmt.Print("Would you like to draw a card? (y/n): ")
			fmt.Scanln(&userInput)
			fmt.Print("\n")

			for validInput != 1 {

				//user chooses to draw a card
				if strings.ToLower(userInput) == "y" {

					playerHandLen = playerTurn(playerHandLen)
					validInput = 1

					//user chooses to stop drawing cards
				} else if strings.ToLower(userInput) == "n" {

					validInput = 1

					//counts CPU's hand
					cpuPoints(CPUHandLen)

					//counts player's hand
					playerPoints(playerHandLen)

					//cpu is done drawing cards
					if cpuTotal > MAX_CPU {
						cpuStop = 1

						//CPU is still drawing cards
					} else {

						for cpuTotal <= MAX_CPU {

							CPUHandLen = cpuTurn(CPUHandLen)
							cpuPoints(CPUHandLen)
						}

						//cpu is over 21
						if cpuTotal > 21 {

							fmt.Print("YOU WIN!\n")
							printHands(playerHandLen, CPUHandLen, gameEND)
							return
						}

						cpuStop = 1
					}

					//cpu also stops drawing, calculates total points
					if cpuStop == 1 {

						//Player won
						if playerTotal > cpuTotal {
							fmt.Print("YOU WIN!\n\n")

							printHands(playerHandLen, CPUHandLen, gameEND)

							//game tied
						} else if playerTotal == cpuTotal {
							fmt.Print("YOU TIE!\n\n")

							printHands(playerHandLen, CPUHandLen, gameEND)

							//Player lost
						} else {
							fmt.Print("YOU LOSE!\n\n")

							printHands(playerHandLen, CPUHandLen, gameEND)
						}

						return
					}

					//invalid input
				} else {
					fmt.Print("INVALID INPUT\n\n")
					fmt.Print("Would you like to draw a card? (y/n): ")
					fmt.Scanln(&userInput)
					fmt.Print("\n")
					validInput = 0
				}
			}

			//calculates the player's points
			playerPoints(playerHandLen)

			//player lost
			if playerTotal > 21 {
				fmt.Print("YOU LOSE\n\n")

				printHands(playerHandLen, CPUHandLen, gameEND)

				return

				//player won
			} else if playerTotal == 21 {
				fmt.Print("YOU WIN\n\n")

				printHands(playerHandLen, CPUHandLen, gameEND)

				return
			}

			//counts CPU's hand
			cpuPoints(CPUHandLen)

			//CPU does not draw more cards if their hand totals over MAX_CPU
			if cpuTotal <= MAX_CPU {

				CPUHandLen = cpuTurn(CPUHandLen)

				//counts CPU's hand
				cpuPoints(CPUHandLen)

				//cpu stops drawing a card
			} else {
				cpuStop = 1
			}

			//player won
			if cpuTotal > 21 {
				fmt.Print("YOU WIN\n\n")

				printHands(playerHandLen, CPUHandLen, gameEND)

				return

				//player lost
			} else if cpuTotal == 21 {
				fmt.Print("YOU LOSE\n\n")

				printHands(playerHandLen, CPUHandLen, gameEND)

				return
			}

			//user chooses to quit game
		} else if userInput == "3" {

			return

		} else {
			fmt.Print("INVALID INPUT\n")
		}
	}
}

func main() {

	var playGame int = 1
	var userInput string

	//while the user wants to play a game
	for playGame == 1 {

		fmt.Print("Would you like to play blackjack? (y/n): ")
		fmt.Scanln(&userInput)
		fmt.Print("\n")

		//starts game
		if strings.ToLower(userInput) == "y" {

			gameStart()

			//exits program
		} else if strings.ToLower(userInput) == "n" {

			playGame = 0
		}
	}
}
