/** Number Guessing Game
What to learn
-- Random Numbers
-- Loops
--Conditional
--user input
--error handling basics

Let's Goooo
**/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	arr := [4][3]int{
		{10, 80, 45},
		{20, 37, 22},
		{1, 5, 78},
		{99, 34, 2},
	}

	point := 0

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println("Numbers to guess", arr[i][j])
		}
		randomNum := arr[i][rand.Intn(len(arr[i]))]
		fmt.Print("Enter a number from the lists:\n")
		scanner.Scan()
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Not correct digit:", err)
			break
		}

		if num == randomNum {
			point += 5
		}
	}
	fmt.Printf("You got %d points out of 20 \n", point)
}
