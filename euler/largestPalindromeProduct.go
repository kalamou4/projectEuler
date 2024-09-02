package euler

import (
    "strconv"
    "math"
)

// The function `LargestPalindromProduct` finds the largest palindrome product of two numbers with a
// specified number of digits.
func LargestPalindromProduct(digits int) int {
    max := int(math.Pow10(digits)) - 1
    min := int(math.Pow10(digits - 1))
    largestPalindrome := 0

    for i := max; i >= min; i-- {
        for j := i; j >= min; j-- { // Start from i to avoid duplicate calculations
            product := i * j
            if product <= largestPalindrome {
                // No need to continue this inner loop if product is smaller
                break
            }
            if IsPalindrome(strconv.Itoa(product)) && product > largestPalindrome {
                largestPalindrome = product
            }
        }
    }
    return largestPalindrome
}

func IsPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

//arithmetic approch for checking palindromes
// func IsPalindrom(n int) bool {
//     backwards, temp := 0, n
//     for temp > 0 {
//         backwards = backwards * 10 + temp % 10
//         temp /= 10
//     }
//     return backwards == n
// }