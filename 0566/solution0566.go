package solution0566

// 566. Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/
/*
In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.
You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.
The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.
If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
Output:
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.

Example 2:
Input:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
Output:
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.

Note:
The height and width of the given matrix is in range [1, 100].
The given r and c are all positive.
 */
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nums == nil {
		return nil
	}
	m := len(nums)
	if m == 0 {
		return nums
	}
	n := len(nums[0])
	if n == 0 {
		return nums
	}
	if m*n != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[count/c][count%c] = nums[i][j]
			count++
		}
	}
	return res
}

// Runtime: 92 ms, faster than 91.11% of Go online submissions for Reshape the Matrix.
