package misc

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * Takes an array of (not necessarily sorted) integers and
 * a target value as arguments
 * returns indicies of 2 numbers that add up to the target
 * (or empty array if no 2 numbers like this exist)
 *
 * LeetCode link: https://leetcode.com/problems/two-sum/description/
 */
func TwoSum(nums []int, target int) []int {
	// Holds already checked numbers
	// key is the number, value is the index in original array
	complements := map[int]int{}
	for i, val := range nums {
		// Find the complement value (we're looking for)
		comp := target - val
		// Look for the complement value in complements array (value is index in the original array)
		if compIndex, ok := complements[comp]; ok {
			// Return the indicies of current number and complement's index
			return []int{compIndex, i}
		} else if _, ok := complements[val]; !ok {
			complements[val] = i
		}
	}

	return []int{}
}

/**
 * Takes a 32int number as argument, returns a reverse integer
 * OR 0 if reverse integer overflows
 *
 * Examples:
 * x = 123; return 321
 * x = -123; return -321
 *
 * LeetCode: https://leetcode.com/problems/reverse-integer/description/
 */
func ReverseInteger(x int) int {
	// Check if number is neg (has "-" as first char)
	isNeg := x < 0
	if isNeg {
		// Remove the "-", will reattach it later on
		x = x * -1
	}
	// Convert to string for easier swapping by char
	str := strconv.Itoa(x)
	bytes := []byte(str)
	// Save len of intstring
	strlen := len(bytes)
	// Go through first half and swap with last half (0 with last, 1 with (last - 1)...)
	for i := 0; i < strlen/2; i++ {
		lastI := strlen - 1 - i
		bytes[i], bytes[lastI] = bytes[lastI], bytes[i]
	}
	// Convert string to int
	x, err := strconv.Atoi(string(bytes))
	// If err != nil must be integer overflow
	if err != nil {
		return 0
	}
	// Negate it if needed
	if isNeg {
		x = x * -1
		if x < math.MinInt32 {
			// Check for int overflow
			return 0
		}
	} else if x > math.MaxInt32 {
		// Check for int overflow
		return 0
	}
	return x
}

func IsPalaindrome(x int) bool {
	// Negatives can't be palindrome
	if x < 0 {
		return false
	}
	newNum := 0
	// Iterate through x's digits last to first
	// using %10 (get last digit) and doing i = i / 10 on each step
	for i := x; i > 0; i /= 10 {
		num := i % 10
		// Increase `newNum` times 10, add the new num
		// which will make a reverse digited number
		newNum = (newNum * 10) + num
	}
	return newNum == x
}

// 899 - DCCCIC

// LeetCode: https://leetcode.com/problems/roman-to-integer/description/
func RomanToString(roman string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	bytes := []byte(roman)
	blen := len(bytes)
	sum := 0
	for i := range bytes {
		curr := romans[bytes[i]]
		// Check if next integer is bigger than current
		// and sub by current nuber if that is the case
		// else add
		if (i+1) < blen && curr < romans[bytes[i+1]] {
			fmt.Println("-", curr)
			sum -= curr
		} else {
			fmt.Println("+", curr)
			sum += curr
		}
		fmt.Println(sum)
	}
	return sum
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := []byte{}
	for i := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return string(prefix)
			}
		}
		prefix = append(prefix, strs[0][i])
	}
	return string(prefix)
}

func LengthOfLongestSubstring(s string) int {
	usedChars := map[byte]bool{}
	max := 0
	curr := 0
	for i := range s {
		if usedChars[s[i]] == false {
			usedChars[s[i]] = true
		} else {
			j := i - curr
			for ; s[j] != s[i]; j++ {
				curr--
				usedChars[s[j]] = false
			}
			curr--
			usedChars[s[j]] = false
		}
		usedChars[s[i]] = true
		curr++
		if curr > max {
			max = curr
		}
	}
	return max
}

// LeetCode: https://leetcode.com/problems/implement-strstr/description/
// my: https://leetcode.com/submissions/detail/113899889/
func StrStr(haystack string, needle string) int {
	nlen := len(needle)
	if nlen == 0 {
		return 0
	}
	for i := range haystack {
		if isSubStr(haystack[i:], needle, &nlen) {
			return i
		}
	}
	return -1
}

func isSubStr(haystack string, needle string, nlen *int) bool {
	if len(haystack) < *nlen {
		return false
	}
	for i := range needle {
		if haystack[i] != needle[i] {
			return false
		}
	}
	return true
}

// LeetCode: https://leetcode.com/problems/search-insert-position/description/
// my: https://leetcode.com/submissions/detail/114033378/
func SearchInsert(nums []int, target int) int {
	return subSearch(nums, &target, 0)
}

func subSearch(nums []int, target *int, start int) int {
	nlen := len(nums)
	if nlen == 1 {
		if nums[0] == *target {
			return start
		} else if nums[0] < *target {
			return start + 1
		}
		if start > 0 {
			return start - 1
		}
		return 0
	} else if nlen == 0 {
		return 0
	}
	pivot := nlen / 2
	if nums[pivot] == *target {
		return start + pivot
	} else if nums[pivot] < *target {
		return subSearch(nums[pivot:], target, start+pivot)
	}
	return subSearch(nums[:pivot], target, start)
}

// LeetCode: https://leetcode.com/problems/plus-one/description/
// my: https://leetcode.com/submissions/detail/114039928/
func PlusOne(digits []int) []int {
	dlen := len(digits)
	if digits[dlen-1] < 9 {
		digits[dlen-1]++
	} else {
		i := dlen - 1
		for ; i > 0 && digits[i] == 9; i-- {
			digits[i] = 0
		}
		if i == 0 && digits[i] == 9 {
			digits[i] = 1
			digits = append(digits, 0)
		} else {
			digits[i]++
		}
	}
	return digits
}

// LeetCode: https://leetcode.com/problems/integer-to-roman/description/
// my: https://leetcode.com/submissions/detail/114181883/
func IntToRoman(num int) string {
	dict := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M"}

	str := ""
	start := 1
	for i := num; i != 0; i /= 10 {
		digit := i % 10
		if digit == 4 {
			str = dict[1*start] + dict[5*start] + str
		} else if digit == 9 {
			str = dict[1*start] + dict[10*start] + str
		} else if digit != 0 {
			tmp := ""
			if digit > 5 {
				tmp = dict[5*start]
				digit -= 5
			} else if digit == 5 {
				tmp = dict[5*start]
				digit -= 5
			}
			for j := digit; j != 0; j-- {
				tmp += dict[1*start]
			}
			str = tmp + str
		}
		start *= 10
	}
	return str
}

// Valid parentheses
// LeetCode: https://leetcode.com/problems/valid-parentheses/description/
// my: https://leetcode.com/submissions/detail/114185482/
func IsValid(s string) bool {
	dict := map[byte]byte{
		')': '(',
		'}': '{',
		']': '['}
	exp := []byte{}
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			exp = append(exp, s[i])
		} else {
			lasti := len(exp) - 1
			if lasti < 0 || exp[lasti] != dict[s[i]] {
				return false
			}
			exp = exp[:lasti]
			i--
		}
	}
	if len(exp) != 0 {
		return false
	}
	return true
}

// LeetCode: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
// my: https://leetcode.com/submissions/detail/114906036/
func RemoveDuplicates(nums []int) int {
	remlen := len(nums)
	for i := 0; i < (remlen - 1); i++ {
		if nums[i] == nums[i+1] {
			sameCount := 1
			for j := i + 1; j < remlen-1; j++ {
				if nums[j] != nums[j+1] {
					break
				}
				sameCount++
			}
			nums = append(nums[:i], nums[i+sameCount:]...)
			i--
			remlen -= sameCount
		}
	}
	return remlen
}

// To get around go nazism
func FmtTest() {
	fmt.Println("Works")
}

// d1 / d2 = x
// d1 = d2 * x

func Divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	r := 0
	multi := 1
	d1neg := dividend < 0
	d2neg := divisor < 0

	if d1neg != d2neg {
		multi = -1
	}
	if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	less := (!d1neg && !d2neg) || (!d1neg && d2neg)
	more := (d1neg && d2neg) || (d1neg && !d2neg)
	for tmp := 0; (less && tmp <= dividend) || (more && tmp >= dividend); r += multi {
		if multi == -1 {
			tmp -= divisor
		} else {
			tmp += divisor
		}
	}
	if r != 0 {
		r -= multi
	}
	return r
}
