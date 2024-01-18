package main

import "fmt"

// func calculateTax(price float32) (float32, float32) {
// 	return price * 0.09, price * 0.02
// }

func birthday(age *int) {
	if *age > 40 {
		panic("Dude you're old")
	}

	*age++
}

func main() {
	// defer println("Hello World!")
	// stateTax, cityTax := calculateTax(5)
	age := 33
	birthday(&age)
	println("Age: ", age)

	// if message := "Hello World"; age == 34 {
	// 	println("You are ", age, message)
	// }

	switch {
	case age == 33:
		println("You are 33")
	default:
		println("You are ", age)
	}

	var collection [10]string
	collection[0] = "Test"

	// for i := 0; i < len(collection); i++ {
	// 	println(collection[i])
	// }

	// for index := range collection {
	// 	println(collection[index])
	// }

	// var i int = 0
	// for i < len(collection) {
	// 	println(collection[i])
	// 	i++
	// }

	// var num1, num2 int
	var operation string

	fmt.Println("Enter an operation")
	fmt.Scanf("%s", &operation)

	switch operation {
	case "add":
		println("You are adding")
	}
}
