package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "math"
	// "os"
	// "math"
	// "math/rand/v2"
	// "strings"
)

func main() {
	// fmt.Println("Ah shit, here we go again!")

	// a := 10
	// fmt.Println("a % 3: ", a%3) // 1
	// b := 10.5
	// fmt.Println("b % 3: ", math.Mod(b, 3)) // если находить остаток от деления от числа float64, то юзаем Mod

	// /*
	// 	Sample Input:
	// 	44.4
	// 	Sample Output :
	// 	Исходное число: 44.4
	// 	Исходное число, увеличенное на 10%: 48.84000
	// 	Исходное число является четным: false
	// 	Предпоследняя цифра целой части исходного числа: 4
	// */
	// random := math.Floor((rand.Float64()*100)*10) / 10

	// fmt.Printf("Исходное число: %.1f\n", random)
	// fmt.Printf("Исходное число, увеличенное на 10%%: %.5f\n", random*1.1)
	// even := int(random)%2 == 0 && (random-math.Trunc(random)) == 0 // random (return float64) - math.Trunc(random) (return float64). Нельзя юзать int(random) вместо math.Trunc(random), потомучто random float64
	// fmt.Println("Исходное число является четным:", even)
	// fmt.Println("Предпоследняя цифра целой части исходного числа:", int(random)/10%10) // 343.346 -> 343 -> 343/10 (целочисленное деление) -> 34 -> 34 % 10 == 4

	// str1 := "Hi, "
	// str2 := "bitch!"

	// result := str1 + str2
	// fmt.Println(result)

	// result2 := fmt.Sprintf("%s%s", str1, str2) // чтобы хранить в переменной кастомные строки
	// fmt.Println(result2)

	// result3 := strings.Join([]string{str1, str2}, "")
	// fmt.Println(result3)

	// var buffer strings.Builder
	// buffer.WriteString(str1)
	// buffer.WriteString(str2)
	// buffer.WriteString("!!!!!!")
	// result4 := buffer.String()
	// fmt.Println(result4) // for really big lines


	// const (
	// 	execute = 0b001 // 1 - право выполнения
	// 	write   = 0b010 // 2 - право записи
	// 	read    = 0b100 // 3 - право чтения
	// )

	// executeRead := execute | read
	// fmt.Printf("execute (001) | read (100): число %d, бинарно %b\n", executeRead, executeRead)

	// writeRead := write | read
	// fmt.Printf("write (010) | read (100): число %d, бинарно %b\n", writeRead, writeRead)

	// permission := 5
	// fmt.Printf("Наши права: %b\n", permission)
	// fmt.Printf("Можем выполнять: %t\n", permission&execute > 0)
	// fmt.Printf("Можем записывать: %t\n", permission&write > 0)
	// fmt.Printf("Можем читать: %t\n", permission&read > 0)


	// num := 10
	// if num > 5 {
	// 	fmt.Println("a bigger than 5")
	// } else if a == 5{
	// 	fmt.Println("a equal 5")
	// } else {
	// 	fmt.Println("a not bigger than")
	// }

	// if num1 := rand.IntN(100); num1 > 50 {   // num1 доступно только в блоке if/else, not global variable 
	// 	fmt.Println("good")
	// } else {
	// 	fmt.Println("bad")
	// }

	// day := 6
	// switch day {
	// case 1: 
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wendsday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// 	fallthrough  // провалится в след кейс
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Unknown day") 		
	// }

	// switch {
	// case day < 1:
	// 	fmt.Println("Incorrect day")
	// case day >= 1 && day <= 5:
	// 	fmt.Println("Work day")
	// case day >= 6 && day <= 7:
	// 	fmt.Println("Weekends")
	// default:
	// 	fmt.Println("Unknown week")
	// }

	// const(
	// 	saturday = 6
	// 	sunday = 7
	// )

	// switch day {
	// case 1, 2, sunday - 4, 4, 5:
	// 	fmt.Println("Work day")
	// case saturday, sunday:
	// 	fmt.Println("Weekends")
	// default: 
	// 	fmt.Println("Unknown week")
	// }

	// var x any = 23
	// switch v := x.(type) {
	// case int:
	// 	fmt.Println("x is int:", v)
	// case string:
	// 	fmt.Println("x is string:", v)
	// default:
	// 	fmt.Println("Type of x unknown")
	// }

	// var val any = "dgfghdg"              // any == interface{}
	// switch v := val.(type) {
	// case int, string, bool, float64:
	// 	fmt.Printf("В переменной val находится тип: %T", v)
	// default:
	// 	fmt.Println("В переменной val находится неизвестный тип данных")
	// }


	/*
	Определение времени суток
	Вы были на встрече с начальником, где было озвучено задание, выполните его.

	Стенограмма разговора
	Начальник: Так, задача для тебя есть. Надо написать программу, в которую я буду вводить время, а она будет говорить, сейчас день, ночь или что...
	Вы: Окей, не проблема, вводить в консоли будем в каком формате?
	Начальник: Да, просто число вводится, 14 например, значит что 14 часов.
	Вы: Хорошо, утро, день, вечер, ночь, во сколько начинается?
	Начальник: Когда спокойной ночи малыши заканчиваются, тогда и ночь наступает, ах-хах - шучу :)) 
	Ладно, давай лучше сделаем так:
	Утро — это с 6 до 12
	День — с 12 до 18
	Вечер — с 18 до 23
	Ночь — с 23 до 6
	Вы: Понял. А что, если пользователь введет что-то не то?
	Начальник: Пусть тогда идет нахрен! Напиши сообщение: "Неверно задано время".
	Вы: Так, а если все правильно введено, то просто нужно вывести, например "Утро" или "День", да?
	Начальник: Не душни, выведи красиво, например "Сейчас 22ч. - вечер".
	Давай работай и постарайся без косяков сегодня.
	Вы: Какие косяки, начальник, все сделаем в лучшем виде! :)
	*/

	// var hours int
	// fmt.Print("Enter the time: ")
	// fmt.Scan(&hours)

	// if hours >= 6 && hours < 12 {
	// 	fmt.Printf("Now %dh. - Morning\n", hours)
	// } else if hours >= 12 && hours < 18 {
	// 	fmt.Printf("Now %dh. - Afternoon\n", hours)
	// } else if hours >= 18 && hours < 23 {
	// 	fmt.Printf("Now %dh. - Evening\n", hours)
	// } else if (hours >= 23 && hours <= 24) || (hours < 6 && hours >= 1) {
	// 	fmt.Printf("Now %dh. - Night\n", hours)
	// } else {
	// 	fmt.Println("Time is incorrect")
	// }


	// switch {
	// case hours >= 6 && hours < 12:
	// 	fmt.Printf("Now %dh. - Morning\n", hours)
	// case hours >= 12 && hours < 18:
	// 	fmt.Printf("Now %dh. - Afternoon\n", hours)
	// case hours >= 18 && hours < 23:
	// 	fmt.Printf("Now %dh. - Evening\n", hours)
	// case (hours >= 23 && hours <= 24) || (hours < 6 && hours >= 1):
	// 	fmt.Printf("Now %dh. - Night\n", hours)
	// default:
	// 	fmt.Println("Time is incorrect")
	// }


	// // Best case
	// var input int
	// fmt.Println("Enter the time in hours (0-23): ")
	// _, err := fmt.Scan(&input)

	// if err != nil || input < 0 || input > 23 {
	// 	fmt.Println("Incorrect time")
	// 	os.Exit(1)
	// }

	// var timeOfDay string
	// switch {
	// case input >= 6 && input < 12:
	// 	timeOfDay = "Morning"
	// case input >= 12 && input < 18:
	// 	timeOfDay = "Afternoon"
	// case input >= 18 && input < 23:
	// 	timeOfDay = "Evening"
	// default:
	// 	timeOfDay = "Night"
	// }

	// fmt.Printf("Now %02dh. - %s\n", input, timeOfDay)



	// /*
	// Ваша задача — создать программу, которая будет вычислять индекс массы тела (ИМТ) пользователя на основе его веса и роста. 
	// Программа должна запрашивать у пользователя ввод данных, производить необходимые вычисления и выводить результат с соответствующей категорией.

	// Формула для расчета ИМТ:

	// ИМТ = вес(кг) / рост(м2)
	// ​
	// На основе полученного значения ИМТ, программа должна определить и вывести соответствующую категорию:

	// Недостаточный вес: ИМТ < 18.5
	// Нормальный вес: 18.5 ≤ ИМТ < 25
	// Избыточный вес: 25 ≤ ИМТ < 30
	// Ожирение: ИМТ ≥ 30

	// Пример работы программы:
	// Введите ваш вес (кг): 70
	// Введите ваш рост (см): 175
	// Ваш ИМТ: 22.86
	// Категория: Нормальный вес
	// */

	// var weight int
	// fmt.Println("Введите ваш вес (кг): ")
	// _, err1 := fmt.Scan(&weight)
	
	// if err1 != nil || weight <= 0 {
	// 	fmt.Println("Incorrect weight")
	// 	os.Exit(1)
	// }

	// var height int
	// fmt.Println("Введите ваш рост (см): ")
	// _, err2 := fmt.Scan(&height)

	// if err2 != nil || height <= 0 {
	// 	fmt.Println("Incorrect height")
	// 	os.Exit(1)
	// }

	// heightM1 := float64(height) / 100
	// heightM2 := heightM1 * heightM1
	// IMB := float64(weight) / heightM2

	// fmt.Printf("Ваш ИМТ: %.2f\n", IMB)

	// if IMB < 18.5 {
	// 	fmt.Println("Категория: Недостаточный вес")
	// } else if IMB < 25 {
	// 	fmt.Println("Категория: Нормальный вес")
	// } else if IMB < 30 {
	// 	fmt.Println("Категория: Избыточный вес")
	// } else {
	// 	fmt.Println("Категория: Ожирение")
	// }


	// // best case
	// var weight1 float64
	// fmt.Printf("Введите ваш вес (кг): ")
	// if _, err3 := fmt.Scan(&weight1); err3 != nil || weight1 <= 0 {
	// 	log.Fatalf("Incorrect input of weight: %s\n", err)
	// }

	// var heightCm float64
	// fmt.Printf("Input your height (cm): ")
	// if _, err4 := fmt.Scan(&heightCm); err4 != nil || heightCm <= 0 {
	// 	log.Fatalf("Incorrect input of height: %s\n", err4)
	// }

	// BMI := weight1 / math.Pow(heightCm/100, 2)

	// var category string
	// switch {
	// case BMI < 18.5:
	// 	category = "Недостаточный вес"
	// case BMI < 25:
	// 	category = "Нормальный вес"
	// case BMI < 30:
	// 	category = "Избыточный вес"
	// default:
	// 	category = "Ожирение"
	// }

	// fmt.Printf("Your BMI: %.2f\n", BMI)
	// fmt.Printf("Category is: %s\n", category)



	/*
	Напишите программу, которая будет выполнять поиск товара по его названию. 
	Пользователь должен ввести название товара или его часть, 
	после чего программа должна вывести в консоль цену данного товара в формате:

	Название товара: цена товара
					
	Доступные товары и их цены:
	"Клавиатура JZ9": 19200
	"Наушники N45": 9600
	"Смартфон S10": 55000
	
	Если пользователь введет название товара, которого нет в списке, программа должна вывести сообщение:
	Товар "название_введенного_товара" не найден.
		
	Поиск должен быть нечувствителен к регистру. 
	Например, если пользователь введет "наУШНИКИ n45", программа должна вернуть цену 9600.

	Примеры:
	Введите название товара: наУШНИКИ n45
	Наушники N45: 9600
				
	Введите название тоСмартфон S10": 55000вара: сМАРТ
	Смартфон S10: 55000
					
	Введите название товара: планшет
	Товар "планшет" не найден.
	*/

	scanner := bufio.NewScanner(os.Stdin)   // Создаёт сканнер для чтения того что мы передадим
	fmt.Printf("Enter the product name: ")  // передаем
	scanner.Scan()							// читает, ждет ввод (когда нажмешь enter), и сохраняет переданное в сканер
	if err := scanner.Err(); err != nil {
		log.Fatalf("Input error: %s", err)
	}

	input := strings.ToLower(scanner.Text())  // возвращает введеную строку
	
	if strings.Contains("клавиатура jz9", input) {
		fmt.Println("Клавиатура JZ9: 19200")
	} else if strings.Contains("наушники n45", input) {
		fmt.Println("Наушники N45: 9600")
	} else if strings.Contains("смартфон s10", input) {
		fmt.Println("Смартфон S10: 55000")
	} else {
		fmt.Printf("The product %q didn't find.\n", input)  // %q - вывод в кавычках
	}

	switch {
	case strings.Contains("клавиатура jz9", input):
		fmt.Println("Клавиатура JZ9: 19200")
	case strings.Contains("наушники n45", input):
		fmt.Println("Наушники N45: 9600")
	case strings.Contains("смартфон s10", input):
		fmt.Println("Смартфон S10: 55000")
	default:
		fmt.Printf("The product %q didn't find.\n", input)
	}
	



	
	


}
