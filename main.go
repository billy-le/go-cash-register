package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getEasyInt(val float64) int32 {
	ratio := math.Floor(val * 100)
	return int32(ratio)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

outer:
	for {
		fmt.Print("What is the total cost of the items? ")

		scanner.Scan()

		costInput := scanner.Text()
		if len(costInput) == 0 {
			break
		}

		costValue, err := strconv.ParseFloat(costInput, 64)
		switch {
		case err != nil:
			fmt.Println("Bad input. Please enter a number")
		case costValue <= 0:
			fmt.Println("Please enter a number greater than 0")
		default:
		getChange:
			fmt.Print("What is the given amount? ")

			scanner.Scan()

			givenInput := scanner.Text()

			if len(givenInput) == 0 {
				goto getChange
			}

			givenValue, err := strconv.ParseFloat(givenInput, 64)

			switch {
			case err != nil:
				fmt.Println("Bad input. Please enter a number")
				goto getChange
			case givenValue <= 0:
				fmt.Println("Please enter a number greater than 0")
				goto getChange
			case costValue > givenValue:
				fmt.Println("The cost is greater than the amount given")
				break outer
			default:
				value := getEasyInt(givenValue - costValue)
				change := float64(value) / 100

				changeMap := map[string]int32{
					"100 dollars": 0,
					"50 dollars":  0,
					"20 dollars":  0,
					"10 dollars":  0,
					"5 dollars":   0,
					"1 dollars":   0,
					"quarters":    0,
					"dimes":       0,
					"nickels":     0,
					"pennies":     0,
				}

				for value > 0 {
					switch {
					case value >= 10000:
						changeMap["100 dollars"] += 1
						value -= 10000
					case value >= 5000:
						changeMap["50 dollars"] += 1
						value -= 5000
					case value >= 2000:
						changeMap["20 dollars"] += 1
						value -= 2000
					case value >= 1000:
						changeMap["10 dollars"] += 1
						value -= 1000
					case value >= 500:
						changeMap["5 dollars"] += 1
						value -= 500
					case value >= 100:
						changeMap["1 dollars"] += 1
						value -= 100
					case value >= 25:
						changeMap["quarters"] += 1
						value -= 25
					case value >= 10:
						changeMap["dimes"] += 1
						value -= 10
					case value >= 5:
						changeMap["nickels"] += 1
						value -= 5
					case value >= 1:
						changeMap["pennies"] += 1
						value -= 1
					default:
						value -= 1
					}
				}

				fmt.Printf("\nYour change is: $%.2f\n\n", change)
				fmt.Printf("(%d) 100 dollar bills\n", changeMap["100 dollars"])
				fmt.Printf("(%d) 50 dollar bills\n", changeMap["50 dollars"])
				fmt.Printf("(%d) 20 dollar bills\n", changeMap["20 dollars"])
				fmt.Printf("(%d) 10 dollar bills\n", changeMap["10 dollars"])
				fmt.Printf("(%d) 5 dollar bills\n", changeMap["5 dollars"])
				fmt.Printf("(%d) 1 dollar bills\n", changeMap["1 dollars"])
				fmt.Printf("(%d) quarters\n", changeMap["quarters"])
				fmt.Printf("(%d) dimes\n", changeMap["dimes"])
				fmt.Printf("(%d) nickels\n", changeMap["nickels"])
				fmt.Printf("(%d) pennies\n\n", changeMap["pennies"])
			}
		}
	}
}
