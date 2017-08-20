package misc

import (
	"fmt"
	"math"
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
func LargestNumber(nums []int) string {
	return "2"
}
func fmtholder() {
	fmt.Println("test")
}

func Qsort(arr []int) []int {
    alen := len(arr)

    if alen <= 1 {
        return arr
    }
    left, right := 0, alen - 1
    pivot := arr[right]
    for i := range arr {
        if arr[i] < pivot {
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }

    arr[left], arr[right] = arr[right], arr[left]

    Qsort(arr[:left])
    Qsort(arr[left + 1:])

    return arr
}
