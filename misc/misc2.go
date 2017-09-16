package misc

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
	"unicode"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestList() ListNode {
	return ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 10}}}}
}
func TestList2() ListNode {
	return ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 9}}}
}

// LeetCode https://leetcode.com/problems/merge-two-sorted-lists/description/
// my: https://leetcode.com/submissions/detail/114635943/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}

	recMerge(l1, l2)

	return l1
}

func recMerge(lower *ListNode, higher *ListNode) *ListNode {
	if lower == nil || higher == nil {
		return nil
	}
	if lower.Next == nil {
		lower.Next = higher
		return nil
	} else if lower.Next.Val > higher.Val {
		tmp := higher.Next
		lower.Next, higher.Next = higher, lower.Next
		return recMerge(lower.Next, tmp)
	}
	return recMerge(lower.Next, higher)
}

func T(l *ListNode) {
	for ; l != nil; l = l.Next {
		fmt.Println(l.Val)
	}
}

// LeetCode: https://leetcode.com/problems/string-to-integer-atoi/description/
// my: https://leetcode.com/submissions/detail/114740090/
func MyAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	isNeg := false
	wscheck := true
	i := 0
	r := 0
	dict := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9}
	for ; i < len(str); i++ {
		if wscheck {
			if unicode.IsSpace(rune(str[i])) {
				continue
			}
			if str[i] == '-' {
				isNeg = true
				wscheck = false
				continue
			}
			if str[i] == '+' {
				wscheck = false
				continue
			}
		}
		digit, ok := dict[str[i]]
		if !ok {
			break
		}
		wscheck = false
		r *= 10
		r += digit
		if r > math.MaxInt32 {
			break
		}
	}

	if isNeg {
		r = -r
	}
	if r < math.MinInt32 {
		return math.MinInt32
	}
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r
}

func MajorityElement(nums []int) int {
	seen := map[int]int{}
	nlen := len(nums)
	overhalf := (nlen / 2) + 1
	maxi := 0
	maxn := 0
	for i := 0; i < nlen; i++ {
		n := seen[nums[i]]
		n++
		if n > maxi {
			maxn = nums[i]
			maxi = n
		}
		if n > overhalf {
			return maxn
		}
		seen[nums[i]] = n
	}
	return maxn
}

// https://leetcode.com/problems/length-of-last-word/description/
// my: https://leetcode.com/submissions/detail/114918067/
func LengthOfLastWord(s string) int {
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		if r != 0 && unicode.IsSpace(rune(s[i])) {
			return r
		}
		if !unicode.IsSpace(rune(s[i])) {
			r++
		}
	}
	return r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestTreeNode() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4},
			Right: &TreeNode{
				Val: 12}},
		Right: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val: -1}}}
}

// LeetCode: https://leetcode.com/problems/path-sum/description/
// my: https://leetcode.com/submissions/detail/114920149/
func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return HasPathSum(root.Left, sum) || HasPathSum(root.Right, sum)
}

// LeetCode: https://leetcode.com/problems/climbing-stairs/description/
// my: https://leetcode.com/submissions/detail/114922894/
// fibonacci
func ClimbStairs(n int) int {
	if n < 2 {
		return 1
	}
	l1 := 1
	l2 := 1
	for i := 2; i <= n; i++ {
		l1, l2 = l2, l1+l2
	}
	return l2
}

// https://leetcode.com/problems/shuffle-an-array/description/
//
type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{original: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nlen := len(this.original)
	shuffled := make([]int, nlen)
	copy(shuffled, this.original)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := nlen - 1; i > 0; i-- {
		swapi := rnd.Intn(i)
		shuffled[i], shuffled[swapi] = shuffled[swapi], shuffled[i]
	}
	return shuffled
}

// LeetCode: https://leetcode.com/problems/fizz-buzz/description/
// my: https://leetcode.com/submissions/detail/115534964/
func FizzBuzz(n int) []string {
	if n < 0 {
		return []string{}
	}
	strings := make([]string, n)
	for i := 1; i <= n; i++ {
		three := i % 3
		five := i % 5
		if three == 0 {
			if five == 0 {
				strings[i-1] = "FizzBuzz"
			} else {
				strings[i-1] = "Fizz"
			}
		} else if five == 0 {
			strings[i-1] = "Buzz"
		} else {
			strings[i-1] = strconv.Itoa(i)
		}
	}
	return strings
}

// https://leetcode.com/problems/min-stack/description/
// my: https://leetcode.com/submissions/detail/118220020/
type MinStack struct {
	nums []int
	min  []int
}

/** initialize your data structure here. */
func ConstructorMS() MinStack {
	return MinStack{nums: []int{}, min: []int{}}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	minlen := len(this.min)
	if minlen < 1 || x <= this.min[minlen-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	nlen := len(this.nums)
	if nlen == 0 {
		return
	}
	mlen := len(this.min)
	if this.min[mlen-1] == this.nums[nlen-1] {
		this.min = this.min[:mlen-1]
	}
	this.nums = this.nums[:nlen-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	mlen := len(this.min)
	if mlen == 0 {
		return 0
	}
	return this.min[mlen-1]
}

func (ms *MinStack) Print() {
	fmt.Println(ms.nums)
	fmt.Println(ms.min)
}

// https://leetcode.com/problems/find-peak-element/description/
// my: https://leetcode.com/submissions/detail/118398015/
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)
	for i := (right + left) / 2; right-left != 0; i = (right + left) / 2 {
		if i-1 >= left && nums[i-1] >= nums[i] {
			right = i
		} else if i+1 < right && nums[i+1] >= nums[i] {
			left = i
		} else {
			return i
		}
	}
	return right + left/2
}

// my: https://leetcode.com/submissions/detail/118412320/
func LargestNumber(nums []int) string {
	r := ln(nums)
	if len(r) > 0 && r[0] == '0' {
		return "0"
	}
	return r
}

func ln(nums []int) string {
	left, right := 0, len(nums)-1
	if right < 0 {
		return ""
	}
	if right == 0 {
		return strconv.Itoa(nums[right])
	}
	pivoti := right / 2
	pivot := nums[pivoti]
	var pow, npow int
	if pivot == 0 {
		pow = 10
	} else {
		pow = int(math.Pow10(int(math.Log10(float64(pivot))) + 1))
	}
	nums[pivoti], nums[right] = nums[right], nums[pivoti]
	for i := 0; i < right; i++ {
		if nums[i] == 0 {
			npow = 10
		} else {
			npow = int(math.Pow10(int(math.Log10(float64(nums[i]))) + 1))
		}
		if nums[i]*pow+pivot >= pivot*npow+nums[i] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	return ln(nums[:left]) + strconv.Itoa(nums[left]) + ln(nums[left+1:])
}
func LargestNumber2(nums []int) string {
	left, right := 0, len(nums)-1
	if right < 0 {
		return ""
	}
	if right == 0 {
		return strconv.Itoa(nums[right])
	}
	pivoti := right / 2
	nums[pivoti], nums[right] = nums[right], nums[pivoti]
	pivotstr := strconv.Itoa(nums[right])
	var first, second int
	for i := 0; i < right; i++ {
		first, _ = strconv.Atoi(strconv.Itoa(nums[i]) + pivotstr)
		second, _ = strconv.Atoi(pivotstr + strconv.Itoa(nums[i]))
		if first >= second {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	return LargestNumber(nums[:left]) + strconv.Itoa(nums[left]) + LargestNumber(nums[left+1:])
}

func fmtholder() {
	fmt.Println("test")
}

func Qsort(arr []int) []int {
	alen := len(arr)

	if alen <= 1 {
		return arr
	}
	left, right := 0, alen-1
	pivot := arr[right]
	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	Qsort(arr[:left])
	Qsort(arr[left+1:])

	return arr
}
