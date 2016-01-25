package main

import "fmt"

func f(origem string, canal chan string) {
	for i := 0; i < 100; i++ {
		fmt.Println(origem, ":", i)
		if i == 50 {
			canal <- "\n ######## " + origem + " chegou a 50 \n"
		}
	}
}

func main() {
	canal := make(chan string, 1)
	go f("goroutine 1", canal)
	go f("goroutine 2", canal)
	go func(msg string) {
		fmt.Println(msg)
	}("Olá")

	fmt.Println(<-canal)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
