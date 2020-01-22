package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	timeLimit := flag.Int("TimeLimit", 30, "The time limit of each decision round.")
	flag.Parse()

	var x int

	fmt.Println("Welcome to the Go Or No Go, what would you like to do?")
	fmt.Println("1: Play go or no go\t 2: Check the scoreboard")
	fmt.Scanln(&x)

	if x == 1 {

		box := []string{
			"box 1", "box 2", "box 3", "box 4", "box 5", "box 6", "box 7", "box 8", "box 9", "box 10",
		}
		fmt.Println(box)

		//generate a random nonrepeated array to mix the order the box value.
		rand.Seed(time.Now().UnixNano())
		order := rand.Perm(10)

		//Prize list of remaining box value
		Prize := []int{1000, 2500, 5000, 10000, 20000, 30000, 45000, 60000, 80000, 100000}

		//make map dictionary for box
		Value := make(map[string]int)

		//giving value to the boxes
		for i := 0; i < len(box); i++ {
			Value[box[i]] = Prize[order[i]]
		}

		//user's selection of boxes
		var lucky1 int
		var luckybox int

		//select user's lucky box and keep the value.
		fmt.Println("Now you have 10 boxes, please select one box that you would like to keep as your lucky box.")
		fmt.Scanln(&lucky1)
		luckybox = Value[box[lucky1-1]]

		//drop boxes that have been selected
		box[lucky1-1] = " "

		//create a slice to restore all the box that have been selected
		picked := []int{}
		picked = append(picked, lucky1)

		for {
			fmt.Println(box)
			//start to select the boxes to be dropped
			var select1 int
			fmt.Println("please select your next box to be dropped")
			fmt.Scanln(&select1)

			//check if the box has been selected before.
			found := find(picked, select1)
			if found {
				fmt.Println("Pleast do not select the box that has not been dropped.")
			} else {

				//add the dropped box into the trashcan slice
				picked = append(picked, select1)

				//drop boxes that have been selected
				box[select1-1] = " "

				//drop the value from the prize list
				Prize[order[select1-1]] = 0

				//remain number of boxes function
				var remain int
				for k := 0; k < len(box); k++ {
					if Value[box[k]] != 0 {
						remain++
					}
				}

				fmt.Println("Here is the remaining boxes: ")
				fmt.Println(box)

				//total prize pool
				var pool int

				//calculate the total prize pool
				for i := 0; i < len(Prize); i++ {
					pool = pool + Prize[i]
				}

				//show the remaining prize list
				fmt.Println("Here is the total prize pool: ")
				fmt.Println(Prize)

				//banker's offer
				var risk float32 = 0.9
				offer := int(float32(pool) / (float32(remain + 1)) * risk)

				fmt.Println("\nThe banker is offering to buy your lucky box, if you take the offer, the game will end and you can go with the price. Or you can reject the offer and continue the game")
				fmt.Println("You have 30 seconds to make your decision, after 30 seconds, the game is over and your final prize is the value inside your lucky box")
				fmt.Println("\nBanker's offer: ", int32(offer))

				//Decision time, 30 sec. Start to counter right after the offer is given:
				roundTime := time.NewTimer(time.Duration(*timeLimit) * time.Second)

				for i := 0; i < 1; i++ {
					//handle player's choice
					//var decision int
					fmt.Println("\n1: Continue the game\t 2: Take the money and Go")
					decisionCh := make(chan int)

					go func() {
						var decision int
						fmt.Scanln(&decision)
						decisionCh <- decision
					}()

					select {
					//If the player consider over 30 seconds, the game is over and player leave with the luckybox
					//Receive time out signal from roundTime channel first.
					case <-roundTime.C:
						fmt.Println("Ops, your decisioin time ran out, you are now leaving with your lucky box")
						luckyBox(luckybox)
						return

					//The player make his decision on time, the game continue.
					//Receive decision from decisionCh channel first.
					case decision := <-decisionCh:
						if remain == 0 {
							fmt.Println("All the remaining boxese are gone, you can now go with you luckybox")
							luckyBox(luckybox)
							record(decision, offer, luckybox)
							return

						} else if decision == 1 {
							break

						} else if decision == 2 {
							fmt.Println("Accepted banker's offer, here is the prize you earned: ", offer)
							luckyBox(luckybox)
							record(decision, offer, luckybox)
							return

						}

					}
				}
			}
		}
	} else if x == 2 {
		ScoreCheck()

	} else {
		fmt.Println("Please make a selection between 1 and 2")
		fmt.Scanln(&x)
		return
	}

}

//reusable ending quote, to show the value inside the "luckybox"
func luckyBox(luckybox int) {
	fmt.Println("Here is the prize inside your luckybox: ", luckybox)
}

//func find will loop through the slice and check if the number has previously picked and return a boolean result.
func find(drop []int, select1 int) bool {
	for _, picked := range drop {
		if picked == select1 {
			return true
		}
	}
	return false
}

type player struct {
	Name  string
	Score []byte
}

//func record creates a file that has player's nick name and score.
func record(decision int, offer int, luckybox int) {
	var a int
	fmt.Println("Would you like to save your score under a nick name? enter 1 for yes, or any other number to stay annoymous")
	fmt.Scanln(&a)
	if a == 1 {

		var b string
		fmt.Println("Please type your nick name: ")
		fmt.Scanln(&b)

		if decision == 2 {
			player1 := &player{Name: b, Score: []byte(strconv.Itoa(offer))}
			player1.save()
		} else {
			player1 := &player{Name: b, Score: []byte(strconv.Itoa(luckybox))}
			player1.save()
		}
		fmt.Println("You nick name and score has recorded to the file!")

	} else {
		return
	}
}

func (p *player) save() error {
	filename := p.Name + ".txt"
	return ioutil.WriteFile(filename, p.Score, 0600)
}

//func load takes a name and read the content into variable "score"
func load(name string) (*player, error) {
	filename := name + ".txt"
	score, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &player{Name: name, Score: score}, nil
}

//func scorecheck() will take the user's input of player name and call func load() to display the result.
func ScoreCheck() {
	var name string
	fmt.Println("Please enter the name you would like to check: ")
	fmt.Scanln(&name)
	player, err := load(name)
	if err != nil {
		fmt.Println("There is no such player or the system is having problem")
		return
	}
	fmt.Println("Here is the record for: ", name)
	fmt.Println((string(player.Score)))
}
