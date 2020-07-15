package algo

//InsertionSort will sort a slice by using InsertionSort
func InsertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		x := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}

//InsertionSort1 will sort slice reversely
func InsertionSort1(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		x := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}

//InsertionSort2 will sort slice from end to start
func InsertionSort2(arr []int) {
	len := len(arr)
	for i := len - 2; i >= 0; i-- {
		x := arr[i]
		j := i + 1
		for j <= len-1 && arr[j] < x {
			arr[j-1] = arr[j]
			j++
		}
		arr[j-1] = x
	}
}

//InsertionSort3 will sort slice from end to start reversely
func InsertionSort3(arr []int) {
	len := len(arr)
	for i := len - 2; i >= 0; i-- {
		x := arr[i]
		j := i + 1
		for j <= len-1 && arr[j] > x {
			arr[j-1] = arr[j]
			j++
		}
		arr[j-1] = x
	}
}

//BinaryInsertionSort 插入排序时使用二分法搜索插入位置
//二分搜索的基本条件是确保集合有序
//那么首先就要假定数组的一部分是有序的
//假定当前下标之前有序
func BinaryInsertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		x := arr[i]
		left := 0
		right := i - 1
		for left <= right {
			mid := left + (right-left)/2
			//x位置在mid右边
			//x < arr[mid]，希望正序排列，所以要去左边的区间找x的位置
			if x < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		//j一定大于left
		//正序排列，前部为假定的有序区间
		//候选的位置是left
		//那么left和它之前的元素全部向前移动一位
		for j := i; j > left; j-- {
			arr[j] = arr[j-1]
		}
		arr[left] = x
	}
}

//BinaryInsertionSort1 倒序实现
func BinaryInsertionSort1(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		x := arr[i]
		left := 0
		right := i - 1
		for left <= right {
			mid := left + (right-left)/2
			//x位置在mid右边
			//x > arr[mid]，希望倒序排列，所以要去左边的区间找x的位置
			if x > arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i; j > left; j-- {
			arr[j] = arr[j-1]
		}
		arr[left] = x
	}
}

//BinaryInsertionSort2 反向遍历实现
func BinaryInsertionSort2(arr []int) {
	len := len(arr)
	//假定后面的是有序的，且正序排列
	for i := len - 2; i >= 0; i-- {
		x := arr[i]
		left := i + 1
		right := len - 1
		for left <= right {
			mid := left + (right-left)/2
			if x > arr[mid] {
				//x应该去mid右边找位置
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		//j小于left
		for j := i; j < left-1; j++ {
			//[i, left]都要往前移一位
			arr[j] = arr[j+1]
		}
		arr[left-1] = x
	}
}

//BinaryInsertionSort3 reverse order, loop from end to start
func BinaryInsertionSort3(arr []int) {
	len := len(arr)
	for i := len - 2; i >= 0; i-- {
		x := arr[i]
		left := i + 1
		right := len - 1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] > x {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		for j := i; j < left-1; j++ {
			arr[j] = arr[j+1]
		}
		arr[left-1] = x
	}
}
