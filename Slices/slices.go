package main

import (
	"fmt"
	"math/rand"
	"slices"
	"unicode/utf8"

	// "sort"

	// "slices"
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

	// numbers := []int{10, 20, 30, 40, 50}

	// ptr := &numbers[2]
	// fmt.Println("Value:", *ptr)  // 30
	// fmt.Println("Address:", ptr)  // 0x8010

	// numbers = append(numbers, 60)

	// ptr2 := &numbers[2]
	// fmt.Println("Value:", *ptr2)  // 30
	// fmt.Println("Adress:", ptr2)  // 0x20b0

	// *ptr = 290  // does not changed
	// fmt.Println(numbers)  // after append it will be new slice [10 20 30 40 50 60], cause cap was 5 and after adding it will be another slice

	// slice := []int{1, 2, 3}
	// value := 0
	// slice = append([]int{value}, slice...)  // добавление в началао
	// fmt.Println(slice)  // [0 1 2 3]

	// slice := []int{1, 2, 4, 5}
	// index := 2
	// value := 3

	// // Adding to middle
	// // best case
	// before := slice[:index]
	// after := append([]int{value}, slice[index:]...)  // добавление в середину
	// slice = append(before, after...)
	// fmt.Println(slice)  // [1 2 3 4 5]

	// worst case (don't add element at first part of slice cause 4 will be 3)
	// before := append(slice[:index], value)
	// fmt.Println(before, slice)  // [1 2 3] [1 2 3 5]
	// after := slice[index:]
	// slice = append(before, after...)
	// fmt.Println(slice)  // [1 2 3 3 5]

	// slice := []int{23, 444, 0, -3, 12, 8, 60, 3}
	// fmt.Println(PlayWithSlice(slice))

	// fmt.Println(insertElement([]int{1,2,3,4,5,6,7,89,9}, 5, 5000))

	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Printf("slice: %v, capacity: %d\n", slice, cap(slice))

	// slice = slice[:len(slice) - 1]
	// fmt.Printf("slice: %v\n", slice)  // [1 2 3 4]

	// slice = slice[1:]
	// fmt.Printf("slice: %v\n", slice)  // 2, 3, 4, 5

	// indexToRemove := 2
	// slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	// fmt.Printf("slice: %v, capacity: %d\n", slice, cap(slice))  // slice: [1 2 4 5], capacity: 5

	// slice = slices.Clip(slice)
	// fmt.Printf("slice: %v, capacity: %d\n", slice, cap(slice))  // slice: [1 2 4 5], capacity: 4. But in memory we still have 5 cell

	// newSlice := make([]int, len(slice))
	// copy(newSlice, slice)
	// fmt.Printf("slice: %v, capacity: %d\n", newSlice, cap(newSlice))  // slice: [1 2 4 5], capacity: 4

	// arr := []int{1, 2, 3, 4, 5, 6, 11}
	// fmt.Println(DeletingFromSlice(arr))

	// s := []int{1,2,3,4,5}
	// fmt.Println("Original slice:", s)
	// fmt.Println(changeSlice(s))

	// strings := []string{"banana", "apple", "cherry"}
	// slices.Sort(strings)
	// fmt.Println(strings)

	// numbers := []int{4,2,7,1,3}
	// slices.SortStableFunc(numbers, func(a, b int) int {
	// 	return a - b
	// })
	// fmt.Println(numbers)

	// slice := [][]int{
	// 	{3, 1, 4, 1},
	// 	{2, 2, 2},
	// 	{5, 0, 6, 3, -8, 1},
	// 	{4, 6, 8, 2},
	// }
	// fmt.Println(magicSort(slice))

	// numbers := []int{1, 2, 3, 4, 5, 6}
	// newNumbers := slices.Insert(numbers, 2, 90) // (slice, index, value)
	// fmt.Println(newNumbers)

	// fmt.Println(slices.Contains(numbers, 40))
	// contains := slices.ContainsFunc(numbers, func(v int) bool {
	// 	return v%2 == 0
	// })
	// fmt.Println(contains)

	// numbers1 := []int{1, 2, 3, 2, 5}
	// numbers2 := []int{5, 3, 1, 4, 5}
	// // compared := slices.Compare(numbers1, numbers2)
	// compared := slices.CompareFunc(numbers1, numbers2, func(a, b int) int {
	// 	if a%2 == 1 && b%2 == 1 {
	// 		return 0
	// 	}
	// 	return a - b
	// })
	// fmt.Println(compared)

	// combined := slices.Concat(numbers1, numbers2, []int{100, 233})
	// fmt.Println(combined)

	// fmt.Println(slices.Index(numbers1, 3))
	// idx := slices.IndexFunc(numbers1, func(v int) bool {
	// 	return v%2 == 0
	// })
	// fmt.Println(idx)

	// fmt.Println(slices.Min(numbers1))
	// fmt.Println(slices.Max(numbers1))

	// newNumbers := slices.Delete(numbers1, 2, 4)  // from index to index
	// newNumbers := slices.DeleteFunc(numbers1, func(v int) bool {
	// 	return v%2 == 0
	// })
	// fmt.Println(newNumbers)

	// slices.Reverse(numbers1)
	// fmt.Println(numbers1)  // если функция ничего не возвращает то нужно отдельно ее задействовать и потом печатать

	// fmt.Println(slices.Repeat(numbers1, 5))

	// newNumbers := slices.Replace(numbers1, 1, 3, 34, 45, 99)
	// fmt.Println(newNumbers)  // [1 34 45 99 2 5]

	// fmt.Println(slices.Equal(numbers1, numbers2))
	// equals := slices.EqualFunc(numbers1, numbers2, func(v1, v2 int) bool {
	// 	return v1 == v2 || v1%2 == v2%2  // 1 == 1
	// })
	// fmt.Println(equals)

	// isSorted := slices.IsSorted(numbers1)
	// fmt.Println(isSorted)

	// idx, found := slices.BinarySearch(numbers1, 80)
	// fmt.Println(idx, found)


	// users := []User{
	// 	{"Pavel", 33},
	// 	{"Andrew", 36},
	// 	{"Olga", 41},
	// 	{"Elena", 44},
	// 	{"Ivan", 58},
	// }
	// idx, found := slices.BinarySearchFunc(users, 44, func(v User, target int) int {
	// 	return v.Age - target
	// })
	// fmt.Println(idx, found)


	// str := "Зашифруй меня!"
	// encodedStr := CaesarCode(str, 5, true)
	// fmt.Println(encodedStr)
	
	// decodedStr := CaesarCode(encodedStr, 5, false)
	// fmt.Println(decodedStr)

	slice, err := CreateSlice(5)
	if err != nil {
		fmt.Println(err)
		return	
	} 
	fmt.Println(slice)


}



// Функция для создания слайса
// 	Функция CreateSlice(n int) ([]int, error) принимает целое число n и возвращает слайс
// 	случайных целых чисел от -10 до 10 включительно, длиной в n. 
// 	Если n отрицательное, функция должна возвращать ошибку, а программа должна вывести данную ошибку и завершить свою работу.

func CreateSlice(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("negative length")
	}

	slice := make([]int, n)
	for i := range n {
		slice[i] = rand.Intn(21) - 10  // [0, 20] - 10 = [-10, 10]
	}
	return slice, nil
	
}

// Функция для фильтрации слайса
// 	Функция FilterSlice(numbers []int) []int принимает слайс целых чисел и возвращает новый слайс, 
// 	который будет содержать только те числа, которые соответствуют следующим условиям:

// Перед числом, которое войдет в слайс с результатом, должно стоять число больше (в переданном в функцию слайсе).
// Другими словами - число, которое будет добавлено в результирующий слайс, должно следовать за числом, которое больше него в исходном слайсе.
// Число должно делиться без остатка на одно из чисел: 2, 5, 6, 9.
// Пример:
// 	Переданный в функцию слайс: [1, 6, 2, 5, 8, 6]
// 	Результат, который вернет функция: [2, 6], так как 6 > 2 и 8 > 6, а также оба числа делятся на одно из чисел 2, 5, 6 или 9.

// func FilterSlice(numbers []int) []int {
// 	result := make([]int, len(numbers))

	
// }



// Функция для нахождения максимальной суммы подслайса
// 	Функция MaxSumWithNegative(numbers []int, k int) []int принимает слайс целых чисел и целое число k. 
// 	Она должна вернуть новый слайс, который будет содержать значения подслайса длиной k, 
// 	который имеет максимальную сумму и содержит хотя бы одно отрицательное число.
// 	Если такого слайса не будет, необходимо вернуть nil.

// Под "подслайсом" понимается последовательно расположенные значения в оригинальном слайсе.
// Пример:
// 	Представим, что функция вызывается таким образом: MaxSumWithNegative([]int{-1, 6, 2, 5, 8, -2}, 3)
// 	Функция должна найти и вернуть слайс с наибольшей суммой значений среди всех подслайсов, в которых будет три значения. 
// 	Определим какие подслайсы есть в переданном слайсе:
// 	-1, 6, 2
// 	6, 2, 5
// 	2, 5, 8
// 	5, 8, -2

// 	Необходимо найти подслайс с самой большой суммой значений и вернуть его, при этом, 
// 	в подслайсе должно быть как минимум одно отрицательное число. 
// 	В данном примере это будет последний слайс, именно его нам и нужно вернуть.





// Функция для сортировки слайса
// 	Функция SortByParity(numbers []int) []int должна отсортировать слайс. 
// 	Четные числа в слайсе должны быть в начале, отсортированы по убыванию, 
// 	а нечетные в конце слайса и отсортированы по возрастанию.






func CaesarCode(text string, shift int32, encode bool) string {
	result := make([]rune, 0, utf8.RuneCountInString(text))
	for _, r := range text {
		if encode {
			result = append(result, r+shift)
		} else {
			result = append(result, r-shift)
		}
	}
	return string(result)
}

// func CaesarCode(text string, shift int32, encode bool) string {
// 	if !encode {
// 		shift = -shift
// 	}

// 	result := []rune{}

// 	for _, ch := range text {
// 		if ch >= 'А' && ch <= 'я' {

// 			start := int32('А')
// 			end := int32('я')
// 			size := end - start + 1  // 33 - 1 + 1 = 33

// 			newChar := ch + shift

// 			if newChar > end {
// 				newChar = start + (newChar-end-1)%size
// 			}
// 			if newChar < start {
// 				newChar = end - (start-newChar-1)%size
// 			}

// 			result = append(result, rune(newChar))

// 		} else {
// 			result = append(result, ch)
// 		}
// 	}
// 	return string(result)
// }




type User struct {
	Name string
	Age int
}

/*
Необходимо реализовать функцию magicSort,
которая принимает двумерный слайс целых чисел и сортирует его по следующим критериям:
	Внешний слайс должен быть отсортирован по возрастанию суммы чисел в каждом внутреннем слайсе,
	если слайсы имеют одинаковую сумму, то они должны быть в том же порядке, что и были.
Каждый внутренний слайс должен быть отсортирован в порядке убывания,
при этом четные числа должны идти перед нечетными. Ноль идет перед всеми числами.
*/

func magicSort(arr [][]int) [][]int {
	for _, s := range arr {
		slices.SortStableFunc(s, func(a, b int) int {
			if a == 0 && b != 0 {
				return -1
			}
			if b == 0 && a != 0 {
				return 1
			}

			if a%2 == 0 && b%2 != 0 {
				return -1
			}
			if a%2 != 0 && b%2 == 0 {
				return 1
			}

			return b - a

		})
	}

	sum := func(s []int) int {
		total := 0
		for _, v := range s {
			total += v
		}
		return total
	}

	slices.SortStableFunc(arr, func(a, b []int) int {
		return sum(a) - sum(b)
	})

	return arr
}

func changeSlice(slice []int) []int {
	slice[len(slice)-1] = 100
	slice = append(slice, -1)
	slice[0] = 500
	return slice
}

/*
Необходимо создать функцию с названием DeletingFromSlice, которая принимает слайс целых чисел
и возвращает новый слайс целых чисел. Функция должна выполнять следующие действия:
	Удалить из принятого слайса последнее значение, если это значение существует и оно больше 10.
	Удалить значение по индексу 2, если такое значение есть и вместимость слайса больше 5.
	Удалить первое значение из слайса, если оно присутствует и были удалены значения, указанные в первых двух пунктах.
	Убрать лишнюю вместимость у слайса.
Функция должна вернуть полученный слайс.
*/

func DeletingFromSlice(arr []int) []int {
	slice := make([]int, len(arr))
	copy(slice, arr)

	removedLast := false
	removedSecond := false

	if len(slice) > 0 && slice[len(slice)-1] > 10 {
		slice = slice[:len(slice)-1]
		removedLast = true
	}

	if len(slice) > 2 && cap(slice) > 5 {
		slice = append(slice[:2], slice[3:]...)
		removedSecond = true
	}

	if removedLast && removedSecond && len(slice) > 0 {
		slice = slice[1:]
	}

	slice = append([]int{}, slice...)
	return slice
}

// Example from comments
func insertElement(slice []int, pos int, value int) []int {
	n := len(slice)
	newSlice := make([]int, n+1)
	copy(newSlice[:pos], slice[:pos])    // Копируем левую половину
	newSlice[pos] = value                // Добавляем новый элемент
	copy(newSlice[pos+1:], slice[pos:n]) // Копируем правую половину
	return newSlice
}

/*
Необходимо реализовать функцию PlayWithSlice, которая будет принимать слайс целых чисел
и выполнять с ним несколько операций и возвращать новый слайс целых чисел:

	Клонирование слайса: Создайте новый слайс, который будет являться клоном переданного слайса.
	Все дальнейшие операции должны выполняться только с клоном.

	Вставка числа 100: Найдите первое значение с конца клона, которое больше 10.
	После этого значения вставьте число 100. Если такого значения не найдено, этот шаг можно пропустить.

	Вставка числа 500: Если сумма всех чисел в текущем клоне больше 100, добавьте число 500 в конец слайса.

	Вставка числа 1000: Если в оригинальном слайсе четных чисел больше, чем нечетных, вставьте число 1000 в начало клона слайса.

Функция должна вернуть модифицированный слайс.
Дополнительные указания

	Обратите внимание на то, что слайсы в Go передаются по ссылке,
	поэтому вам нужно создать новый слайс, чтобы избежать изменения оригинального.
	Убедитесь, что ваша функция корректно обрабатывает пустые слайсы и слайсы, не содержащие чисел, удовлетворяющих условиям.
*/
func PlayWithSlice(arr []int) []int {
	clone := make([]int, len(arr))
	copy(clone, arr)

	for i := len(clone) - 1; i >= 0; i-- {
		if clone[i] > 10 {
			newSlice := []int{}
			newSlice = append(newSlice, clone[:i+1]...)
			newSlice = append(newSlice, 100)
			newSlice = append(newSlice, clone[i+1:]...)
			break
		}
	}

	sum := 0
	for _, v := range clone {
		sum += v
	}

	if sum > 100 {
		clone = append(clone, 500)
	}

	even := 0
	odd := 0
	for _, v := range arr {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		clone = append([]int{1000}, clone...)
	}

	return clone
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
