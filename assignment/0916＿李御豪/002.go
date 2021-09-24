package main

// ## Two Sum Less Than K

// ---

// Given an array A of integers and integer K. Any interger pair in A could be combined as a set. Find and return all combinations of sets where the sum S of elements is the maximum but less than K ( a, b are in A and a + b = S < K and S is maximum, return [a, b] ). The answer must exclude duplicates, but it's fine in any order. e.g. [1, 2] or [2, 1] may be treated as the same, no need to return twice in one output. Please answer the time complexity of your program, toof.

// Free to use any programing language

// Example 1:

// ```
// Input: A = [1, 2, 3, 4]. K = 4
// Output: [1, 2], because 1+2 = 3 < 4
// ```

// Example 2:

// ```
// Input: A = [1, 2, 3]. K = 3
// Output: []
// ```

// Example 3:

// ```
// Input: A = [1, 2, 2, 3, 4]. K = 5
// Output: [[1, 3], [2, 2]], because both `1+3` and `2+2` are 4 < 5
// ```

func twosum(arr []int, target int) [][]int {
	l := len(arr)
	if l < 2 {
		return [][]int{}
	}
	dp := map[int]bool{}
	for _, v := range arr {
		dp[v] = true
	}
	res := [][]int{}
	for i := 1; i < target; i++ {
		spec := target - i

		for j := 0; j < l; j++ {
			if _, ok := dp[spec-arr[j]]; ok {
				delete(dp, arr[j])
				temp := []int{arr[j], spec - arr[j]}
				res = append(res, temp)
			}
		}
		if len(res) > 0 {
			return res
		}
	}
	return [][]int{}

}
