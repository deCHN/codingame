package main

import "fmt"

/*
Within an infinite loop, read the heights of the mountains from the standard input and print to the standard output the index of the mountain to shoot.
Input for one game turn 8 lines: one integer mountainH per line. Each represents the height of one mountain given in the order of their index (from 0 to 7).

Output for one game turn
A single line with one integer for the index of which mountain to shoot.

Constraints
0 ≤ mountainH ≤ 9
Response time per turn ≤ 100ms
*/

//import "os"

func main() {
	for {
		mountainH, t, shoot := 0, 0, 0

		for i := 0; i < 8; i++ {
			// mountainH: represents the height of one mountain, from 9 to 0.
			fmt.Scan(&mountainH)
			if t < mountainH {
				t = mountainH
				shoot = i
			}
		}

		fmt.Println(shoot) // The number of the mountain to fire on.
	}
}
