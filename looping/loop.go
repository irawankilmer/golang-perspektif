package looping

import "fmt"

func LoopSatu(ulang int) {
	for i := 1; i <= ulang; i++ {
		if i%2 == 1 {
			fmt.Printf("%d. I Love Coding\n", i+1)
			continue
		}
	}
}

func LoopDua() {
	for i := 20; i >= 1; i-- {
		if i%2 == 0 {
			fmt.Printf("%d. I Will Become a FullStack Developer\n", i)
			continue
		}
	}
}

func LoopTiga() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 == 0 {
			fmt.Printf("%d. Super Bootcamp\n", i)
		} else if i%3 == 0 && i%2 == 1 {
			fmt.Printf("%d. I Love Coding\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d. Berkualitas\n", i)
		} else {
			fmt.Printf("%d. Santai\n", i)
		}
	}
}

func LoopEmpat() {
	for i := 0; i < 7; i++ {
		for a := 0; a <= i; a++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
