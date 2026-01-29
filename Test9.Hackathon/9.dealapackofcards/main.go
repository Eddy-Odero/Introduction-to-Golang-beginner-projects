package main
import "fmt"

func DealAPackOfCards (deck []int) {
	player :=1
	for i := 0; i+2 < len(deck); i += 3 {
	fmt.Printf("Player %d: %d, %d, %d$\n",
			player, deck[i], deck[i+1], deck[i+2])
		player++

}

}
func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
