package main

import (
	"fmt"
	"slices"
)

func main() {
	// var m map[string]int
	// fmt.Println(m == nil) // true
	// // m["sgsgsgsgsg"] = 1  // cant add in nil map
	
	// n := make(map[string]int)
	
	// s := map[string]int {
	// 	"apple": 2,
	// 	"banana": 13,
	// }

	// fmt.Println(s["fdfdefdf"])  // zero value: 0
	// val, ok := s["fdfdefdf"]
	// fmt.Println(val, ok)  // 0, false


	// value, exists := s["banana12313"]  // 13
	// if exists {
	// 	fmt.Println("Value:", value)
	// } else {
	// 	fmt.Println("not found")
	// }

	// fmt.Println(m, n, s)  // map[](nill), map[], map[apple:2 banana:13]

	// for key, value := range s {
	// 	fmt.Printf("key: %s, value: %d\n", key, value)
	// }

	// delete(s, "apple")
	// fmt.Println(s)


	// nums := []int{1, 2, 2, 3, 3, 3}
	// fmt.Println(CountMaxFrequency(nums))

	// m := make(map[int]int, 10)
	// for i := range 10 {
	// 	m[i] = i
	// }

	// s := make([]int, 0, 10) 
	// for key := range m {
	// 	if key == 5 {
	// 		delete(m, 7)  // don't change map in for loop !!!
	// 	}
	// 	s = append(s, key)
	// }

	// fmt.Printf("Slice: %v, count of values: %d\n", s, len(s))

	// points := make(map[any]string)   // only hashable and comparable types
	// points[Point{1, 2}] = "Point A"
	// points[123] = "Point B"
	// points["asdfg"] = "Point C"
	// fmt.Println(points)

	// m := make(products)

	// m["fruits"] = map[string]int{
	// 	"apple": 8,
	// 	"banana": 10,
	// }
	// m["vegetables"] = map[string]int{
	// 	"carrot": 2,
	// }

	// fmt.Println(m)
	// fmt.Println(m["fruits"])  // map[apple:8 banana:10]
	// fmt.Println(m["fruits"]["banana"])  // 10
	// fmt.Println(m["hello"] == nil)  // true
	// fmt.Println(m["hello"]["hello"])  // zero value 0

	// val, ok := m["fruits"]["banana"]
	// fmt.Println(val, ok)  // 10, true  or  0, false

	m := map[string]int{
		"banana":     2,
		"apple":      1,
		"grapefruit": 3,
		"cherry":     1,
	}
	invertedMap := invertMap(m)
	printMap(invertedMap)

}

/*
Вам необходимо реализовать две функции, первая - invertMap, которая будет инвертировать входную map, 
вторая - printMap, которая будет выводить данные, полученные от invertMap в правильном формате.
	Функция invertMap должна принимать map типа map[string]int и возвращать новую map типа map[int][]string. 
	В новой map ключами будут значения из исходной map, а значениями — срезы ключей, которые соответствуют этим значениям.
Функция printMap должна принимать то, что возвращает функция invertMap, 
то есть, функция printMap должна принимать map[int][]string и выводить эти данные в определенном формате

// Ваш код
func invertMap(m map[string]int) map[int][]string {
    im := make(map[int][]string)
    for key, value := range m {
        im[value] = append(im[value], fmt.Sprintf("\"%s\"", key))
    }
    return im
}

func printMap(m map[int][]string) {
    keys := make([]int, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    slices.Sort(keys)
    fmt.Println("{")
    for _, key := range keys {
        values := m[key]
        slices.Sort(values)
        fmt.Printf("  %d: [%s],\n", key, strings.Join(values, ", "))
    }
    fmt.Println("}")
}
*/
func invertMap(m map[string]int) map[int][]string {
	result := make(map[int][]string)

	for key, value := range m {
		result[value] = append(result[value], key)
	}
	return result
}

func printMap(m map[int][]string) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	fmt.Println("{")

	for i, k := range keys {
		values := m[k]
		slices.Sort(values)

		fmt.Printf("  %d: [", k)

		for j, v := range values {
			fmt.Printf("\"%s\"", v)
			if j != len(values)-1 {
				fmt.Print(", ")
			}
		}

		fmt.Print("]")

		if i != len(keys)-1 {
			fmt.Println(",")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("}")
}





type products map[string]map[string]int

type Point struct {
	X int
	Y int
}

/*
Вам необходимо реализовать функцию CountMaxFrequency, которая будет определять,
сколько раз встречается самое часто упоминаемое значение в переданном слайсе целых чисел.
Функция CountMaxFrequency должна принимать один аргумент - слайс целых чисел. 
Она должна проанализировать этот слайс и вернуть количество повторений самого часто встречающегося числа. 
Если в слайсе несколько чисел имеют одинаковую максимальную частоту, функция должна вернуть количество повторений любого из них.
*/
func CountMaxFrequency(arr []int) int {
	result := map[int]int{}
	max := 0

	for _, num := range arr {
		result[num]++
	}

	for _, count := range result {
		if count > max {
			max = count
		}
	}
	return max
}
