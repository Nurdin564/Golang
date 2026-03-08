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


	fmt.Println(printMagic([]int{1,2,3,4}))


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
