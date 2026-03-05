package main

import "fmt"

func main() {
	// var nums [5]int
	// fmt.Println(nums)

	// var nums1 = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(nums1)

	// nums2 := [5]int{8, 9, 0, 3, 5}
	// fmt.Println(nums2)

	// fruits := [...]string{"apple", "cherry", "banana"}
	// fmt.Println(fruits)

	// nums3 := [10]int{1: 10, 4: 30}
	// fmt.Println(nums3)

	// nums4 := [6]int{3, 5, 7, 6, 3, 9}
	// fmt.Println(nums4[1])
	// nums4[1] = 400
	// fmt.Println(nums4)

	// i := 3
	// if i >= 0 && i < len(nums4) {
	// 	fmt.Println(nums4[i])
	// } else {
	// 	fmt.Printf("Value %d exit out of bound [0, %d]\n", i, len(nums4)-1)
	// }

	// for i := 0; i < len(nums4); i++ {
	// 	fmt.Println(i, nums4[i])
	// }

	// for i, v := range nums4 {
	// 	fmt.Println(i, v)
	// }
	// for _, v := range nums4 {
	// 	fmt.Println(v)
	// }
	// for i := range nums4 {
	// 	fmt.Println(i)
	// }

	arr := [10]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	isPalindrome(arr)

}

func isPalindrome(nums [10]int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] != nums[j] {
			fmt.Println("It is not palindrome")
			return
		}
	}
	fmt.Println("It is palindrome")
}
