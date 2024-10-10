package main

import "fmt"

func main() {
	var myFavnum float64 = 1.614
	var uniqueNum int = 10
	fmt.Printf("myFavnum: %v\nmyUniqueNum: %v\n", myFavnum, uniqueNum)

	var (
		num    int    = 13
		wrd    string = "A frog"
		status string = "submitted"
	)
	fmt.Print(num, wrd, status, "\n")
	var a, b, c, d int = 2, 4, 6, 8
	fmt.Print("int line: ", a, b, c, d)
}
