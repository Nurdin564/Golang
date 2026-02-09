package main

import (
	// "bufio"
	"fmt"
	"math"
	"math/rand"
	// "os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	a = 5
	b = 2
	c int
)

const (
	s = 1
	t = 5.5
	v
)

func main() {
	
	var age1 int = 30
	var age2 = 20
	age3 := 40
	fmt.Println(age1, age2, age3)

	var x, y int = 1, 2
	x1, y2 := 4, 5
	fmt.Println(x, y, x1, y2)

	fmt.Println(a, b ,c)

	h := 35
	fmt.Println(h)
	h = 49
	fmt.Println(h)

	const MY_VAR = 3.1415
	fmt.Println(MY_VAR)

	fmt.Println(s, t ,v)

	randomNum0 := rand.Intn(500) // [0, 500)
	fmt.Println(randomNum0)

	min := 10
	max := 50
	randomNum := rand.Intn(max-min) + min // [0, 40) + 10 == [10, 50)
	fmt.Println(randomNum)

	randomFloat := rand.Float64() // [0, 1)
	fmt.Println(randomFloat)

	min1 := 20.0
	max1 := 50.0
	randomFloat1 := rand.Float64()*(max1-min1+1) + min1 // [0,31) + 20 == [20, 51) == [20, 50]
	fmt.Println(randomFloat1)


	fmt.Println(math.Pow(2, 3))   // Returns 2^3 = 8.0
	fmt.Println(math.E)
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Log(5))
	fmt.Println(math.Log10(5))
	minValue := math.Min(3.0, 6.0)
	maxValue := math.Max(3.0, 6.0)
	fmt.Println("minValue:", minValue, "maxValue:", maxValue)

	fmt.Println(math.Round(3.4))   // 3
	fmt.Println(math.Round(-3.5))  // -4
	fmt.Println(math.Floor(3.4))  // округляет вниз
	fmt.Println(math.Ceil(3.4))   // округляет вверх
	fmt.Println(math.Trunc(3.4))  // отбрасывает дробную часть 
	
	number := 3.1415926535
	rounded := math.Round(number*100) / 100  // если нужно округлить до 2 знаков после запятой
	fmt.Println(rounded)


	str := "Hello"
	fmt.Println(len(str))  // 5 bytes == print 5
	str1 := "Привет"
	fmt.Println(len(str1))  // 12 bytes == print 12
	fmt.Println("Runes count: ", utf8.RuneCountInString(str1))  // print count of symbols == 6

	str2 := " Hello my \"nigga\" "  // экранирование print Hello my "nigga"
	fmt.Println(str2)


	str3 := "Bitch ganster shit"
	words := strings.Split(str3, " ")
	fmt.Println(words)

	fmt.Println(strings.Contains(str3, "shit"))  // true (регистрозависимая)


	i := 555
	s1 := strconv.Itoa(i)  // print string i (decimal)
	fmt.Println(s1)

	s2 := strconv.FormatInt(int64(i), 2)   // translate number to binary presentation (string)
	fmt.Println(s2)
	
	s3 := strconv.FormatInt(int64(i), 16)   // translate number to hexidecimal (string)
	fmt.Println(s3)
	
	number00 := 10.233744
	fixed := strconv.FormatFloat(number00, 'f', 3, 64)  // 3 - знаков после запято,также округляет; 64 - тип number00 /(string) 
	fmt.Println(fixed)

	str4 := "1234567"
	num, err := strconv.Atoi(str4)  // первевод строки в число
	// num, err := strconv.ParseInt(str4, 10, 0)  // 10-система счисления, 0 - int64
	if err != nil {
		fmt.Println("Ошибка преобразования: ", err)
	} else {
		fmt.Println("Преобразованное число: ", num)
	}

	str5 := "3.1415"
	num01, err01 := strconv.ParseFloat(str5, 64)  // 64 - то к какому типу переведется str5, даже если str5 - int, переведется к float64
	if err01 != nil {
		fmt.Println("Ошибка преобразования: ", err01)
	} else {
		fmt.Println("Преобразованное число: ", num01)
	}
	
	priceStr := "100.56"
	quantityStr := "67"
	// find "price*quntity" with using Atoi/ParseFloat/FormatFloat
	priceFloat, _ := strconv.ParseFloat(priceStr, 64)   // _ - when we want translate string to number without catching errors
	quantityFloat, _ := strconv.ParseFloat(quantityStr, 64)
	total := strconv.FormatFloat(priceFloat * quantityFloat, 'f', 2, 64)  // from float to string
	fmt.Println(total)  // print string(6737.52)


	name := "Samuel"
	age := 34
	fmt.Printf("My name is %s, I am %d years old. \n", name, age)     // f-строка
	
	str6 := fmt.Sprintf("My name is %s, I am %d years old.", name, age)  // не для выведения инфы а для сохранения отфарматированной инфы в переменную
	fmt.Println(str6)

	fmt.Printf("%f\n", 2345.66)  // 2345.66000000 - 6 знаков после запятой
	fmt.Printf("%.7f\n", 7584.83756) // 7 знаков после запятой
	fmt.Printf("%v %v\n", "fgfgfgfg", 978795.099)   // %v - универсальный


	//pointers
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)  // type *int, value nil

	// nil pointer panic
	// fmt.Printf("%T %#v %#v\n", intPointer, intPointer, *intPointer)

	x2 := 43
	fmt.Println("Значение х2: ", x2)  // 43
	fmt.Println("Указатель на х2: ", &x2)  // 0xc00000a1f8, & - указывает на адресс

	p := &x2  // var p (*int), указывает на тот же адресс что и x2
	fmt.Println("Значение p: ", p)  // 0xc00000a1f8, p хранит адресс x2
	fmt.Println("Значение по адресу p: ", *p)  // 43, *p (*int) указывает на значение по адресу (int)

	*p = 100
	fmt.Println("Значение по адресу p: ", *p)  // 100            // & - адрес, * - значение
	fmt.Println("Значение x2: ", x2)  // 100 

	p1 := new(int)  // создает указатель на тип int
	fmt.Println(p1)  // 0xc00000a210
	fmt.Println(*p1) // 0, not nil pointer panic
	*p = 51
	fmt.Println(*p)  // 51
	


	// var firstName, lastName string
	// fmt.Print("Enter your name and surname with space: ")
	// fmt.Scan(&firstName, &lastName)
	// fmt.Printf("Hello %s. Surname %s seems like familiar...\n", firstName, lastName)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Enter the line: ")
	// scanner.Scan()
	// fmt.Println("You entered:", scanner.Text())

	var firstName1, lastName1 string
	fmt.Print("Enter your name and surname with space: ")
	fmt.Scan(&firstName1, &lastName1)
	var old int
	fmt.Print("Enter your old: ", )
	fmt.Scan(&old)
	fmt.Printf("Nice to meet you, %s. 5 years ago, I have met with man, who's surname is %s too, you were %d. How we were yuang!\n", firstName1, lastName1, old - 5)





}

/*
многострочные комментарии
... 
*/
