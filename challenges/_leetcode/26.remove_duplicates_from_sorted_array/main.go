package main

func main() {
	r := removeDuplicates([]int{
		0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
		// 1, 1, 2,
	})

	println("result: ", r)
}

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
