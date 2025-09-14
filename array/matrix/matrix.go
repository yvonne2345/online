package main

// 螺旋矩阵
//func generateMatrix(n int) [][]int {
//	startX, startY := 0, 0
//	offset, num := 0, 0
//	l := n
//	arr := make([][]int, n)
//	for i := range arr {
//		arr[i] = make([]int, n)
//	}
//	for n/2 > 0 {
//		x, y := startX, startY
//		for y = startY; y < l-offset-1; y++ {
//			num++
//			arr[x][y] = num
//		}
//		for x = startX; x < l-offset-1; x++ {
//			num++
//			arr[x][y] = num
//		}
//		for y = l - offset - 1; y > startY; y-- {
//			num++
//			arr[x][y] = num
//		}
//		for x = l - offset - 1; x > startX; x-- {
//			num++
//			arr[x][y] = num
//		}
//		startX++
//		startY++
//		offset++
//		n--
//	}
//	if l%2 == 1 {
//		arr[l/2][l/2] = l * l
//	}
//	return arr
//}

//func generateMatrix(n int) [][]int {
//	top, right, bottom, left := 0, n-1, n-1, 0
//	arr := make([][]int, n)
//	for i := range arr {
//		arr[i] = make([]int, n)
//	}
//	var i int
//	count := 1
//	for i < n*n {
//		for j := left; j <= right && i < n*n; j++ {
//			arr[top][j] = count
//			count++
//			i++
//		}
//		top++
//		for j := top; j <= bottom && i < n*n; j++ {
//			arr[j][right] = count
//			count++
//			i++
//		}
//		right--
//		for j := right; j >= left && i < n*n; j-- {
//			arr[bottom][j] = count
//			count++
//			i++
//		}
//		bottom--
//		for j := bottom; j >= top && i < n*n; j-- {
//			arr[j][left] = count
//			count++
//			i++
//		}
//		left++
//	}
//	return arr
//}

func generateMatrix(n int) [][]int {
	top, bottom, left, right := 0, n-1, 0, n-1
	count := 0
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	for count < n*n {
		for i := left; i <= right && count < n*n; i++ {
			count++
			a[top][i] = count
		}
		top++
		for i := top; i <= bottom && count < n*n; i++ {
			count++
			a[i][right] = count
		}
		right--
		for i := right; i >= left && count < n*n; i-- {
			count++
			a[bottom][i] = count
		}
		bottom--
		for i := bottom; i >= top && count < n*n; i-- {
			count++
			a[i][left] = count
		}
		left++
	}
	return a
}

//func spiralOrder(matrix [][]int) []int {
//	a := len(matrix)
//	b := len(matrix[0])
//	top, left, bottom, right := 0, 0, a-1, b-1
//	var i int
//	var arr []int
//	for i < a*b {
//		for j := left; j <= right && i < a*b; j++ {
//			arr = append(arr, matrix[top][j])
//			i++
//		}
//		top++
//		for j := top; j <= bottom && i < a*b; j++ {
//			arr = append(arr, matrix[j][right])
//			i++
//		}
//		right--
//		for j := right; j >= left && i < a*b; j-- {
//			arr = append(arr, matrix[bottom][j])
//			i++
//		}
//		bottom--
//		for j := bottom; j >= top && i < a*b; j-- {
//			arr = append(arr, matrix[j][left])
//			i++
//		}
//		left++
//	}
//	return arr
//}

//func spiralArray(array [][]int) []int {
//	if len(array) == 0 || (len(array) != 0 && len(array[0]) == 0) {
//		return []int{}
//	}
//	a := len(array)
//	b := len(array[0])
//	top, right, bottom, left := 0, b-1, a-1, 0
//	var i int
//	var nums []int
//	for i < a*b {
//		for j := left; j <= right && i < a*b; j++ {
//			nums = append(nums, array[top][j])
//			i++
//		}
//		top++
//		for j := top; j <= bottom && i < a*b; j++ {
//			nums = append(nums, array[j][right])
//			i++
//		}
//		right--
//		for j := right; j >= left && i < a*b; j-- {
//			nums = append(nums, array[bottom][j])
//			i++
//		}
//		bottom--
//		for j := bottom; j >= top && i < a*b; j-- {
//			nums = append(nums, array[j][left])
//			i++
//		}
//		left++
//	}
//	return nums
//}

func main() {
	//n := 3
	//fmt.Println(generateMatrix(n))

	//matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//fmt.Println(spiralOrder(matrix))

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	spiralArray(matrix)
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	top, bottom, left, right := 0, m-1, 0, n-1
	count := 0
	nums := make([]int, 0)
	for count < m*n {
		for i := left; i <= right && count < m*n; i++ {
			nums = append(nums, matrix[top][i])
			count++
		}
		top++
		for i := top; i <= bottom && count < m*n; i++ {
			nums = append(nums, matrix[i][right])
			count++
		}
		right--
		for i := right; i >= left && count < m*n; i-- {
			nums = append(nums, matrix[bottom][i])
			count++
		}
		bottom--
		for i := bottom; i >= top && count < m*n; i-- {
			nums = append(nums, matrix[i][left])
			count++
		}
		left++
	}
	return nums
}

func spiralArray(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}
	m := len(array)
	n := len(array[0])
	top, bottom, left, right := 0, m-1, 0, n-1
	nums := make([]int, 0)
	count := 0
	for count < m*n {
		for i := left; i <= right && count < m*n; i++ {
			nums = append(nums, array[top][i])
			count++
		}
		top++
		for i := top; i <= bottom && count < m*n; i++ {
			nums = append(nums, array[i][right])
			count++
		}
		right--
		for i := right; i >= left && count < m*n; i-- {
			nums = append(nums, array[bottom][i])
			count++
		}
		bottom--
		for i := bottom; i >= top && count < m*n; i-- {
			nums = append(nums, array[i][left])
			count++
		}
		left++
	}
	return nums
}
