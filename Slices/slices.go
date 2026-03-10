package main

import (
	"fmt"
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


func sum(values ...int) int {   // variative parameter - можно передать сколько угодно int
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

	result := make([]int, minLen)    // range doesnot work with int
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
		result = append(result, arr1[i] + arr2[i])
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
