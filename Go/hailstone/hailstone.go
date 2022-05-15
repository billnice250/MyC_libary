/*
this a program that prints the hailstone sequence for a given number
*/
package hailstone

func Generate(n int64) []int64 {
	var sequence []int64
	if n == 1 {
		sequence = append(sequence, n)
		return sequence
	}
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
