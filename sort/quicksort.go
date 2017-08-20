package algorithms

func Qsort(arr []int) []int {
    alen := len(arr)

    // If array is 1 number or less wrong => it's sorted already
    // return as is
    if alen <= 1 {
        return arr
    }

    // left = first index of arr
    // right = last index of arr
    left, right := 0, alen - 1
    // save the pivot for comparison in a var
    pivot := arr[right]

    // iterate through numbers,
    // place those smaller than pivot on
    // index `left`, inc the index afterwards
    for i := range arr {
        if arr[i] < pivot {
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }

    // place the pivot in the correct spot (on the `left` index
    // since we know numbers small than that will have smaller indexes)
    arr[left], arr[right] = arr[right], arr[left]

    // sort left and right part recursively
    Qsort(arr[:left])
    Qsort(arr[left + 1:])

    return arr
}
