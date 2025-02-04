package main

import "fmt"

func main() {
	var n int = 10
	var leg_height int = n / 2
	var leg rune = '🎃'
	var leaf rune = '🌱'
	var sky rune = '🔵'
	var soil rune = '🟫'
	for i := 0; i < int(n+leg_height); i++ {
		for j := 1; j < int(n*2); j++ {
			if j >= n-i && j <= n+i && i < n {
				fmt.Printf("%c", leaf)
			} else if j == n {
				fmt.Printf("%c", leg)
			} else if i > n {
				fmt.Printf("%c", soil)
			} else {
				fmt.Printf("%c", sky)
			}
		}
		println()
	}
}
