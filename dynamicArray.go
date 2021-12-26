package main

import (
	"fmt"
)

func mainMock1() {
	var n int32 = 3
	queries := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	res := dynamicArray(n, queries)
	fmt.Println("hello world ", 5%5)
	fmt.Println("hello world ", res)
}

func dynamicArray(n int32, queries [][]int32) []int32 {
	seqList := make([][]int32, n)
	var answersArr []int32
	var lastAnswer int32 = 0
	for _, query := range queries {
		x := query[1]
		y := query[2]
		if query[0] == 1 {
			index, seq := appendIndex(x, y, n, lastAnswer, seqList)
			seqList[index] = seq
		}
		if query[0] == 2 {
			rowIndex := (x ^ lastAnswer) % n
			colIndex := int(y) % len(seqList[rowIndex])
			lastAnswer = seqList[rowIndex][colIndex]
			answersArr = append(answersArr, lastAnswer)
		}
	}

	return answersArr
}

func appendIndex(x int32, y int32, n int32, lastAnswer int32, seqList [][]int32) (int32, []int32) {
	rowIndex := (x ^ lastAnswer) % n
	seq := seqList[rowIndex]
	seq = append(seq, y)
	return rowIndex, seq
}
