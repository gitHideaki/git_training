// goのFizzBuzz
package main

import "fmt"

func main() {
<<<<<<< HEAD
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
=======
	for (i := 1; i <= 100; i++) {
		if i%15 == 0 {
>>>>>>> parent of 597b686 (code.go)
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}