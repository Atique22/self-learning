package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// Fibonacci Sequence: Write a function that calculates the Fibonacci sequence up to a given number n and returns the sequence as an array or slice.
func fibonacciSequence(n int) {
	first := 0
	second := 1
	fmt.Print("\n Fibonacci sequence: ", first, ", ", second)
	for i := 2; first+second < n; i++ {
		next := first + second
		fmt.Print(", ", next)
		first = second
		second = next
	}
}

// Prime Numbers: Write a function that determines whether a given number n is prime or not.
func primeNumbers(n int) bool {

	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i < limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Palindrome Check: Write a function that checks whether a given string is a palindrome (reads the same forwards and backwards)
func palindromeCheck(str string) bool {
	// Convert the string to a byte array
	bytes := []byte(str)

	length := len(bytes)
	for i := 0; i < length/2; i++ {
		// Swap the characters at i and length-i-1 indices
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
	// Convert the byte array back to a string
	reversedStr := string(bytes)

	if str == reversedStr {
		return true
	}
	return false
}

// Reverse String: Write a function that reverses a given string without using the built-in reverse function.
func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	// Iterate over the rune slice and swap characters
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		// Swap the characters at i and j indices
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the rune slice back to a string
	reversedStr := string(runes)

	return reversedStr
}

// Factorial: Write a function that calculates the factorial of a given number n.
func factorial(n int) int {
	res := 1
	if n == 0 {
		return res
	} else {
		for i := 1; i <= n; i++ {
			res = res * i
		}
	}

	return res
}

// Find Duplicates: Write a function that finds and returns any duplicates in a given array or slice.
func duplicates(arr []int, size int) []int {
	duplicates := []int{}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				duplicates = append(duplicates, arr[i])
			}
		}
	}
	return duplicates
}

// Binary Search: Implement a binary search algorithm that searches for a given target value in a sorted array.
func binarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2

		if array[mid] == target {
			return mid // Target found at index mid
		} else if array[mid] < target {
			left = mid + 1 // Target is in the right half
		} else {
			right = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found in the array
}

// Anagram Check: Write a function that checks whether two given strings are anagrams (contain the same characters in a different order).
func areAnagrams(str1, str2 string) bool {
	// Remove spaces and convert strings to lowercase
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Convert strings to slices of characters
	slice1 := strings.Split(str1, "")
	slice2 := strings.Split(str2, "")

	// Sort the slices
	sort.Strings(slice1)
	sort.Strings(slice2)

	// Convert the sorted slices back to strings
	sortedStr1 := strings.Join(slice1, "")
	sortedStr2 := strings.Join(slice2, "")

	// Check if the sorted strings are equal
	return sortedStr1 == sortedStr2
}

// Word Count: Write a function that counts the occurrence of each word in a given string and returns the word count as a map.

// Matrix Operations: Implement functions to perform basic matrix operations such as matrix addition, multiplication, and transposition.

func main() {

	var input int
	for {
		fmt.Print("\n 1.Fibonacci Sequence: \n 2.Prime Numbers: \n 3.Palindrome Check: \n 4.Reverse String: \n 5.Factorial:\n 6.Find Duplicates:\n 7.Binary Search:\n 8.Anagram Check: \n 9.same characters in a different order\n 10.Word Count: \n 11.Matrix Operations: \n Other.Break this Program: \n\n Enter a number for problem give in sequence: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Print("\n\n 1.Fibonacci Sequence: \n\n To calculates the Fibonacci sequence up to a given number n, enter n number: ")
			var n int
			fmt.Scan(&n)
			fibonacciSequence(n)

		case 2:
			fmt.Print("\n\n 2.Prime Numbers: \n\n To calculates whether a given number n is prime or not., enter n number: ")
			var n int
			fmt.Scan(&n)
			primeNumbers(n)
			if primeNumbers(n) {
				fmt.Println("  ", n, "is a prime number.")
			} else {
				fmt.Println("  ", n, "is not a prime number.")
			}
		case 3:
			fmt.Println(" 3.Palindrome Check: \n\n To Check Palindrome, enter string: ")
			var str string
			fmt.Scan(&str)
			palindromeCheck(str)
			if palindromeCheck(str) {
				fmt.Println("  ", str, "is Palindrome")
			} else {
				fmt.Println("  ", str, "is No Palindrome")
			}
		case 4:
			fmt.Print("\n\n4.Reverse String:\n\n Enter string: ")
			var str string
			fmt.Scan(&str)
			myStr := reverseString(str)
			fmt.Println("\nReverse string is: ", myStr)

		case 5:
			fmt.Print("\n\n5.Factorial: \n\n Enter n number: ")
			var n int
			fmt.Scan(&n)
			res := factorial(n)
			fmt.Print("Factorial is :", res)

		case 6:
			fmt.Println("6.Find Duplicates:")
			var size int

			fmt.Print("Enter the size of the array: ")
			fmt.Scanln(&size)

			// Create an array with the specified size
			array := make([]int, size)

			// Prompt the user to enter values for each element
			for i := 0; i < size; i++ {
				fmt.Printf("Enter element at index %d: ", i)
				fmt.Scan(&array[i])
			}
			dpArr := duplicates(array, size)
			fmt.Println("Duplicates:", dpArr)

		case 7:
			fmt.Println("7.Binary Search: \n\n Enter target number: ")
			array := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
			var target int
			fmt.Scan(&target)
			result := binarySearch(array, target)
			if result == -1 {
				fmt.Println("Target not found in the array.")
			} else {
				fmt.Printf("Target found at index %d.\n", result)
			}
		case 8:
			fmt.Println("8.Anagram Check: ")
			var str1, str2 string
			fmt.Println("Enter string1: ")
			fmt.Scan(&str1)
			fmt.Println("Enter string1: ")
			fmt.Scan(&str2)

			if areAnagrams(str1, str2) {
				fmt.Printf("%s and %s are anagrams.\n", str1, str2)
			} else {
				fmt.Printf("%s and %s are not anagrams.\n", str1, str2)
			}
		case 9:
			fmt.Println("9.same characters")
		case 10:
			fmt.Println("10.Word Count:")
		case 11:
			fmt.Println("11.Matrix Operations:")
		default:
			os.Exit(0)
		}
	}
}
