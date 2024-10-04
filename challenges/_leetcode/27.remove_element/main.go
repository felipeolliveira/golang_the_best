package main

func main() {
	r := removeElement([]int{
		0, 1, 2, 2, 3, 0, 4, 2,
		// 3, 2, 2, 3, 3,
	}, 2)

	println(r)
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}
