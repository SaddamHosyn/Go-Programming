package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '0'; i-- {
		for j := i - 1; j >= '0'; j-- {
			for k := j - 1; k >= '0'; k-- {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				if !(i == '2' && j == '1' && k == '0') { // Avoid trailing comma and space after "210"
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

/*Instructions
Write a program that prints in descending order on a single line all unique combinations of three different digits so that the first digit is greater than the second and the second is greater than the third.

These combinations are separated by a comma and a space.

Usage
Here is an incomplete output :

$ go run . | cat -e
987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
$
999 or 000 are not valid combinations because the digits are not different.

789 should not be shown because the first digit is not greater than the second.*/
