package main

import "fmt"

func main() {
	fmt.Println("Hello, AHA!")

  fmt.Println("--- Now we shall print an array of numbers")

  someNums := []float64{1.2, 3.14, 2.71}
  for _, x := range someNums {
    fmt.Println("The numbers are:", x)
  }
}
