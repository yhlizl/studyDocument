package assignment

import "fmt"

// Q1: Given an array with N+1 integers, with each int in the range from 1 to N, there is at least a number with duplicates, please find the duplicate, and try to optimize time/space complexity if possible.
// Please note, the number of duplicates might be greater than 2.
// Example:
// Input:
// [1,3,3,3,2]
// Output: 3

// /**
//  * @param {number[]} nums
//  * @return {number}
//  */
// const findDuplicated = function(nums) {
//   // your code here.
// };
//use binary search, 一直切半找
func q1FindDuplicated(n []int) int {
	if len(n) == 0 {
		return 0
	}
	if len(n) == 1 {
		return n[0]
	}
	res := 0
	pre := n[0]
	flag := true
	sum := 0
	for _, v := range n {
		flag = flag && (pre == v)
		pre = v
		sum += v
	}
	if flag == true {
		return pre
	}
	fmt.Println(n)
	//binary search
	l := len(n)
	mid := float64(sum) / float64(l)
	lcount := 0
	llist := []int{}
	rcount := 0
	rlist := []int{}
	for _, v := range n {
		if float64(v) >= mid {
			rcount++
			rlist = append(rlist, v)
		} else {
			lcount++
			llist = append(llist, v)
		}
	}
	if lcount > rcount {
		res = q1FindDuplicated(llist)
	} else {
		res = q1FindDuplicated(rlist)
	}
	return res

}

// 
// Q2: An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).
// Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.
// To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all the pixels with the newColor.
// At the end, return the modified image.
// Example 1:
// Input:
// image = [[1,1,1], [1,1,0], [1,0,1]]
// sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2], [2,2,0], [2,0,1]]
// Explanation:
// From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
// by a path of the same color as the starting pixel are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected.
// to the starting pixel.

// /**
//  * @param {number[][]} image
//  * @param {number} sr
//  * @param {number} sc
//  * @param {number} newColor
//  * @return {number[][]}
//  */
// const floodFill = function(image, sr, sc, newColor) {
//     // your code here.
// };
//bfs method to realize!
type vertex struct {
    i,j int
}


func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    // Init
    i,j := sr, sc
    color := image[i][j]
    hash := make(map[vertex]bool)
    q := make([]vertex, 0)
    q = append(q, vertex{i,j})
    // BFS 
    for len(q) > 0 {
        v := q[0]
        q = q[1:]
        hash[v] = true
        image[v.i][v.j] = newColor
        for _, n := range getNeighbors(image, hash, color, v.i, v.j) {
            q = append(q, n)
        }
    }
    return image
}

func getNeighbors(image [][]int, hash map[vertex]bool, color, i, j int) []vertex {
    offset := [][]int {
        {0,1},
        {1,0},
        {0,-1},
        {-1,0},
    }
    var res []vertex
    
    for _, v := range offset {
        _i := i + v[0]
        _j := j + v[1]
        if _i < 0 || _i >= len(image) {
            continue
        } else if _j < 0 || _j >= len(image[_i]) {
            continue
        } else if hash[vertex{_i,_j}] == true {
            continue
        } else if image[_i][_j] == color {
            res = append(res, vertex{_i,_j})
        }
    }
    fmt.Println(i, j, res)
    return res
}


// Q3.1: Please implement below multiply function to make it work.
// multiply(2)(5); // 10

// function multiply(num) {
//   // your code here
// } j
//javascript
function multiply(x) {
    return function(y) {
        if (typeof y !== 'undefined') {
            x = x * y;
            return arguments.callee;
        } else {
            return x;
        }
    };
}