/*
this a program that prints the hailstone sequence for a given number
*/
package util

func Hailstone(n int) []int {
	var sequence []int
	for n != 1 {
		sequence = append(sequence, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return sequence
}
