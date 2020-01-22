package Player

import(
	"fmt"
	"io/ioutil"
)

type Player struct {
	Name  string
	Score []byte
}

func (p *Player) Save() error {
	filename := p.Name + ".txt"
	return ioutil.WriteFile(filename, p.Score, 0600)
}

//func load takes a name and read the content into variable "score"
func Load(name string) (*Player, error) {
	filename := name + ".txt"
	score, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Player{Name: name, Score: score}, nil
}

//func scorecheck() will take the user's input of player name and call func load() to display the result.
func ScoreCheck() {
	var name string
	fmt.Println("Please enter the name you would like to check: ")
	fmt.Scanln(&name)
	player, err := Load(name)
	if err != nil {
		fmt.Println("There is no such player or the system is having problem")
		return
	}
	fmt.Println("Here is the record for: ", name)
	fmt.Println((string(player.Score)))
}
