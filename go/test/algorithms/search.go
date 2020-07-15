package algo

// Search function return target position or -1 means target not exist
func Search(arr []int, target int) int {
	for i, x := range arr {
		if x == target {
			return i
		}
	}
	return -1
}

// BinarySearch only used to search target in an ordered set from small to large
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	//采用左右都是闭区间的做法
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// BinarySearch1 是左闭右开的写法
func BinarySearch1(arr []int, target int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
