package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var GameOver bool

func main() {
	var deckOfCards []int = []int{}
	//computer cards
	var player1Cards []int = []int{}
	//player cards
	var player2Cards []int = []int{}
	var isWantCard int

	chooseCard(&deckOfCards, &player1Cards, 1, false)
	chooseCard(&deckOfCards, &player2Cards, 2, true)

	for {
		fmt.Print("kart istiyorsanız 1 istemiyorsanız 0 giriniz : ")
		fmt.Scanf("%d\n", &isWantCard)
		if isWantCard == 0 {
			for {
				if checkTotalIsBigger(player1Cards, player2Cards) {
					break
				}
				chooseCard(&deckOfCards, &player1Cards, 1, false)
				time.Sleep(4000000000)
			}
			if GameOver {
				fmt.Println("Kazandınız")
				break
			} else {
				fmt.Println("Kaybettiniz")
				break
			}
		}
		chooseCard(&deckOfCards, &player2Cards, 1, true)
		if GameOver {
			fmt.Println("Kaybettiniz")
			break
		}
	}
}

func checkTotalIsBigger(player1Cards []int, player2Cards []int) bool {
	var total1 int
	for i := 0; i < len(player1Cards); i++ {
		total1 += player1Cards[i]
	}
	isGameOver(total1)
	var total2 int
	for i := 0; i < len(player2Cards); i++ {
		total2 += player2Cards[i]
	}
	if total1 > total2 || (total2 == 21 && total1 == 21) {
		return true
	}
	return false
}

func printPlayerCardsAndCalculateTotal(playerCards *[]int, isYours bool) {
	var total int
	if isYours {
		fmt.Print("Kartların:")
	} else {
		fmt.Print("Player1'in kartları:")
	}
	for i := 0; i < len(*playerCards); i++ {
		print(strconv.Itoa((*playerCards)[i]) + " ")
		total += (*playerCards)[i]
	}
	isGameOver(total)
	print("\n")
	fmt.Println("total = " + strconv.Itoa(total) + "\n")
}

func isGameOver(total int) {
	if total > 21 {
		GameOver = true
	}
}

///deckOfCards = cards of already choosen
///playerCards = cards in player's hand
func chooseCard(deckOfCards *[]int, playerCards *[]int, numberOfChoosen int, isYourTurn bool) {
	var randomCard int
	var howManySameCardInDeck int
	for i := 0; i < numberOfChoosen; i++ {
		for {
			randomCard = rand.Intn(9) + 1
			rand.Seed(time.Now().UnixNano())
			howManySameCardInDeck = 0
			for j := 0; j < len(*deckOfCards); j++ {
				if randomCard == (*deckOfCards)[j] {
					howManySameCardInDeck += 1
				}
			}
			if howManySameCardInDeck != 4 {
				break
			}
		}
		*deckOfCards = append(*deckOfCards, randomCard)
		*playerCards = append(*playerCards, randomCard)
	}
	printPlayerCardsAndCalculateTotal(playerCards, isYourTurn)
}
