package main

import (
	"fmt"
)

type Team struct {
	Name  string
	score int
}

func changescore(s *Team, score int) {
	// Modify the Person's age directly using the pointer
	s.score = score
}

func main() {
	var P1 = Team{"Aus", 329}
	fmt.Println(P1)
	changescore(&P1, 567)
	fmt.Println(P1)

}
