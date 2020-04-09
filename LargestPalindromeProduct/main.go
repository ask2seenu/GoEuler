/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
)

func main() {


	for i,j=0; i,j<=1000; i,j++; {
		if i*j == reverse(i*j){

			fmt.Println(i*j)
			
		}

	}
}
