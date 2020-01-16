package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main(){
	var x int

	fmt.Println("\nWelcome to the Go Or No Go, what would you like to do?")
	fmt.Println("\n1: Play go or no go\t 2: Check the scoreboard\n")
	fmt.Scanln(&x)

	if x == 1{

		box := []string{
		"box 1","box 2","box 3","box 4","box 5","box 6","box 7","box 8","box 9","box 10",
		}
		fmt.Println(box,"\n")

		//generate a random nonrepeated array to mix the order the box value.
		rand.Seed(time.Now().UnixNano())
		order := rand.Perm(10)
		
		//Prize list of remaining box value
		Prize := []int {1000,2500,5000,10000,20000,30000,45000,60000,80000,100000,}
	
		//make map dictionary for box
		Value := make(map[string]int)



		//giving value to the boxes
		for i := 0; i <len(box); i++{
			Value[box[i]] = Prize[order[i]]
		}
		//fmt.Println(Value) cheat sheet
		//user's selection of boxes
		var lucky1 int
		var luckybox int
	
		//select user's lucky box and keep the value.
		fmt.Println("Now you have 10 boxes, please select one box that you would like to keep as your lucky box.")
		fmt.Scanln(&lucky1)
		luckybox = Value[box[lucky1-1]]

		//drop boxes that have been selected
		box[lucky1-1] = " "
		
		for {
			fmt.Println(box)
			//start to select the boxes to be dropped
			var select1 int
			fmt.Println("\nplease select your next box to be dropped")
			fmt.Scanln(&select1)

			//drop boxes that have been selected
			box[select1-1] = "   "

			//drop boxes that have been selected
			Prize[select1-1] = 0

			//clear the value of seleted boxes from the prize pool
			Value[box[select1-1]] = 0


			//remain number of boxes function
			var remain int
			for k := 0; k < len(box); k++{
				if Value[box[k]] != 0{
					remain++
				}
			}

			fmt.Println("\nHere is the remaining boxes: \n")
			fmt.Println(box)

			//total prize pool
			var pool int

			//calculate the total prize pool
			for j := 0; j < len(box); j++{
				pool = pool + Value[box[j]]
			}
	
			fmt.Println("\nHere is the total prize pool: \n", Prize)
			
			//banker's offer
			offer := pool/(remain+1)
	
			fmt.Println("\nThe banker is offering to buy your lucky box, if you take the offer, the game will end and you can go with the price. Or you can reject the offer and continue the game")
			fmt.Println("\nBanker's offer: ",offer)

			var y int
			fmt.Println("\n1: Continue the game\t 2: Take the money and Go")
			fmt.Scanln(&y)

			if remain == 0{
				fmt.Println("All the remaining boxese are gone, you can now go with you luckybox")
				fmt.Println("\nHere is your luckybox's value!: ", luckybox)
				break
		
			} else if y == 1 {
			
			} else if y == 2{
				fmt.Println("\nAccepted banker's offer, here is the prize you earned: ",offer)
				fmt.Println("\nHere is the prize inside your luckybox: ",luckybox)
				break
			}	//else if y == selection pool array {
				// fmt.Println("Please select the box numbers inside the pool")
			//}
		}
	} else if x == 2{
			fmt.Println("Ops, the scoreboard is empty for now, bye~")
	} else {
		fmt.Println("Please make a selection between 1 and 2")
	}

	// func droppedbox(){
	// 	dropped := []string

	// }

	// func restore(){
	// 	var r int
	// 	fmt.Println("Here is your first reborn game, the box you select will be restore to the pool")
	// 	fmt.Println(boxdropped)
	// 	fmt.Scanln(&r)
	//  restore the selected box back to the pool
	// }
	
}