package main

import (
	"fmt"
	"math"
)

/**
* Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
  **/

const (
	width  int = 16000
	height int = 9000
)

func main() {
	threshold := 0.5 * math.Sqrt(width*width+height*height)
	var thrust int = 100
	var speed int = 0

	for {
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		if nextCheckpointAngle > 90 || nextCheckpointAngle < -90 {
			thrust = 0
		} else {
			thrust = int(100 * math.Cos(float64(nextCheckpointAngle)/180*math.Pi))
		}

		if nextCheckpointAngle == 0 && float64(nextCheckpointDist) > threshold {
			fmt.Printf("%d %d %s\n", nextCheckpointX, nextCheckpointY, "BOOST")
		}

		fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY, thrust)
	}
}
