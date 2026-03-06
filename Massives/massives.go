package main

import (
	"fmt"
	// "math"
	// "strconv"
)

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

	// arr := [10]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	// // isPalindrome(arr)
	// fmt.Println(SumNeighbors(arr))

	arr11 := [5]int{3,8,1,8,1}
	fmt.Println(generateCode(arr11))

}

/*
Секретный код
	Вы работаете в секретной лаборатории, где разрабатываются новые методы шифрования. 
	Ваша задача — создать функцию, которая будет обрабатывать
	массив целых чисел и генерировать "секретный код" на основе определенных правил.
Входные данные
	Функция принимает массив целых чисел фиксированного размера — 5 значений, 
	и возвращает строку, представляющую сгенерированный код.
Генерация секретного кода
	В начале кода должно находиться минимальное значение массива.
	В конце кода должно находиться максимальное значение массива.
Между максимальным и минимальными значениями, должны быть расположены все значения, 
перед которыми должна быть проставлена буква.
	Если элемент четный, перед ним добавляется префикс "E" (от слова Even — четный).
	Если элемент нечетный, перед ним добавляется префикс "O" (от слова Odd — нечетный).
Пример
Для входного массива [3, 8, 1, 8, 1]:
	Минимальное значение: 1
	Максимальное значение: 8
	Префиксы: O3, E8, O1
	Сгенерированный код будет: 1O3E8O1E8O18
*/
func generateCode(arr [5]int) string {
	min, max := arr[0], arr[0]
	middlePart := ""

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		if v%2 == 0 {
			middlePart += fmt.Sprintf("E%d", v)
		} else {
			middlePart += fmt.Sprintf("O%d", v)
		}
	}

	return fmt.Sprintf("%d%s%d", min, middlePart, max)
}




/*
Напишите функцию SumNeighbors, которая принимает массив из 10 целых чисел
и возвращает созданный новый массив такой же длины, 
где каждый элемент равен сумме своих соседей в исходном массиве. 
Для крайних элементов используйте только одного соседа. Новый массив необходимо вернуть из функции
*/

func SumNeighbors(arr [10]int) [10]int {
	result := [10]int{}
	for i := range arr {
		if i == 0 {
			result[i] = arr[i+1]
		} else if i == len(arr)-1 {
			result[i] = arr[i-1]
		} else {
			result[i] = arr[i-1] + arr[i+1]
		}
	}
	return result
}


func isPalindrome(nums [10]int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] != nums[j] {
			fmt.Println("It is not palindrome")
			return
		}
		i++
		j--
	}
	fmt.Println("It is palindrome")
}
