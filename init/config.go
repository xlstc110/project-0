package domain

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"strconv"
	"text/template"
)

//Server alive for score checking.
func webPage() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe(":8080", nil)
}

//read content and return *template.Template
var templates = template.Must(template.ParseFiles("mainpage.html", "result.html"))

//execute template mainpage.html
func mainPageHandler(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "mainpage.html", nil)
}

//handle request and execute template result.html
func resultHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	score, err := load(name)
	if err != nil {
		score1 := &player{Name: name, Score: []byte("No record for this player")}
		score = score1
	}

	templates.ExecuteTemplate(w, "result.html", score)
}

//Player did not make a decision in 30 seconds, box stage is over.
//Show luckybox then create a default decision variable and call record func()
func timesUp(luckybox int, offer int) {
	log.Println("Game phrase: player's time up, game over and open the lucky box as final prize.")

	fmt.Println("Ops, your decisioin time ran out, you are now leaving with your lucky box")
	luckyBox(luckybox)
	decision := 0
	record(decision, offer, luckybox)
}

//Player accepted offer, box stage is over
//Show luckybox then call record func()
func acceptOffer(luckybox int, decision int, offer int) {
	log.Println("Game phrase: player accepted the offer, game over and open the lucky box.")

	fmt.Println("Accepted banker's offer, here is the prize you earned: ", offer)
	luckyBox(luckybox)
	record(decision, offer, luckybox)
}

//All boxes are picked, box stage is over
//Show luckybox then call record func()
func boxGone(luckybox int, decision int, offer int) {
	log.Println("Game phrase: all the boxes on stage is gone, game over and open the lucky box.")

	fmt.Println("All the staging boxes are gone, you can now go with you luckybox")
	luckyBox(luckybox)
	record(decision, offer, luckybox)
}

//reusable ending quote, to show the value inside the "luckybox"
func luckyBox(luckybox int) {
	fmt.Println("Here is the prize inside your luckybox: ", luckybox)
}

//func find loop through the slice and check if the number has previously picked and return a boolean result.
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
	log.Println("Record stage: if the player wants to save his record.")

	var a int
	fmt.Println("Would you like to save your score under a nick name? enter 1 for yes, or any other number to stay annoymous")
	fmt.Scanln(&a)
	if a == 1 {
		log.Println("Record stage: player type the nick name to be saved.")

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
		fmt.Println("You can check your or other's record on the site http://localhost:8080/ or in main menu!")

	} else {
		log.Println("Record stage: player stay annoymous.")
		return
	}
}

//func save create a txt file that has player's name and score
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

//ScoreCheck will take the user's input of player name and call func load() to display the result.
func ScoreCheck() {

	var name string
	fmt.Println("Please enter the name you would like to check: ")
	fmt.Scanln(&name)
	player, err := load(name)
	if err != nil {
		fmt.Println("There is record for this player")
		return
	}
	fmt.Println("Here is the record for: ", name)
	fmt.Println((string(player.Score)))

	log.Println("Record stage: an user checked player record for: ." + name)
}
