package countTriplets

import "fmt"

func Solution() {
	test1()
	test2()
	test3()
}

func test1() {
	arr := []int64{1, 2, 2, 4}
	r := int64(2)

	fmt.Println("countTriplets result is: ", countTriplets(arr, r))
}

func test2() {
	arr := []int64{1, 3, 9, 9, 27, 81}
	r := int64(3)

	fmt.Println("countTriplets result is: ", countTriplets(arr, r))
}

func test3() {
	arr := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	r := int64(1)

	fmt.Println("countTriplets result is: ", countTriplets(arr, r))
}

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	v2 := make(map[int64]int64)
	v3 := make(map[int64]int64)

	var count int64 = 0
	for _, ele := range arr {
		count += v3[ele]
		v3[ele*r] += v2[ele]
		v2[ele*r] += 1
	}

	return count
}

// Complete the countTriplets function below.
func countTripletsFailed(arr []int64, r int64) int64 {
	counts := make(map[int64]int64)

	for _, ele := range arr {
		counts[ele]++
	}

	var countOfTriplets int64 = 0
	for _, ele := range arr {
		count1 := counts[ele*r]
		if count1 > 0 {
			count2 := counts[ele*r*r]
			if count2 > 0 {
				countOfTriplets += count1 * count2
			}
		}
	}

	return countOfTriplets
}
