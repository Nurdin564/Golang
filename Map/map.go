package main

import (
	"fmt"
	"maps"
	"sort"

	// "maps"
	"slices"
	"strings"
	// "unicode"
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

	// m := map[string]int{
	// 	"banana":     2,
	// 	"apple":      1,
	// 	"grapefruit": 3,
	// 	"cherry":     1,
	// }
	// invertedMap := invertMap(m)
	// printMap(invertedMap)


	// text := "Да здравствует прекрасный язые, да здравствует golang!"
	// text = strings.Map(func (r rune) rune {
	// 	if unicode.IsPunct(r) {
	// 		return -1
	// 	}
	// 	return unicode.ToLower(r)
	// }, text)

	// words := strings.Split(text, " ")
	// wordsCount := make(map[string]int)
	// for _, word := range words {
	// 	wordsCount[word]++
	// }

	// for word, count := range wordsCount {
	// 	fmt.Printf("Word %q occurs in line %d times.\n", word, count)
	// }


	// input := map[string]int { 
	// 	"Mitchel Resnick": 5,  
	// 	"Linus Torvalds": 5,
	// 	"Donald Knuth": 3,
	// 	"Tim Berners-Lee": 4,
	// 	"Bjarne Stroustrup": 5,
	// }
	// fmt.Println(countVotes(input))


	// m1 := map[string][]int{
	// 	"a": {5, 2},
	// 	"b": {20, 11},
	// }
	// m2 := map[string][]int{
	// 	"a": {1, 5, 1},
	// 	"b": {11, 20},
	// }
	
	// result := maps.EqualFunc(m1, m2, func(v1, v2 []int) bool {  // смотрит значения по одинаковому ключу
	// 	sum1 := 0
	// 	for _, v := range v1 {
	// 		sum1 += v
	// 	}
	// 	sum2 := 0
	// 	for _, v := range v2 {
	// 		sum2 += v
	// 	}
	// 	return sum1 == sum2
	// })
	// fmt.Println(result)

	// fmt.Println(CompareMaxValues(m1, m2))


	// m1 := map[string]int{
	// 	"a": 5,
	// 	"b": 10,
	// }
	// m2 := map[string]int{
	// 	"a": 50,
	// 	"b": 10,
	// }
	// maps.Copy(m1, m2)
	// fmt.Println(m1)  // map[a:50 b:10]
	// m3 := maps.Clone(m1)
	// fmt.Println(m3)  // map[a:50 b:10]


	// m := map[string][]int{
	// 	"a": {5, 2},
	// 	"b": {20, 11},
	// 	"c": {1, 3, 9},
	// 	"d": {6, 111, 5},
	// }
	// maps.DeleteFunc(m, func(key string, value []int) bool {
	// 	return slices.Contains(value, 5)
	// 	// for _, v := range value {
	// 	// 	if v == 5 {
	// 	// 		return true
	// 	// 	}
	// 	// }
	// 	// return false
	// })
	// fmt.Println(m)  // map[b:[20 11] c:[1 3 9]]

	// m := map[string][]int{
	// 	"a": {1, 2, 3},
	// 	"b": {4, 5},
	// 	"c": {1, 1, 1},
	// 	"d": {2, 2, 3},
	// }
	// RemoveSlicesBySum(m)
	// fmt.Println(m)


	// m := map[string]map[string]float64{ 
	// 	"Экшен": {
	// 		"Фильм1": 8.52, 
	// 		"Фильм2": 6.0,
	// 	}, 
	// 	"Драма": {
	// 		"Фильм3": 7.524, 
	// 		"Фильм4": 7.527, 
	// 		"Фильм5": 5.54,
	// 	}, 
	// }
	// printRecommendations(m)


	friendsData := map[string][]string{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Елена":    {"Алексей", "Сергей", "Дмитрий"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}
	countFriends := countFriends(friendsData)
	users := make([]string, 0, len(countFriends))
	for user := range countFriends {
		users = append(users, user)
	}
	slices.Sort(users)

	fmt.Println("Count of friends:")
	for _, user := range users {
		fmt.Printf("%s: %d\n", user, countFriends[user])
	}

	user1, user2 := "Иван", "Елена"
	commonFriends := commonFriends(friendsData, user1, user2)
	slices.Sort(commonFriends)
	fmt.Printf("Общие друзья между пользователями %s и %s: %s.\n", user1, user2, strings.Join(commonFriends, ", "))

	popularUsers, maxFriends := mostPopularUsers(friendsData)
	fmt.Printf("Наиболее популярные пользователи: %s (количество друзей: %d).\n", strings.Join(popularUsers, ", "), maxFriends)

	
}
/*
В этот раз нам требуется разработать программу, которая будет анализировать данные о пользователях и их друзьях в социальной сети. 
Программа должна использовать структуру данных map для хранения информации о пользователях и их друзьях, 
а также реализовать функции для анализа этих данных.

Структура данных
Для хранения пользователей и их друзей будет использоваться map[string][]string. 
Ключом будет имя пользователя, а значением — список его друзей. 

Подсчет друзей
Реализуйте функцию countFriends, которая принимает map с пользователями и их друзьями и возвращает map, 
где ключом является имя пользователя, а значением — количество его друзей.

Общие друзья
Реализуйте функцию commonFriends, которая принимает map с пользователями и их друзьями, а также имена двух пользователей. 
Функция должна возвращать список общих друзей между этими двумя пользователями.

Наиболее популярных пользователей
Реализуйте функцию mostPopularUsers, которая принимает map с пользователями и их друзьями 
и возвращает список имен пользователей с наибольшим количеством друзей и это количество.

Вывод результатов
Программа должна вывести определенные данные:

Количество друзей для каждого пользователя, пользователи должны быть выведены в алфавитном порядке.
Список общих друзей между двумя заданными пользователями (например, "Иван" и "Елена", выбирайте сами), 
перечисление общих друзей должно быть в алфавитном порядке.
Имена наиболее популярных пользователей и количество их друзей, перечисление имен должно быть в алфавитном порядке.

Пример вывода
Количество друзей:
Алексей: 3
Анна: 1
Дмитрий: 3
Елена: 3
Иван: 3
Мария: 2
Ольга: 2
Петр: 2
Светлана: 2
Сергей: 2
Общие друзья между пользователями Иван и Елена: Алексей, Дмитрий.
Наиболее популярные пользователи: Алексей, Дмитрий, Елена, Иван (количество друзей: 3).
*/
func countFriends(m1 map[string][]string) map[string]int {
	result := map[string]int{}
	for name, friends := range m1 {
		result[name] = len(friends)
	}
	return result
}

func commonFriends(m1 map[string][]string, user1, user2 string) []string {
	friends1 := make(map[string]struct{})
	for _, friend := range m1[user1] {
		friends1[friend] = struct{}{}
	}

	commonFriends := []string{}
	for _, friend := range m1[user2] {
		if _, ok := friends1[friend]; ok {
			commonFriends = append(commonFriends, friend)
		}
	}
	return commonFriends
}

func mostPopularUsers(m1 map[string][]string) ([]string, int) {
	MaxFriends := 0
	users := []string{}

	for user, friends := range m1 {
		friendsCount := len(friends)
		if friendsCount > MaxFriends {
			users = []string{user}
			MaxFriends = friendsCount
		} else if MaxFriends == friendsCount {
			users = append(users, user)
		}
	}
	return users, MaxFriends
}




/*
Вам необходимо реализовать функцию printRecommendations, которая будет анализировать данные о фильмах 
и рекомендовать их к просмотру на основе заданных критериев. Функция должна принимать вложенный map[string]map[string]float64, 
где ключи первого уровня представляют жанры фильмов, а ключи второго уровня — названия фильмов с их соответствующими рейтингами.

Условия
Функция должна принимать один аргумент movies типа map[string]map[string]float64, где:

Ключ первого уровня (string) — это название жанра (например, "Экшен", "Драма").
Ключ второго уровня (string) — это название фильма.
Значение (float64) — это рейтинг фильма.
Функция должна выводить на экран все жанры в алфавитном порядке, в которых есть хотя бы один фильм с рейтингом 7 и выше.

Для каждого жанра, в котором есть фильмы с рейтингом 7 и выше, необходимо вывести названия всех фильмов (с рейтингом 7 и выше) в порядке убывания их рейтинга, если рейтинг одинаков, тогда сортировать такие фильмы нужно в алфавитном порядке.

Если в жанре нет фильмов с рейтингом 7 и выше, в таком случае, жанр вовсе не должен выводиться.

Пример
Входные данные:

m := map[string]map[string]float64{ 
  "Экшен": {
    "Фильм1": 8.52, 
    "Фильм2": 6.0,
  }, 
  "Драма": {
    "Фильм3": 7.524, 
    "Фильм4": 7.527, 
    "Фильм5": 5.54,
  }, 
}

                  
Ожидаемый вывод:

Драма: Фильм4 (7.5), Фильм3 (7.5).
Экшен: Фильм1 (8.5).
*/
func printRecommendations(m1 map[string]map[string]float64) {
	genres := []string{}
	for genre := range m1 {
		genres = append(genres, genre)
	}
	slices.Sort(genres)

	for _, genre := range genres {
		movies := m1[genre]

		filtered := []string{}
		for name, rating := range movies {
			if rating >= 7 {
				filtered = append(filtered, name)
			}
		}

		if len(filtered) == 0 {
			continue
		}

		sort.Slice(filtered, func(i, j int) bool {
			ri := movies[filtered[i]]
			rj := movies[filtered[j]]

			if ri != rj {
				return ri > rj // по убыванию рейтинга
			}
			return filtered[i] < filtered[j] // по алфавиту
		})

		fmt.Printf("%s: ", genre)
		for i, name := range filtered {
			fmt.Printf("%s (%.1f)", name, movies[name])
			if i < len(filtered)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println(".")
	}
}



/*
Необходимо написать функцию RemoveSlicesBySum, которая принимает на вход map типа map[string][]int, 
и удаляет из этой карты все слайсы, сумма элементов которых превышает 6.

func RemoveSlicesBySum(m map[string][]int) {
	for s, i := range m {
		sum := 0
		for _, j := range i {
			sum += j
		}
		if sum > 6 {
			delete(m, s)
		}
	}
}
*/
func RemoveSlicesBySum(m map[string][]int) {
	maps.DeleteFunc(m, func(key string, value []int) bool {
		sum := 0
		for _, v := range value {
			sum += v
		}
		return sum > 6
	})
}


/*
Необходимо реализовать функцию mergeMaps, которая принимает два аргумента типа map[string]int
и возвращает новый мап типа map[string]int. Функция должна объединять значения из двух мап, 
суммируя значения для одинаковых ключей. Если ключ присутствует только в одной из мап, 
он должен быть добавлен в результирующую мапу с соответствующим значением.
Пример
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 3, "c": 4, "d": 5}
result := mergeMaps(m1, m2)
// result будет равен map[string]int{"a": 1, "b": 5, "c": 7, "d": 5}
*/
func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := map[string]int{}
	maps.Copy(result, m1)
	for j, k := range m2 {
		result[j] += k
	}
	return result
}



/*
Напишите функцию CompareMaxValues, которая принимает два параметра: два map типа map[string][]int. 
Функция должна сравнить максимальные значения в срезах для каждого ключа. 
Если максимальные значения для всех соответствующих ключей в обоих map равны, функция должна вернуть true, в противном случае — false

func CompareMaxValues(m1, m2 map[string][]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, arr1 := range m1 {
		arr2, ok := m2[key]
		if !ok {
			return false
		}

		if len(arr1) == 0 && len(arr2) == 0 {
			continue
		}

		if len(arr1) == 0 || len(arr2) == 0 {
			return false
		}

		max1 := arr1[0]
		max2 := arr2[0]

		for _, v := range arr1 {
			if v > max1 {
				max1 = v
			}
		}
		for _, v := range arr2 {
			if v > max2 {
				max2 = v
			}
		}

		if max1 != max2 {
			return false
		}

	}
	return true
	
}
*/
func CompareMaxValues(m1, m2 map[string][]int) bool {
	return maps.EqualFunc(m1, m2, func(v1, v2 []int) bool {
		if len(v1) == 0 && len(v2) == 0 {
			return true
		}

		if len(v1) == 0 || len(v2) == 0 {
			return false
		}

		return slices.Max(v1) == slices.Max(v2)
	})
}


/*
Подсчет голосов
Необходимо реализовать функцию countVotes, которая будет определять кандидата с наибольшим количеством голосов на выборах.
	Функция countVotes принимает один аргумент: votes типа map[string]int, 
	где ключами являются имена кандидатов (типа string), а значениями — количество голосов, полученных каждым кандидатом (типа int). 
Функция countVotes должна возвращать имя кандидата (типа string), который получил наибольшее количество голосов.
	В случае, если несколько кандидатов имеют одинаковое максимальное количество голосов, 
	верните имена всех кандидатов через запятую, в алфавитном порядке. 
Если в аргументе votes не будет ни одного кандидата, функция должна вернуть строку Кандидаты потерялись.
	Если кандидаты были, но ни у одного кандидата не было ни одного голоса, необходимо вернуть строку Все голоса похищены!
*/
func countVotes(votes map[string]int) string {
	if len(votes) == 0 {
		return "Кандидаты потерялись"
	}

	maxVotes := 0
	for _, v := range votes {
		if v > maxVotes {
			maxVotes = v
		}
	}

	if maxVotes == 0 {
		return "Все голоса похищены!"
	}

	winners := make([]string, 0)
	for name, v := range votes {
		if v == maxVotes {
			winners = append(winners, name)
		}
	}

	slices.Sort(winners)
	return strings.Join(winners, ", ")

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
