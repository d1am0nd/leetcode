package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "time"
	_ "math"

	"algorithms/misc"
	"algorithms/sort"
)

func main() {
	// qsort()
	// twoSum()
	// reverseInt()
	// isPalindrome()
	// romanToString()
	// longestCommonPrefix()
	// longestCommonPrefix()
	// lengthOfLongestSubstring()
	// strStr()
	// searchInsert()
	// plusOne()
	// intToRoman()
	// isValid()
	// divide()

	// t := misc.TestList()
	// t2 := misc.TestList2()
	// misc.MergeTwoLists(&t, nil)

	// myAtoi()
	// majorityElement()
	// largestNumber()
	// removeDuplicates()
	// hasPathSum()
	// climbStairs()
	// shuffle()

	// fizzBuzz()
	// minStack()
	// findPeakElement()

	a := []int{1,
		32,
		43,
		34,
		3}

	fmt.Println("Sup boy", a)

	/*
		then := time.Now()
		for i := 0; i < 10000000; i++ {
			largestNumber()
		}
		fmt.Println(time.Since(then))
		then2 := time.Now()
		for i := 0; i < 10000000; i++ {
			largestNumber2()
		}
		fmt.Println(time.Since(then2))
	*/
	/*
		(largestNumber())
	*/

	// fmt.Println(4/(-2))
}

// 4334
func largestNumber() {
	(misc.LargestNumber([]int{1, 32, 43, 34, 3, 100, 1003210032, 43, 999999999999, 3, 10032, 43, 100321003210032, 3, 100}))
	(misc.LargestNumber([]int{0, 0, 1, 0, 0}))
	(misc.LargestNumber([]int{0, 0, 0}))
	(misc.LargestNumber([]int{0, 0}))
}
func largestNumber2() {
	(misc.LargestNumber2([]int{1, 32, 43, 34, 3, 100, 1003210032, 43, 999999999999, 3, 10032, 43, 100321003210032, 3, 100}))
	(misc.LargestNumber2([]int{0, 0, 1, 0, 0}))
	(misc.LargestNumber2([]int{0, 0, 0}))
	(misc.LargestNumber2([]int{0, 0}))
}

func findPeakElement() {
	fmt.Println(misc.FindPeakElement([]int{1, 2, 3, 4, 5, 6, 2, 1, 2, 3, 4, 5, 3, 2, 1}))
	fmt.Println(misc.FindPeakElement([]int{}))
	fmt.Println(misc.FindPeakElement([]int{3}))
	fmt.Println(misc.FindPeakElement([]int{1, 2, 3}))
	fmt.Println(misc.FindPeakElement([]int{3, 2, 4, 2, 3}))
	fmt.Println(misc.FindPeakElement([]int{3, 2, 2, 2, 3}))
	fmt.Println(misc.FindPeakElement([]int{3, 2, 1}))
	fmt.Println(misc.FindPeakElement([]int{2, 2, 2, 2}))
	fmt.Println(misc.FindPeakElement([]int{3, 2, 3, 2, 3}))
}

func minStack() {
	minStack := misc.ConstructorMS()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.Print()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Print()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Print()
}

// 1 2 3 123 122 111 12
// .1 .2 .3 .123 .122 .111 0.12
// 3 2 1 123 122 12 111
// 3 2 123 122 12 111 1
func fizzBuzz() {
	fmt.Println(misc.FizzBuzz(-1))
}

func shuffle() {
	s := misc.Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
}

func climbStairs() {
	fmt.Println(misc.ClimbStairs(11))
}

func hasPathSum() {
	fmt.Println(misc.HasPathSum(misc.TestTreeNode(), 8))
}
func lengthOfLastWord() {
	fmt.Println(misc.LengthOfLastWord("test 123123 123"))
}
func removeDuplicates() {
	fmt.Println(misc.RemoveDuplicates([]int{1, 1, 2, 3, 4, 4, 5, 5, 5}))
	fmt.Println(misc.RemoveDuplicates([]int{}))
	fmt.Println(misc.RemoveDuplicates([]int{-12, -11, -9, -9, 0, 12, 33, 33, 33}))
}

func majorityElement() {
	fmt.Println(misc.MajorityElement([]int{1, 2, 2, 2, 2, 3, 3, 6, 4}))
	fmt.Println(misc.MajorityElement([]int{3, 4, 5, 6, 6, 6}))
	fmt.Println(misc.MajorityElement([]int{1, 2, 2, 2, 2, 3, 3, 6, 4}))
}

func myAtoi() {
	fmt.Println(misc.MyAtoi("-9223372036854775809"))
}

func divide() {
	fmt.Println(misc.Divide(0, -2))
}
func isValid() {
	fmt.Println(misc.IsValid("[])"))
}

func intToRoman() {
	fmt.Println(misc.IntToRoman(105))
}

func plusOne() {
	fmt.Println(misc.PlusOne([]int{}))
}

func searchInsert() {
	fmt.Println(misc.SearchInsert([]int{1, 3, 5, 9}, 0))
}

func strStr() {
	fmt.Println(misc.StrStr("123123", "312"))
}

func lengthOfLongestSubstring() {
	fmt.Println(misc.LengthOfLongestSubstring("abcabcbb"))
}

func longestCommonPrefix() {
	fmt.Println(misc.LongestCommonPrefix([]string{"123", "123as dada", "123123123sa", "1231233123"}))
}

func romanToString() {
	fmt.Println(misc.RomanToString("LXXXIX"))
}

func isPalaindrome() {
	fmt.Println(misc.IsPalaindrome(1223221))
}

func reverseInt() {
	fmt.Println(misc.ReverseInteger(7463847412))
}

func twoSum() {
	fmt.Println(misc.TwoSum([]int{1, 1, 3, 5, 10}, 13))
}

func qsort() {
	// Get numbers from file
	var numbers []int

	file, err := os.Open("./numbers.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&numbers)
	if err != nil {
		panic(err)
	}
	// numbers var now holds all (5000) numbers from file

	fmt.Println(len(numbers))

	// Quicksort
	algorithms.Qsort(numbers)
	for _, num := range numbers {
		fmt.Println(num)
	}

}
