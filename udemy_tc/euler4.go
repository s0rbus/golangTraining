package main

//Golang solution to Project Euler Problem 4:
//A palindromic number reads the same both ways. The largest palindrome made
//from the product of two 2-digit numbers is 9009 = 91 × 99.
//
//Find the largest palindrome made from the product of two 3-digit numbers.
//

import (
   "fmt"
   "strconv"
)

//Convert int to string then iterate from both ends, checking for equality
//Only have to iterate to half way. If even length then each half
//has been checked. If odd length the middle value will equal itself
func palcheck(i int) bool {
   s := strconv.Itoa(i)
   l := len(s)
   r := []rune(s)
   for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
      if r[i] != r[j] {
         return false
      }
   }
   return true
}

func main() {
   //Checking multiples of 3 digit numbers and want max so count down from 999*999
   //and limit is 100
   lim := 100
   max := lim * lim
   st := lim*10 - 1
   var mi, mj int
   for i := st; i >= lim; i-- {
      //quick check here to see if already reached max possible
      if max >= (i * st) {
         break
      }
      for j := st; j >= i; j-- {
         //check for palindrome and if a new max
         if palcheck(i*j) && ((i * j) >= max) {
            max = i * j
            mi, mj = i, j
         }
      }
   }
   fmt.Printf("Max palindrome is %d = %d * %d\n", max, mi, mj)
}
