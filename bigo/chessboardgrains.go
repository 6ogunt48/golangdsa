package main

import "fmt"

/* Imagine you have a chessboard and put a single grain of rice on one square. On the second square, you put 2
grains of rice, since that is double the amount of rice on the previous square. On the third square, you put 4
grains. On the fourth square, you put 8 grains, and on the fifth square, you put 16 grains, and so on.
The following function calculates which square you’ll need to place a certain number of rice grains.
For example, for 16 grains, the function will return 5, since you will place the 16 grains on the fifth square.
*/

func chessBoardSpace(numberofGrains int) int {
	chessBoardSpaces := 1
	placedGrains := 1

	for placedGrains < numberofGrains {
		placedGrains *= 2
		chessBoardSpaces += 1
	}
	return chessBoardSpaces

}

func main() {
	noOfgrains := chessBoardSpace(35)
	fmt.Printf("No of grains: %v\n", noOfgrains)
}

// Time complexity logarithmic time operation
