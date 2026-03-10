package main

import (
	"fmt"
	// "unicode/utf8"

	// "slices"
	"strings"
)

func main() {
	// slice := []int{6, 4, 5, 9, 7}
	// fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	// fmt.Println(slice[1])

	// slice1 := make([]int, 5, 10)
	// fmt.Printf("slice: %v, len: %d, cap: %d\n", slice1, len(slice), cap(slice))

	// var s []int
	// fmt.Println(s[0])  // index out of range [0] with length 0

	// user := GetUserParam("https://site.ru/page?user=pavel&id=12345&limit=20")
	// fmt.Println(user)

	// numbers := []int{6, 4, 5, 9, 7}
	// sum := 0

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }
	// for _, v := range numbers {
	// 	sum += v
	// }
	// fmt.Println("Sum of slice's elements:", sum)
	// fmt.Println(numbers)  // [6 4 5 9 7]

	// for _, v := range numbers {
	// 	v += 10
	// }
	// fmt.Println(numbers) // [6 4 5 9 7] original slice didn't change

	// for i := range numbers {
	// 	numbers[i] *= 10
	// }
	// fmt.Println(numbers) // [60 40 50 90 70]

	// a := []int{1,2,3,4,5}
	// a = append(a, 7, 8, 9)
	// fmt.Println(a, cap(a))   // cap(10)
	// a = append(a, 2, 78, 90, 23, 45, 67)
	// fmt.Println(a, cap(a))   // cap(20)

	// fmt.Println(printMagic([]int{1,2,3,4}))

	// b := []int{34, 56, 67, 78, 89, 90}
	// sum := SumSlices(a, b)
	// sum1 := SumSlices1(a, b)
	// fmt.Println(sum, sum1)

	// fmt.Println("Sum:", sum(1,2,3,4,5))
	// s := []int{1,2,3,4,5,6,7,8,9}
	// fmt.Println("Sum:", sum(s...))   // spread operator - передает все содержимое слайса

	// s1 := []int{1,2,3}
	// s2 := []int{6,9}
	// s1 = append(s1, s2...)
	// fmt.Println(s1)

	// q := []int{1,2,3,4,5,6,7,8,9}
	// fmt.Println(filterEven(q...))

	// slices can be compared only with nil
	// s1 := []string{"Sup", "nigga"}
	// s2 := []string{"Sup1", "nigga1"}
	// fmt.Println(s1 == s2)

	// var s3 []int  // nil
	// fmt.Println(s3, len(s3), cap(s3))

	// s := [][]int{
	// 	{2,3,4,5},
	// 	{0,3,6,5,78, 8},
	// 	{8,7,3,2,3,5,6,7,},
	// }
	// fmt.Println(s)
	// fmt.Println(s[1])
	// fmt.Println(s[1][0])

	// s[2] = append(s[2], 100)
	// fmt.Println(s)

	// s[2][2] = -4
	// fmt.Println(s, len(s[2]))

	// s = append(s, []int{-1, -2, -3})
	// fmt.Println(s)

	// sum := 0
	// for _, innerSlice := range s {
	// 	for _, num := range innerSlice {
	// 		if num%2 == 1 {
	// 			sum += num
	// 		}
	// 	}
	// }
	// fmt.Println(sum)

	// fmt.Println(replaceEvenOnEvenIndices(s))

	// s1 := []int{10, 12, 11}  // 3
	// s2 := []int{1, 2, 3, 4, 5}  // 5
	// copy(s1, s2)  // in s1 put s2, return minimum length of arrays
	// fmt.Println(s1, s2)  // [1 2 3] [1 2 3 4 5]

	// s3 := []int{10, 12, 11}
	// s4 := s3  // pointer to same address
	// s4[0] = 100
	// fmt.Println(s3, s4)  // [100 12 11] [100 12 11]

	// s5 := []int{10, 12, 11}
	// s6 := make([]int, len(s5))
	// copy(s6, s5)
	// s6[0] = 678
	// fmt.Println(s5, s6)  // [10 12 11] [678 12 11]

	// s1 := []int{10, 12, 11}  // 3
	// s2 := slices.Clone(s1)
	// s2[0] = 458
	// fmt.Println(s1, s2)  // [10 12 11] [458 12 11]

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := []int{3, 4, 5, 6, 7}
	// fmt.Println(intersectSlices(slice1, slice2))

	// original := []int{1, 2, 3, 4, 5}
	// subSlice := original[1:4]
	// fmt.Println(subSlice)

	// str := "Ку-ку!"
	// str1 := str[:4] // take bytes
	// fmt.Println(str1)
	// fmt.Println(utf8.ValidString(str1))

	// original1 := make([]int, 5, 10)
	// copy(original1, []int{1, 2, 3, 4, 5})
	// fmt.Println(original1, len(original1), cap(original1))

	// subSlice1 := original1[1:10]
	// fmt.Println(subSlice1)

	// original := []int{1, 2, 3, 4, 5}
	// fmt.Println(original)

	// subSlice := original[1:4]
	// fmt.Println(subSlice)

	// subSlice[1] = 243
	// fmt.Println(subSlice)  // [2 243 4]
	// fmt.Println(original) // [1 2 243 4 5]

	numbers := []int{10, 20, 30, 40, 50}

	ptr := &numbers[2]
	fmt.Println("Value:", *ptr)  // 30
	fmt.Println("Address:", ptr)  // 0x8010

	numbers = append(numbers, 60)

	ptr2 := &numbers[2]
	fmt.Println("Value:", *ptr2)  // 30
	fmt.Println("Adress:", ptr2)  // 0x20b0

	*ptr = 290  // does not changed
	fmt.Println(numbers)  // after append it will be new slice [10 20 30 40 50 60], cause cap was 5 and after adding it will be another slice

}

/*
Необходимо реализовать функцию intersectSlices, которая будет принимать
два отсортированных слайса целых чисел и возвращать новый слайс,
содержащий все пересечения этих слайсов (тоже в отсортированном порядке), а также ошибку.
Если один из переданных слайсов равен nil, функция должна возвращать ошибку с сообщением slices cannot be nil.
*/
func intersectSlices(arr1, arr2 []int) ([]int, error) {
	if arr1 == nil || arr2 == nil {
		return nil, fmt.Errorf("slices cannot be nil")
	}

	result := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			result = append(result, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return result, nil
}

/*
Напишите функцию replaceEvenOnEvenIndices, которая принимает двумерный слайс целых чисел
и возвращает новый двумерный слайс целых чисел.
В возвращаемом слайсе должны быть те же значения, что у принятого слайса, однако, все четные числа,
находящиеся на четных индексах (включая 0, в каждом из собственных внутренних слайсов), должны быть заменены на 0.
*/
func replaceEvenOnEvenIndices(arr [][]int) [][]int {
	result := make([][]int, len(arr))

	for i, row := range arr {
		newRow := make([]int, len(row))

		for j, v := range row {
			if j%2 == 0 && v%2 == 0 {
				newRow[j] = 0
			} else {
				newRow[j] = v
			}
		}
		result[i] = newRow
	}
	return result

}

/*
Вам необходимо реализовать функцию Max, которая будет находить максимальный элемент в слайсе целых чисел.
Функция должна принимать слайс в качестве аргумента и возвращать максимальное значение и возможную ошибку.
Если слайс равен nil или пуст, функция должна возвращать ошибку с сообщением slice is nil or empty
*/

func Max(arr []int) (int, error) {
	if arr == nil || len(arr) == 0 {
		return 0, fmt.Errorf("slice is nil or empty")
	}
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max, nil
}

/*
Необходимо реализовать функцию filterEven, которая будет фильтровать четные числа из переданного набора целых чисел.
Функция должна принимать переменное количество аргументов (целых чисел) и возвращать слайс,
содержащий только четные числа, сохраняя порядок переданных в функцию значений в возвращенном слайсе.
*/
func filterEven(numbers ...int) []int {
	result := []int{}
	for _, v := range numbers {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func sum(values ...int) int { // variative parameter - можно передать сколько угодно int
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

/*
Необходимо написать функцию SumSlices, которая будет выполнять сложение двух слайсов целых чисел.
Функция должна принимать два слайса в качестве аргументов и возвращать новый слайс,
содержащий суммы соответствующих элементов из входных слайсов.
Если длины двух слайсов различаются, необходимо игнорировать лишние элементы более длинного слайса.
То есть, если один слайс длиннее другого, то при сложении следует учитывать только те элементы, которые находятся в пределах длины более короткого слайса.
Результирующий слайс должен содержать суммы элементов, расположенных на одинаковых позициях в исходных слайсах.
*/

func SumSlices(arr1, arr2 []int) []int {
	minLen := len(arr1)
	if len(arr2) < minLen {
		minLen = len(arr2)
	}

	result := make([]int, minLen) // range doesnot work with int
	for i := range result {
		result[i] = arr1[i] + arr2[i]
	}
	return result
}

func SumSlices1(arr1, arr2 []int) []int {
	minArr := len(arr1)
	if minArr > len(arr2) {
		minArr = len(arr2)
	}

	result := []int{}
	for i := 0; i < minArr; i++ {
		result = append(result, arr1[i]+arr2[i])
	}
	return result
}

/*
Напишите функцию printMagic, которая принимает слайс целых чисел и выводит строку, похожую на массив, в консоль.
Каждое выведенное число - это произведение всех остальных элементов исходного массива.

Пример
Для входного массива [1, 2, 3, 4] программа должна вывести [24, 12, 8, 6], где:

Первый элемент 24 — это произведение 2 * 3 * 4
Второй элемент 12 — это произведение 1 * 3 * 4
Третий элемент 8 — это произведение 1 * 2 * 4
Четвертый элемент 6 — это произведение 1 * 2 * 3
*/
func printMagic(numbers []int) string {
	result := []int{}

	for i := range numbers {
		product := 1

		for j := range numbers {
			if j == i {
				continue
			}
			product *= numbers[j]
		}
		result = append(result, product)
	}
	return fmt.Sprintln(result)
}

/*
Необходимо будет создать функцию с названием GetUserParam, которая будет принимать строку с записанным в ней url.
Нужно будет вернуть из функции строку, в которой будет записано значение для параметра user.

Для примера, если в функцию будет передано значение:
https://site.ru/page?user=pavel&id=12345&limit=20

Функция должна найти GET-параметры (это часть после знака ?), найти среди параметров значение под ключом user
и вернуть это значение. В данном примере, мы должны вернуть значение pavel.

Данный параметр user может быть в середине, в конце списка GET-параметров, а может и вовсе отсутствовать.
Если параметра user нет, необходимо будет вернуть строку с текстом not found.

Примечание:
Для решения этой задачи есть готовые пакеты, для парсинга URL, однако, мне бы хотелось,
чтобы вы использовали базовые операции со строками (включая те что есть в пакете strings).
Обратите внимание, если значения у параметра user нет,
то необходимо будет вернуть строку not found, как и в случае отсутствия ключа.
Если ключ user повторяется в GET-параметрах, необходимо вернуть значение первого ключа.
*/
func GetUserParam(url string) string {
	parts := strings.Split(url, "?") // Split(строка, разделитель), ["https://site.ru/page", "user=pavel&id=12345&limit=20"]
	if len(parts) < 2 {
		return "not found"
	}

	second := parts[1]
	params := strings.Split(second, "&") // ["user=pavel", "id=12345", "limit=20"]

	for _, v := range params {
		pair := strings.Split(v, "=") // ["user", "pavel"], ["id", 12345], ["limit", 20]

		if pair[0] == "user" {
			if len(pair) > 1 && pair[1] != "" {
				return pair[1]
			}
			return "not found"
		}
	}
	return "not found"
}
