package game_of_life

import "fmt"
import "overview/game_of_life/dashboard"

func main() {
	d := dashboard.New(5, 5)
	fmt.Println(d.Width())
	fmt.Println(*d)
}
