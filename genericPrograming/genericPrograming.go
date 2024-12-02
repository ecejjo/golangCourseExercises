package main

import (
	"fmt"
	"strings"
)

import "golang.org/x/exp/constraints"


// AnyOf checks if any element in the slice satisfies the given predicate.
func AnyOf[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

func FindIf[T any] (slice []T, predicate func (T) bool) (int, bool) {
	for index, item := range slice {
		if predicate(item) {
			return index, true
		}
	}
	return -1, false
}

func AdjacentFind[T any] (slice []T, predicate func (T, T) bool) int{
	for i := 0; i < len(slice)-1; i++ {
		if predicate(slice[i], slice[i+1]) {
			return i
		}
	}
	return -1	
}

func Equal[T comparable] (a[]T, b[]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Notice: if only chaning an element in the slice, we can keep using slice by value.
func ReplaceIf[T comparable] (slice[]T, element T, compare func (T) bool) (int, bool){
	for i := 0; i < len(slice); i++ {
		if compare(slice[i]) {
			slice[i] = element
			return i, true
		}
	}
	return -1, false
}

// Notice: as here we are modifying the slice itselft (and not only an element), wee need to pass slice by reference.
func RemoveIf[T comparable] (slice*[]T, compare func (T) bool) (int, bool) {
	for i := 0; i < len(*slice); i++ {
		if compare((*slice)[i]) {
			if (i==0) {
				fmt.Println("Removing index 0 ...")
				*slice = (*slice)[i+1:]
			} else {
				*slice = append((*slice)[:i-1], (*slice)[i+1:]...)
			}
			fmt.Println("New slice is:", slice)
			return i, true
		}
	}
	return -1, false
}

// IsSorted checks if a slice of any ordered type is sorted in ascending order.
func IsSorted[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

func Merge[T any] (slice1[]T, slice2[]T) []T {
	var returnSlice []T
	if (len(slice1) == 0) {
		return slice2
	} else if (len(slice2) == 0) {
		return slice1
	} else if (len(slice1) < len(slice2)) {
		for i := 0; i < len(slice1); i++ {
			returnSlice = append(returnSlice, slice1[i])
			returnSlice = append(returnSlice, slice2[i])
		}
		for i := len(slice1); i < len(slice2)-len(slice1); i++ {
			returnSlice = append(returnSlice, slice2[i])
		}
	} else if (len(slice1) < len(slice2)) {
		returnSlice = Merge(slice2, slice1)
	}
	return returnSlice
}


// Example usage
func main() {
	
	consecutiveNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nonConsecutiveNumbers := []int{1, 8, 90, 12, 50, 66, 73, 28, 19, 110}

	evenNumbers := []int{2, 4, 6}
	oddNumbers := []int{1, 3,  5, 7, 9}

	withConsecutiveLetters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	withoutConsecutiveLetters := []rune{'k', 'b', 'm', 'h', 'o', 's', 'w', 'z', 'c', 't'}

	isEven := func(n int) bool { return n%2 == 0 }

	fmt.Println("Testing AnyOf isEven on evenNumbers: ")
	fmt.Println(AnyOf(evenNumbers, isEven)) // Output: true

	fmt.Println("Testing AnyOf isEven on oddNumbers: ")
	fmt.Println(AnyOf(oddNumbers, isEven)) // Output: false


	fruitsWithA := []string{"apple", "banana", "cherry"}
	fruitsWithoutA := []string{"cucumber"}
	hasLetterA := func(s string) bool {
		return strings.Contains(strings.ToLower(s), "a")
	}

	fmt.Println("Testing AnyOf hasLeterA on fruitsWithA: ")
	fmt.Println(AnyOf(fruitsWithA, hasLetterA)) // Output: true

	fmt.Println("Testing AnyOf hasLeterA on fruitsWithoutA: ")
	fmt.Println(AnyOf(fruitsWithoutA, hasLetterA)) // Output: false


	fmt.Println("Testing FindIt isEven on evenNumbers: ")
	result, index := FindIf(evenNumbers, isEven)
	fmt.Println("Result: ", result, "Index: ", index) // Output: true, 0

	fmt.Println("Testing FindIf isEven on oddNumbers: ")
	result, index = FindIf(oddNumbers, isEven)
	fmt.Println("Result: ", result, "Index: ", index) // Output: false, -1

	isIntAdjacent := func(n int, m int) bool { return n+1 == m || n-1 == m }

	fmt.Println("Testing AdjacentFind isIntAdjacent on consecutiveNumbers: ")
	fmt.Println(AdjacentFind(consecutiveNumbers, isIntAdjacent)) // Output: 0

	fmt.Println("Testing AdjacentFind isIntAdjacent on nonConsecutiveNumbers: ")
	fmt.Println(AdjacentFind(nonConsecutiveNumbers, isIntAdjacent)) // Output: -1

	isCharAdjacent := func(a rune, b rune) bool { return a+1 == b || a-1 == b }

	fmt.Println("Testing AdjacentFind isCharAdjacent on withConsecutiveLetters: ")
	fmt.Println(AdjacentFind(withConsecutiveLetters, isCharAdjacent)) // Output: 0

	fmt.Println("Testing AdjacentFind isCharAdjacent on withoutConsecutiveLetters: ")
	fmt.Println(AdjacentFind(withoutConsecutiveLetters, isCharAdjacent)) // Output: -1


	fmt.Println("Testing Equal on consecutiveNumbers and nonConsecutiveNumbers: ")
	fmt.Println(Equal(consecutiveNumbers, nonConsecutiveNumbers)) // Output: false

	fmt.Println("Testing Equal on consecutiveNumbers and consecutiveNumbers: ")
	fmt.Println(Equal(consecutiveNumbers, consecutiveNumbers)) // Output: true

	fmt.Println("Testing ReplaceIf on evenNumbers and 48 and isEven: ")
	fmt.Println("evenNumbers: ", evenNumbers)
	result, index = ReplaceIf(evenNumbers, 48, isEven)
	fmt.Println("Result: ", result, "Index: ", index) // Output: 0, true
	fmt.Println("evenNumbers: ", evenNumbers)

	fmt.Println("Testing ReplaceIf on oddNumbers, 5 and isEven: ")
	fmt.Println("oddNumbers: ", oddNumbers)
	result, index = ReplaceIf(oddNumbers, 5, isEven)
	fmt.Println("Result: ", result, "Index: ", index) // Output: -1, false
	fmt.Println("oddNumbers: ", oddNumbers)

	fmt.Println("Testing RemoveIf on evenNumbers on isEven: ")
	fmt.Println("evenNumbers: ", evenNumbers)
	result, index = RemoveIf(&evenNumbers, isEven)
	fmt.Println("Result: ", result, "Index: ", index) // Output: 1, true
	fmt.Println("evenNumbers: ", evenNumbers)

	fmt.Println("Testing isSorted on evenNumbers: ")
	fmt.Println(IsSorted(evenNumbers)) // Output: true

	fmt.Println("Testing isSorted on nonConsecutiveNumbers: ")
	fmt.Println(IsSorted(nonConsecutiveNumbers)) // Output: false

	fmt.Println("Testing Merge with evenNumbers and nonConsecutiveNumbers: ")
	fmt.Println("evenNumbers: ", evenNumbers)
	fmt.Println("nonConsecutiveNumbers: ", nonConsecutiveNumbers)
	fmt.Println(Merge(evenNumbers, nonConsecutiveNumbers)) // Output: true

}