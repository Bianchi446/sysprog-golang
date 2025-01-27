package main

import "fmt"

func customSwitch(tier, age int) {
	switch tier {
	case 1:
		fmt.Println("T-shirt")
		if age > 18 {
			break
		}
		fallthrough
	case 2:
		fmt.Println("Smug")
		fallthrough
	case 3:
		fmt.Println("Sticker pack")
	default:
		fmt.Println("No reward")
	}
}

func looping(b int) {
label:
	for i := b; i < b+2; i++ {
		switch i % 3 {
		case 0:
			fmt.Println("Divisible by 3")
			break label
		default:
			fmt.Println("Not divisible by 3")
		}
	}
}
