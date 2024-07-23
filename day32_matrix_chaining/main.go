package main

import (
	"fmt"
	"log"
)

type MatrixDimension struct {
	row    int
	column int
}

type Range struct {
	i int
	j int
}

// var r = [20][20]int{}
var location = map[Range]int{}

func Cal(i int, j int, matrices []MatrixDimension, location map[Range]int) int {
	log.Printf("Called %v, %v \n", i, j)
	r := Range{i, j}
	// if r[i][j] >= 0 {
	// 	return r[i][j]
	// }
	// if _, ok := r[Range{i, j}]; ok {
	// 	return
	// }
	if i >= j {
		return 0
		// r[Range{i, j}] = true
		// r[i][j] = true
		// return
	}
	if j == i+1 {
		return matrices[i].row * matrices[j].row * matrices[j].column
		// r[Range{i, j}] = true
		// return
	}
	var res = 1<<32 - 1

	for k := i; k < j; k++ {
		// res = min()
		x1 := Cal(i, k, matrices, location)
		x2 := Cal(k+1, j, matrices, location)
		x3 := matrices[i].row * matrices[j].column * matrices[k].column
		res = min(res, x1+x2+x3)
		fmt.Printf("%v - %v - %v at %v in %v - %v \n", x1, x2, x3, k, i, j)
		location[r] = k
	}
	fmt.Printf("%v:%v is %v at %v \n", i, j, location[r], res)
	return res
	// r[Range{i, j}] = true
	// return 0
}
func CalDP(matrices []MatrixDimension) int {
	var res = [20][20]int{}
	for i := 0; i < 20; i++ {
		for j := i + 1; j < 20; j++ {
			res[i][j] = 1<<32 - 1
		}
	}
	if len(matrices) <= 1 {
		return 0
	}
	for i := 0; i <= len(matrices); i++ {
		for j := 0; j < i; j++ {
			res[i][j] := min(res[i][j], )
		}
	}

}
func main() {
	matrices := []MatrixDimension{
		{50, 40},
		{40, 30},
		{30, 70},
		{70, 100},
		{100, 10},
	}
	// res := [20][20]int{}
	// fmt.Println(res[0][0])
	fmt.Println(Cal(0, 4, matrices, location))
	fmt.Println(location)

}
