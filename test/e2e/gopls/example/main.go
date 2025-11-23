package main

// Add adds two integers and returns the result.
// This function demonstrates basic arithmetic operation.
//
// Example usage:
//
//	result := Add(5, 3)
//	fmt.Println(result) // Output: 8
func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(5, 3)
	println(result)
}
