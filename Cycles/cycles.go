package main

import (
	"bufio"
	"errors"
	"fmt"

	// "log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
	// "os"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// a := 0
	// for a < 5 {
	// 	fmt.Println(a)
	// 	a++
	// }
	// fmt.Println(a)    // a = 5

	// for i := range 10 {
	// 	fmt.Println(i + 1)
	// }

	// for range 6 {
	// 	fmt.Println("Sup")
	// }

	// rollDice(8)

	// str := "Ку-ку!"
	// for i, r := range str {
	// 	fmt.Println(i, r)
	// }

	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(i, str[i])
	// }

	// fmt.Println()

	// for i, b := range []byte(str) {
	// 	fmt.Println(i, b)
	// }

	// a1 := "Кукушка"
	// PrintReplaced(a1)

	// count := 0
	// for {
	// 	fmt.Println("Counter: ", count)
	// 	count++
	// 	if count >= 3 {
	// 		break
	// 	}
	// }

	// for i := 0; i < 4; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println("Value of i: ", i)
	// }


	// for range 1 {
	// 	random = rand.Intn(100) + 1
	// 	guesses = 0
		
	// 	result := play()
	// 	fmt.Printf("Загадано: %d, Угадано: %d\n", random, result)
		
	// 	if result != random {
	// 		fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
	// 		os.Exit(-1)
	// 	}	
	// }

	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= 2; j++ {
	// 		fmt.Printf("i: %d, j: %d\n", i, j)
	// 	}
	// }

	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		if i == 1 && j == 2 {
	// 			break
	// 		}
	// 		fmt.Printf("i: %d, j: %d\n", i, j)
	// 	}
	// }

	// printTable(5)

	// printDiamond(2)

	// for {
	// 	if err := playGame(); err != nil {
	// 		if err == ErrUserEndGame {
	// 			printEndGameMassage()
	// 			break
	// 		}
	// 		fmt.Printf("An error occured during the game: %v", err)
	// 		os.Exit(1)
	// 	}

	// 	var playAgain string
	// 	fmt.Printf("Want you play again? (if yes, write word \"yes\"): ")
	// 	fmt.Scanln(&playAgain)
	// 	if playAgain != "yes" {
	// 		printEndGameMassage()
	// 		break
	// 	}
	// }

	input, err := GetInput()
	if err != nil {
		fmt.Printf("Error from getting text: %v\n", err)
		os.Exit(1)
	}

	DisplayResults(CountCharacters(input))
// 	letters, digits, spaces, punctuation := CountCharacters(input)
// 	DisplayResults(letters, digits, spaces, punctuation)
}

/*
Анализ строки
	Ваша задача — разработать программу, которая будет анализировать текст, введенный пользователем. 
	Программа должна выполнять несколько ключевых функций, позволяющих получить различные статистические данные о тексте.
Ввод текста
	Создайте функцию GetInput() (string, error), которая запрашивает у пользователя ввод текста и возвращает его для дальнейшего анализа. 
	Если функция вернет ошибку, программу необходимо завершить.
Подсчет символов
	Реализуйте функцию CountCharacters(text string) (letters, digits, spaces, punctuation int), которая принимает текст и 
	возвращает количество:
	букв (letters)
	цифр (digits)
	пробелов (spaces)
	знаков препинания (punctuation)
Вывод результатов
	Напишите функцию DisplayResults(letters, digits, spaces, punctuation int), которая принимает результаты анализа и 
	выводит их на экран в удобочитаемом формате. Например:
Количество букв: 50
Количество цифр: 10
Количество пробелов: 5
Количество знаков препинания: 3
*/

func GetInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input some text: ")
	if !scanner.Scan() {
		fmt.Println()
		if err := scanner.Err(); err != nil {
			return  "", fmt.Errorf("scan error: %w", err)
		}
		return "", errors.New("unable to read input")
	}

	return scanner.Text(), nil
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	for _, char := range text {
		switch {
		case unicode.IsLetter(char):
			letters++
		case unicode.IsDigit(char):
			digits++
		case unicode.IsSpace(char):
			spaces++
		case unicode.IsPunct(char):
			punctuation++
		}
	}
	return
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Printf("Count of letters: %d\n", letters)
	fmt.Printf("Count of digits: %d\n", digits)
	fmt.Printf("Count of spaces: %d\n", spaces)
	fmt.Printf("Count of punctuations: %d\n", punctuation)
}



/*
Игра "Угадай число"
Предлагаю создать очень простую версию игры, где программа загадывает случайное число в диапазоне от 1 до 100 включительно, 
а пользователь должен попытаться его угадать. Программа должна предоставлять подсказки о том, 
больше или меньше загаданное число по сравнению с предположением пользователя, а также учитывать количество попыток.
Детали
Генерация случайного числа:
	Необходимо создать функцию, которая будет генерировать случайное число от min до max включительно. 
	В данной игре min будет равен 1, а max — 100.
Ввод пользователя:
	Необходимо создать функцию, которая будет запрашивать у пользователя ввод числа. 
	Важно учесть, что человек может устать играть и захочет выйти из игры, поэтому нужно продумать вариант выхода. 
	Например, пользователь может ввести специальное слово выход, чтобы завершить игру. 
	Также необходимо учитывать, что при вводе могут возникать ошибки (например, ввод нечисловых значений). 
	Эта функция должна вызываться каждый раз, когда требуется получить значение от пользователя. 
Подсказки:
	После каждого ввода программа должна сообщать пользователю, 
	является ли загаданное число больше или меньше введенного значения.
Отслеживание попыток:
	Программа должна отслеживать и выводить количество удачных попыток ввода, 
	которые потребовались пользователю для угадывания числа.
Завершение игры:
	Игра должна завершаться, когда пользователь угадает число, с выводом сообщения о том, 
	что число угадано, и с указанием количества попыток.
Повторная игра:
	Необходимо также реализовать возможность повторной игры после завершения текущей, 
	если пользователь не вышел из игры намеренно.
Пример
	Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!
	Ваше предположение (либо, для завершения, введите слово выход): 50
	Загаданное число больше.
	Ваше предположение (либо, для завершения, введите слово выход): 75
	Загаданное число меньше.
	Ваше предположение (либо, для завершения, введите слово выход): 63
	Загаданное число меньше.
	Ваше предположение (либо, для завершения, введите слово выход): 56
	Загаданное число меньше.
	Ваше предположение (либо, для завершения, введите слово выход): 53
	Правильно! Вы угадали число с 5 попытки.
	Хотите сыграть еще раз? (если хотите, напишите слово да): нет
	Спасибо за игру! До свидания!
	*/

var ErrUserEndGame = errors.New("User finished game")

func printEndGameMassage() {
	fmt.Println("Thanks for play! Bye!")
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getUserInput() (int, error) {
	var input string
	fmt.Printf("Your assumption (or write \"exit\" to finish): ")
	fmt.Scanln(&input)

	if strings.ToLower(input) == "exit" {
		return 0, ErrUserEndGame
	}
	
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}
	return number, nil
	

}

func playGame() error {
	min := 1
	max := 100
	randomNumber := generateRandomNumber(min, max)
	attempts := 0

	fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

	for {
		input, err := getUserInput()
		if err != nil {
			if err == ErrUserEndGame {
				return err
			}
			fmt.Println("Incorrect input. Please, write integer.")
			continue
		}
		attempts++

		if input < randomNumber {
			fmt.Println("The hidden number is greater.")
		} else if input > randomNumber {
			fmt.Println("The hidden number is less.")
		} else {
			fmt.Printf("Right! You guessed a number since %d attempt.\n", attempts)
			break
		}
	}
	return nil
}






/*
Ромб из звездочек
Порисуем? Хочется увидеть ромбик в терминале, для этого необходимо создать функцию func printDiamond(n int), 
которая выводит текст Мой бриллиант:, а с новой строки ромб (смотри примеры), где n - это ребро ромба.
*/
func printDiamond(n int) {
	fmt.Println("Мой бриллиант:")

	for i := 0; i < n; i++ {
		for s := 0; s < n-i-1; s++ {
			fmt.Print(" ")
		}

		fmt.Print("#")

		if i > 0 {
			for s := 0; s < 2*i-1; s++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
		}

		fmt.Println()
	}

	for i := n - 2; i >= 0; i-- {
		for s := 0; s < n-i-1; s++ {
			fmt.Print(" ")
		}

		fmt.Print("#")

		if i > 0 {
			for s := 0; s < 2*i-1; s++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
		}

		fmt.Println()
	}
}



/*
Игра "Угадай число"
Ваша задача — реализовать игру, в которой необходимо будет создать функцию play() int, 
эта функция будет пытаться угадать загаданное число. 
	Загаданное число находится в диапазоне от 1 до 100 включительно.
Условия задания
	Для вас уже создана функция func guess(num int) (int, error), данная функция нужна для проверки загаданного числа.
	Функция guess возвращает один из следующих вариантов:
		-1, nil, если загаданное число меньше переданного.
		1, nil, если загаданное число больше переданного.
		0, nil, если число угадано.
		0, too many attempts, если функция была вызвана больше 6 раз, что будет считаться за невыполненное задание.
Функция play, которую вам необходимо реализовать, должна угадать загаданное число, используя функцию guess для проверки своих предположений.  
Когда функция play угадает число, она должна вернуть это значение.
Примечания
У вашей функции play, которую вам необходимо реализовать, есть всего 6 попыток, чтобы отгадать число, 
так как функция guess начнет возвращать ошибку после 6 попытки. За 6 вызовов можно со 100% вероятностью отгадать загаданное число. 
Такая же тактика может использоваться для поиска значений в отсортированном ряде данных, например в базах данных.
*/
var guesses int
var random int

func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attemts")
	}
	guesses++
	if num > random {       // if mid > random
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func play() int {
	left := 1
	right := 100

	for guesses < 6 {
		mid := (left + right) / 2
        fmt.Printf("Попытка %d: mid=%d (диапазон [%d, %d])\n", guesses+1, mid, left, right)

		res, err := guess(mid)
		if err != nil {
			fmt.Println("Ошибка:", err)
			break
		}

		if res == 0 {
			fmt.Println("Угадали!")
			return mid
		}

		if res == -1 {
			// загаданное число меньше
			right = mid - 1
		} else {
			// загаданное число больше
			left = mid + 1
		}
	}

	// Если по какой-то причине 6 попыток закончились,
	// остаётся единственный возможный вариант
	return left
}


/*
Необходимо симулировать броски двух шестигранных кубиков, для этого необходимо реализовать функцию rollDice. 
Данная функция должна принимать одно целое число, ничего при этом не возвращая.
	В функцию будет передаваться число в диапазоне от 2 до 12 включительно (число всегда передается в правильном диапазоне, проверять не нужно), 
	которое будет представлять целевую сумму, которую необходимо получить при бросках кубиков.
Требования
	Функция rollDice должна генерировать случайные броски двух кубиков до тех пор, пока сумма значений, 
	выпавших на кубиках, не станет равной переданному в функцию числу.
	Результат каждого броска необходимо выводить в консоль.
	Когда целевая сумма будет достигнута, программа должна вывести иное сообщение, добавив информацию об общем количестве бросков, 
	которые были сделаны для получения данной суммы.
Пример работы программы
	Если в функцию было передано число 8, то вывод может выглядеть следующим образом:
	Выпало 4 и 4, в сумме 8, на это потребовался 1 бросок.

Выпало 2 и 3, в сумме 5, бросаем еще раз.
Выпало 1 и 6, в сумме 7, бросаем еще раз.
Выпало 5 и 3, в сумме 8, на это потребовалось 3 броска.
              
Выпало 2 и 3, в сумме 5, бросаем еще раз.
Выпало 1 и 6, в сумме 7, бросаем еще раз.
Выпало 5 и 5, в сумме 10, бросаем еще раз.
Выпало 5 и 6, в сумме 11, бросаем еще раз.
Выпало 5 и 3, в сумме 8, на это потребовалось 5 бросков.
*/
func getRollWord(n int) string {
	// исключения 11-14
	if n%100 >= 11 && n%100 <= 14 {
		return "бросков"
	}

	switch n % 10 {
	case 1:
		return "бросок"
	case 2, 3, 4:
		return "броска"
	default:
		return "бросков"
	}
}

func rollDice(num int) {
	i := 0
	result := 0
	var first, second int

	for result != num {
		first = rand.Intn(6) + 1
		second = rand.Intn(6) + 1
		result = first + second
		i++

		if result != num {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз\n", first, second, result)
		}		
	}
	if i == 1 {
		fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовался %d бросок\n", first, second, result, i)
	} else {
		fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d %s\n", first, second, result, i, getRollWord(i))			
	}
}


/*
Замена символов в строке
Часто бывает, что нужно заменить какой-то символ в строке. 
	Напишите функцию PrintReplaced, которая принимает строку и выводит в консоль новую строку, 
	в которой все буквы "у" заменены на буквы "а" (регистрозависимо, прописные буквы игнорируем). 
Так как мы изучаем циклы, хотелось бы, чтобы и решение было с использованием циклов, а не готовых функций
*/
func PrintReplaced(str string) {
	result := ""
	for _, i := range(str) {
		if i == 'у' {
			result += "а"
		} else {
			result += string(i)
		}
	}
	fmt.Println(result)
}

func PrintReplaced1(s string) {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] == 'у' {
			runes[i] = 'а'
		}
	}
	fmt.Println(string(runes))
}


/*
Таблица умножения
	Хочется вывести таблицу умножения для ребенка. Необходимо написать функцию printTable, 
	которая принимает число (назовем параметр num) и выводит таблицу умножения num x num.
Используйте вложенные циклы и форматируйте вывод в виде таблицы, 
между примерами отступ должен состоять из одного знака табуляции, в конце каждой строки табуляции быть не должно.
	Sample Output 1: (3)
1 x 1 = 1	1 x 2 = 2	1 x 3 = 3
2 x 1 = 2	2 x 2 = 4	2 x 3 = 6
3 x 1 = 3	3 x 2 = 6	3 x 3 = 9
*/
func printTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%d x %d = %d", i, j, i*j)

			if j < num {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}
}





