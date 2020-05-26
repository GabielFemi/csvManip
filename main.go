package main

import (
	"fmt"
	gabby_csv "github.com/gabielfemi/csvManip/csv"
)

func main() {
	fmt.Println("Hi there, This is CSV manip!")
	csvData := gabby_csv.ReadCsv("demo.csv")
	fmt.Println(csvData)
}
