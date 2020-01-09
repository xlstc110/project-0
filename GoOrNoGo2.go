package main

import "fmt"


func main(){
	var x int

	fmt.Println(" ")
	fmt.Println("Welcome to the Go Or No Go, what would you like to do?")
	fmt.Println(" ")
	fmt.Println("1: Play go or no go\t 2: Check the scoreboard")
	fmt.Scanln(&x)

	if x == 1{
	
		fmt.Println(" ")

		box := []string{
		"box 1","box 2","box 3","box 4","box 5","box 6","box 7","box 8","box 9","box 10",
		}
		fmt.Println(box)
		fmt.Println(" ")
	
		//make map dictionary for box
		Value := make(map[string]int)

		//base value of prize pool
		ValueBase := 2

		//giving value to the boxes
		for i := 0; i <len(box); i++{
			Value[box[i]] = ValueBase<<i
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

		

		fmt.Println(" ")

		for {

			fmt.Println(" ")
			
			fmt.Println(box)
			//start to select the boxes to be dropped
			var select1 int
			fmt.Println("please select your next box to be dropped")
			fmt.Scanln(&select1)

			//drop boxes that have been selected
			box[select1-1] = " "

			//clear the value of seleted boxes from the prize pool
			Value[box[select1-1]] = 0


			//remain number of boxes function
			var remain int
			for k := 0; k < len(box); k++{
				if Value[box[k]] != 0{
					remain++
				}
			}
			fmt.Println(" ")
			fmt.Println("Here is the total remaining boxes: ")
			fmt.Println(box)

			//total prize pool
			var pool int

			//calculate the total prize pool
			for j := 0; j < len(box); j++{
				pool = pool + Value[box[j]]
			}
			fmt.Println("Here is the total prize pool amount: ")
			fmt.Println(pool)
			//banker's offer
			offer := pool/(remain+1)
			fmt.Println("The banker is offering to buy your lucky box, if you take the offer, the game will end and you can go with the price. Or you can reject the offer and continue the game")
			fmt.Println(offer)

			var y int
			fmt.Println("1: Continue the game\t 2: Take the money and Go")
			fmt.Scanln(&y)

			if remain == 0{
				fmt.Println("All the remaining boxese are gone, you can now go with you luckybox")
				fmt.Println("Here is your luckybox's value!")
				fmt.Println(luckybox)
				break
		
			} else if y == 1 {
			
			} else if y == 2{
				fmt.Println("Accepted banker's offer, here is the prize you earned: ")
				fmt.Println(offer)
				fmt.Println("Here is the prize inside your luckybox")
				fmt.Println(luckybox)
				break
			}
		}
	} else {
			fmt.Println("Ops, the scoreboard is empty for now, bye~")
		}
		
}



	
