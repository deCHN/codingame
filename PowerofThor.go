/**
The program must first read the initialization data from the standard input, then, in an infinite loop, provides on the standard output the instructions to move Thor.
Initialization input
Line 1: 4 integers lightX lightY initialTX initialTY. (lightX, lightY) indicates the position of the light. (initialTX, initialTY) indicates the initial position of Thor.
Input for a game round
Line 1: the number of remaining moves for Thor to reach the light of power: remainingTurns. You can ignore this data but you must read it.
Output for a game round
A single line providing the move to be made: N NE E SE S SW W ou NW
Constraints
0 ≤ lightX < 40
0 ≤ lightY < 18
0 ≤ initialTX < 40
0 ≤ initialTY < 18
Response time for a game round ≤ 100ms
**/

package main

import "fmt"

func main() {
	var lightX, lightY, initialTX, initialTY int
	fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)

	move := make(chan string)

	go func(c chan string) {
		for dy, dx := lightY-initialTY, lightX-initialTX; dy != 0 || dx != 0; {
			dir := ""
			switch {
			case dy < 0:
				dir += "N"
				dy += 1
			case dy > 0:
				dir += "S"
				dy -= 1
			}

			switch 
			case dx < 0:
				dir += "W"
				dx += 1
			case dx > 0:
				dir += "E"
				dx -= 1
			}
			c <- dir
		}
	}(move)

	for {
		var remainingTurns int
		fmt.Scan(&remainingTurns)
		fmt.Println(<-move)
	}
}
