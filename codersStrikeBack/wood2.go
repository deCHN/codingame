package main

import (
	"fmt"
	"math"
	"os"
)

/**
* Auto-generated code below aims at helping you parse
* the standard input according to the problem statement.
**/

const (
	width  float64 = 16000
	height float64 = 9000
)

func main() {
	//If the next checkpoint is cross half of the map, we boost it.
	threshold := 0.5 * math.Sqrt(width*width+height*height)
	var thrust, remainboost int = 100, 1

	var speed, lastDist = 0, 0

	for {
		// nextCheckpointX: x position of the next check point
		// nextCheckpointY: y position of the next check point
		// nextCheckpointDist: distance to the next checkpoint
		// nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		if nextCheckpointAngle > 90 || nextCheckpointAngle < -90 {
			thrust = 0
		} else {
			//thrust = int(math.Abs(100 * math.Cos(float64(nextCheckpointAngle)/180*math.Pi)))
			thrust = 100
		}

		if remainboost > 0 && nextCheckpointAngle == 0 && float64(nextCheckpointDist) > threshold {
			fmt.Printf("%d %d %s\n", nextCheckpointX, nextCheckpointY, "BOOST")
			remainboost--
		}

		speed = lastDist - nextCheckpointDist
		lastDist = nextCheckpointDist

		if nextCheckpointDist < 2*600+speed {
			thrust = 30
		}

		fmt.Fprintln(os.Stderr, "dist:", nextCheckpointDist, "angle:", nextCheckpointAngle, "thrust:", thrust, "boost:", remainboost, "speed:", speed)

		fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY, thrust)
	}
}
