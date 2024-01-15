package main

// func calculateTax(price float32) (float32, float32) {
// 	return price * 0.09, price * 0.02
// }

func birthday(age *int) {
	*age++
}

func main() {
	// stateTax, cityTax := calculateTax(5)
	age := 33
	birthday(&age)
	println("Age: ", age)
}
