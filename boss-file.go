package main

import (
	"fmt"

	L "./lib"
)

func main() {
	//var inputFile = "./material/sneaker-numbers.csv"
	csvList := L.ReadInputCsv("./material/sneaker-numbers.csv")
	fmt.Println(csvList)
}
