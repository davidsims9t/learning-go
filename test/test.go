package test

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]bool

func init() {
	Countries[0] = "United States"
	Countries[1] = "South Korea"

	qty := cap(Countries)
	fmt.Println("Countries: ", qty)

	names := []string{"Mary", "John"}
	names = append(names, "Carol")
	println(len(names))
}

const Boo = "hoo"
