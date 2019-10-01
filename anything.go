package main

import "fmt"

func main() {
	fmt.Printf("Todd Macleod is the best teacher ever!")
	foo()
	fmt.Println("Print something more baby!")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I am in foo baby!")
}

func bar() {
	fmt.Println("Then we go to the bar and exited baby!")
}
