package main

import (
	"errors"
	"fmt"
	"log"

	// "log"
	"math"
	"math/rand"
	"strings"
	"unicode/utf8"
)

func main(){
	// fmt.Println("Finally, here we go again")

	// petBattle(34, 50)

    // printNumberInfo(25)

    // generateCompliment("Nurdin")

    // helloStr := "Sup, dawg!"
    // _, runesLength := getFullLength(helloStr)
    // fmt.Printf("Line %q has %d symbols\n", helloStr, runesLength)

    // isSecured := securePassword("asaddffwf11111")
    // fmt.Println(isSecured)

    // sum(2, 5, 7, 45, 12)

    // val := 5
    // modifyValue(val)
    // fmt.Println(val)
    // modifyPointer(&val)
    // fmt.Println(val)

    // divide(10, 2)

    // fn()
    // fn1()

    // defer handlePanic()
    // riskyFunction()
    // fmt.Println("This line never be reached, if panic happened")

    // result, err := divide1(10, 2)
    // if err != nil {
    //     log.Fatalf("divide error: %s", err)
    // }
    // fmt.Println("Result of dividing:", result)

    // info, err := userProfileToString("", 94)
    // if err != nil {
    //     fmt.Println("Получена ошибка от функции userProfileToString:", err)
    // } else {
    //     fmt.Println(info)
    // }

    // data, err := calculate1(100.88, 45.77, "multiply")
    // if err != nil {
    //     fmt.Println("Получена ошибка от функции calculate1:", err)
    // } else {
    //     fmt.Println("result of operation:", data)
    // }

    a, b := 10.0, 0.0
    result, err := calculate2(a, b)
    if err != nil {
        log.Fatalf("unable to calculate %s\n", err)
    }
    fmt.Println("result is:", result)
}


func petBattle(cats, dogs int){
    if cats > dogs {
        fmt.Printf("Котики победили со счетом %d:%d!\n", cats, dogs)
    } else if dogs > cats {
        fmt.Printf("Собачки победили со счетом %d:%d!\n", dogs, cats)
    } else {
        fmt.Println("Ничья! Все дружат!")
    }
}

func printNumberInfo(num int){
    if num > 0 {
        fmt.Printf("Number %d is positive\n", num)
    } else if num < 0 {
        fmt.Printf("Number %d is negative\n", num)
    } else {
        fmt.Println("Number is 0")
    }

    if num%2==0 {
        fmt.Printf("Number %d is even\n", num)
    } else {
        fmt.Printf("Number %d is odd\n", num)
    }

    if num > 0 {
        root := math.Sqrt(float64(num))
        if root == float64(int(root)) {
            fmt.Printf("Square root of %d is integer and equal %.0f.\n", num, root)
        } else {
            fmt.Printf("Square root of %d is not integer and equal %.5f.\n", num, root)
        }
    }
    // if num > 0 {
    //     root := math.Sqrt(float64(num))
    //     intRoot := int(root)
    //     if intRoot*intRoot == num {
    //         fmt.Printf("Square root of %d is integer and equal %d.\n", num, intRoot)
    //     } else {
    //         fmt.Printf("Square root of %d is not integer and equal %.5f.\n", num, root)
    //     }
    // }
}

func generateCompliment(name string) string {
    num := rand.Intn(3)
    switch num {
    case 0: 
        return fmt.Sprintf("You are great, %s!\n", name)
    case 1:
        return fmt.Sprintf("You have an amazing smile, %s!\n", name)
    default:
        return fmt.Sprintf("You are inspiring, %s!\n", name)
    }
}

func getFullLength(str string) (bytes int, runes int) {
    bytes = len(str)
    runes = utf8.RuneCountInString(str)
    return  // вернет bytes and runes автоматически
}

// func securePassword(pass string) bool {
//     return utf8.RuneCountInString(pass) >= 6 &&
//         pass != strings.ToLower(pass) &&
//         pass != strings.ToUpper(pass)
// }

func securePassword(pass string) bool {
    if utf8.RuneCountInString(pass) < 6 {
        return false
    }
    if pass == strings.ToLower(pass) {
        return false
    }
    if pass == strings.ToUpper(pass) {
        return false
    }
    return true
}

func sum(a, b int, values ...int) {   // values = [...]
    fmt.Println(a, b, values)
}

func modifyValue(v int) {
    v = v + 10
}

func modifyPointer(v *int) {
    *v = *v + 10
}

func divide(n1, n2 int) {
    defer fmt.Println("End of function divide 1")  // откладывает выполнение функции на конец, выполняются в обратном порядке
    defer fmt.Println("End of function divide 2")
    fmt.Println(n1 / n2)

}

func fn() {
    i := 0
    defer func(val int) {
        fmt.Println(val)
    }(i)
    i = 5
}
func fn1() {
    i := 0
    defer func() {
        fmt.Println(i)
    }()
    i = 12
}

func handlePanic() {
    if r := recover(); r != nil {               // if <инициализация>; <условие> {
        fmt.Println("panic is handled:", r)
    }
}

func riskyFunction() {
    panic("something goes wrong")
}

func someFunction() error {
    return fmt.Errorf("make %s: error in number %d", "someFunction", 29)
    // return errors.New(fmt.Sprintf("make %s: error in number %d", "someFunction", 29))
}

func divide1(n1, n2 int) (int, error) {
    if n2 == 0 {
        return 0, errors.New("divide by 0")
    }
    return n1 / n2, nil
}


/* 
Напишите функцию UserProfileToString, которая принимает имя пользователя name (string) и его возраст age (int).
+ Функция должна возвращать строку с сообщением (string) и ошибку (error), если таковая была.
+ Сообщение должно быть в формате: Имя человека: [ИМЯ], возраст: [ВОЗРАСТ].
+ Если имя было передано в функцию и имеет пробелы слева/справа - эти пробелы нужно убрать.
Возможные ошибки:
+ Если имя пользователя пустое, функция должна вернуть ошибку с сообщением: empty name
+ Если возраст меньше нуля, необходимо вернуть ошибку с сообщением: negative age
+ Если передано имя, состоящее только из пробелов, то функция должна вернуть ошибку с сообщением: name cannot contain only spaces 
*/
func userProfileToString(name string, age int) (string, error) {
    trimmed := strings.TrimSpace(name)

    if name == "" {
        return "", fmt.Errorf("empty name")
    }
    if age < 0 {
        return "", fmt.Errorf("negative age")
    }
    if trimmed == "" {
        return "", fmt.Errorf("name cannot contain only spaces")
    }

    return fmt.Sprintf("Human name is: %s, age is %d.\n", trimmed, age), nil
}


/*
Создадим стандартный калькулятор. Реализуйте функцию calculate, которая принимает три аргумента:
    Число (тип float64) - первый аргумент для вычисления.
    Число (тип float64) - второй аргумент для вычисления.
    Строку, представляющую операцию, которую необходимо выполнить. Возможные операции:
    "add" (сложение)
    "subtract" (вычитание)
    "multiply" (умножение)
    "divide" (деление)
    Функция должна возвращать результат операции и ошибку.
Ошибки
    Если операция не поддерживается, функция должна возвращать ошибку с сообщением: unknown operation.
    Если происходит деление на ноль, возвращаем ошибку с сообщением division by zero.
    В случае любой ошибки, первое возвращаемое значение должно быть равно 0.
*/
func calculate1(n1, n2 float64, operation string) (float64, error) {
    switch operation {
    case "add":
        return n1 + n2, nil
    case "subtract":
        return n1 - n2, nil
    case "multiply":
        return n1 * n2, nil
    case "divide":
        if n2 == 0 {
            return 0, errors.New("division by zero")
        }
        return n1 / n2, nil
    default:
        return 0, errors.New("unknown operation")
    }
}



func divide2(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("unable divide %f to zero", a)
    }
    return a / b, nil
}

func calculate2(a, b float64) (float64, error) {
    if err := logicX(); err != nil {
        return 0, fmt.Errorf("logicX: %w", err)
    }

    result, err := divide2(a, b)
    if err != nil {
        return 0, fmt.Errorf("divide: %w", err)
    }

    return result, nil
}

func logicX() error {
    if err := logicY(); err != nil {
        return fmt.Errorf("logicY: %w", err)  // оборачиваем ошибку для отслеживания
    }
    return nil
}

func logicY() error {
    return errors.New("failed by reason C")  // create error
    // return nil
}

// func userProfile(user string)