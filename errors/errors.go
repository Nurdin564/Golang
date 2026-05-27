package main

import (
	// "bufio"
	"errors"
	"fmt"
	// "log"
	"os"
	"strconv"

	// "log"
	"math/rand/v2"
	// "yourgo/networkerr"
	"yourgo/worker1"
	// "yourgo/worker2"
)

func main() {
	/*
	Is проверяет, равна ли ошибка конкретному значению. Функция проходит по цепочке ошибок и сравнивает каждую из них 
	с заданным значением с помощью оператора ==. Используется для проверки на так называемые сигнальные ошибки (например, io.EOF, sql.ErrNoRows, os.ErrNotExist). 

	As проверяет, можно ли привести ошибку к определённому типу. Если тип подходит, функция сохраняет ошибку в переменную. 
	Используется, когда нужен доступ к данным или поведению, специфичным для этого типа ошибки. 

	Таким образом, Is применяют, когда нужно узнать, произошла ли определённая ситуация, 
	а As — когда необходим доступ к полям или методам ошибки определённого типа. 
	*/



	// result, err := calculate(6, 2)
	// if err != nil {
	// 	log.Fatalf("calculate: %v", err)
	// }
	// fmt.Println("Result of working:", result)


	// if err := someFunc(); err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }


	// if err := run(); err != nil {
	// 	var fileNotExists *worker1.FileNotExistsError
	// 	if errors.As(err, &fileNotExists) {  // если знаем с какой конкретно ошибкой (сами ее создали) сравнивать то errors.Is() / юзаем error.As() чтобы понять из какой структуры ошибка
	// 		// log.Fatalf("check if file exists, worker cannot find file.")
	// 		log.Printf("sending email to admin with text: worker cannot find path %s", fileNotExists.Path)
	// 	} else {
	// 		log.Fatalf("worker error: %s", err)
	// 	}
	// }


	// svc := Service{db: &MockDB{}}

	// fmt.Println("=== Тест 1: успешный чат ===")
	// svc.PrintChat(1)

	// fmt.Println("\n=== Тест 2: DatabaseError код 24 ===")
	// svc.PrintChat(2)

	// fmt.Println("\n=== Тест 3: DatabaseError код 148 (id=264651) ===")
	// svc.PrintChat(264651)

	// fmt.Println("\n=== Тест 4: неизвестная ошибка ===")
	// svc.PrintChat(999)


	// w := worker2.New()
	// if err := w.DoWork(); err != nil {
	// 	var networkerr *networkerr.NetworkError
	// 	switch{
	// 	case errors.Is(err, worker2.ErrNoInternet):  // Unwrap() error - распаковка ошибки, добираемся до дочерней ошибки и сраниваем
	// 		log.Fatalf("check internet")
	// 	case errors.As(err, &networkerr):
	// 		log.Fatalf("network worker error with code %d: %s", networkerr.Code, networkerr)
	// 	default:
	// 		log.Fatalf("unknown worker error: %s", err)
	// 	}
	// }


	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Write the line: ")
	// if !scanner.Scan() {
	// 	if err := scanner.Err(); err != nil {  // don't save the error inside struct
	// 		log.Fatalf("Error when we read insert: %s", err)
	// 	} else {
	// 		log.Fatalf("Can't scan the line.")
	// 	}
	// }
	// fmt.Println("You wrote:", scanner.Text())


	if err := HandlePayments(); err != nil {
		var insFund *InsufficientFundsError
		if errors.Is(err, ErrUnsupportedPaymentMethod) {
			fmt.Println("Ошибка: неподдерживаемый метод платежа.")
		} else if errors.As(err, &insFund) {
			fmt.Printf("Ошибка: недостаточно средств. %v", insFund)
		} else {
			fmt.Printf("Неизвестная ошибка: %v", err)
		}
		os.Exit(1)
	}

	fmt.Println("Платеж успешно обработан!")

}

/*
Вам предоставлен код простой платежной системы. Ваша задача — улучшить обработку ошибок в функции main, 
чтобы сделать вывод сообщений об ошибках более информативным и понятным для пользователя.

Первым делом, внимательно прочитайте предоставленный код и поймите, как работает система обработки платежей. 
Обратите внимание на типы ошибок, которые могут возникнуть.

В функции main добавьте обработку ошибок, чтобы различать разные типы ошибок, 
которые могут возникнуть при вызове функции HandlePayments. Используйте следующие условия:

- Если ошибка связана с неподдерживаемым методом платежа, выведите сообщение: "Ошибка: неподдерживаемый метод платежа."
- Если ошибка связана с недостатком средств, выведите сообщение: "Ошибка: недостаточно средств. [подробности об ошибке]."
- Для всех остальных ошибок выведите сообщение: "Неизвестная ошибка: [текст ошибки]."

Пример вывода
Если пользователь введет неподдерживаемый метод платежа:
Ошибка: неподдерживаемый метод платежа.
                
Если пользователь введет сумму, превышающую максимальную:
Ошибка: недостаточно средств. requested amount: 1500.00, maximum allowed amount: 1000.00
                 
Если произойдет другая ошибка:
Неизвестная ошибка: [текст ошибки].
*/

// Заранее подготавливаем некоторые ошибки
var (
	ErrUnsupportedPaymentMethod = errors.New("unsupported payment method")
)

// Ошибка нехватки средств
type InsufficientFundsError struct {
	RequestedAmount float64
	MaxAmount       float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("requested amount: %.2f, maximum allowed amount: %.2f", e.RequestedAmount, e.MaxAmount)
}

// PaymentProcessor обрабатывает платежи
type PaymentProcessor struct{}

// NewPaymentProcessor создает новый процессор платежей
func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

// ProcessPayment обрабатывает платеж
func (pp *PaymentProcessor) ProcessPayment(method string, amount float64) error {
	if method != "карта" && method != "СБП" {
		return ErrUnsupportedPaymentMethod
	}

	// Имитация проверки средств
	maxAmount := 1000.0
	if amount > maxAmount {
		return &InsufficientFundsError{
			RequestedAmount: amount,
			MaxAmount:       maxAmount,
		}
	}

	return nil
}

func HandlePayments() error {
	var method string
	var amountStr string

	// Спросить метод перевода в консоли
	fmt.Print("Введите метод перевода (карта/СБП): ")
	fmt.Scanln(&method)

	// Спросить сумму перевода в консоли
	fmt.Print("Введите сумму перевода: ")
	fmt.Scanln(&amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %w", err)
	}

	pp := NewPaymentProcessor()
	if err := pp.ProcessPayment(method, amount); err != nil {
		return fmt.Errorf("process payment: %w", err)
	}

	return nil
}



/*
Вам необходимо реализовать метод PrintChat в сервисе, который отвечает за вывод сообщений чата. 
Метод должен выполнять следующие действия:

- Вызвать метод GetChatByIDWithMessages у базы данных, передав ему идентификатор чата (ID).
- Если метод GetChatByIDWithMessages возвращает ошибку, необходимо обработать её в зависимости от типа:
- Если ошибка является DatabaseError и код ошибки равен 24, вывести в консоль сообщение: Ошибка запроса: <текст ошибки>.
- Если ошибка является DatabaseError, но код не равен 24, вывести: Инфраструктурная ошибка: <текст ошибки>.
- Если ошибка не является DatabaseError, вывести: Неизвестная ошибка: <текст ошибки>.
- Если ошибок нет, вывести все сообщения чата в формате: <отправитель>: <сообщение>.
Примечание
	Функция main уже реализована, она создает нужные структуры и вызывает метод PrintChat.
	После вывода каждой строки сообщения чата, необходимо добавить символ переноса строки.
*/

type Message struct {
	From    string
	Message string
}

type Chat struct {
	ID       int
	Messages []Message
}

type DatabaseError struct {
	Message string
	Code    int
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error with code %d: %s", e.Code, e.Message)
}


type Service struct {
	db interface {
		GetChatByIDWithMessages(id int) (*Chat, error)
	}
}

func (w Service) PrintChat(id int) {
    chat, err := w.db.GetChatByIDWithMessages(id)
    if err != nil {
        var dbErr *DatabaseError
        if errors.As(err, &dbErr) {
            if dbErr.Code == 24 {
                fmt.Printf("Ошибка запроса: %s\n", dbErr.Error())
            } else {
                fmt.Printf("Инфраструктурная ошибка: %s\n", dbErr.Error())
            }
        } else {
            fmt.Printf("Неизвестная ошибка: %s\n", err.Error())
        }
        return
    }

    for _, msg := range chat.Messages {
        fmt.Printf("%s: %s\n", msg.From, msg.Message)
    }
}

// Моковая база данных
type MockDB struct{}

func (m *MockDB) GetChatByIDWithMessages(id int) (*Chat, error) {
	switch id {
	case 1:
		// Успешный случай
		return &Chat{
			ID: 1,
			Messages: []Message{
				{From: "Иван", Message: "Привет!"},
				{From: "Мария", Message: "Как дела?"},
			},
		}, nil
	case 2:
		// DatabaseError с кодом 24
		return nil, &DatabaseError{Code: 24, Message: "query timeout"}
	case 264651:
		// DatabaseError с другим кодом
		return nil, &DatabaseError{Code: 148, Message: "connection lost"}
	default:
		// Неизвестная ошибка
		return nil, errors.New("unexpected error")
	}
}



func run() error {
	w := worker1.New()
	if err := w.DoWork("file.exe"); err != nil {
		return fmt.Errorf("do work: %w", err)  // %w save stack trace of functions errors
	}
	return nil
}


type MyError struct {
	Code int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func someFunc() error {
	return &MyError{
		Code: 500,
		Message: "internal error",
	}
}


func calculate(num1, num2 int) (int, error) {
	if rand.IntN(100) > 50 {
		return 0, errors.New("God says there will be an error here")
	}

	result, err := divide(num1, num2)
	if err != nil {
		return 0, fmt.Errorf("divide: %w", err)
	}

	return result, nil
} 

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}


/*
Вам необходимо реализовать функцию UpdateUserName, которая будет обновлять имя пользователя в системе. 
Для этого вам нужно использовать уже существующие функции GetUserByID и SaveUser.

Структура User уже определена и содержит необходимые поля.

Задача
Реализуйте функцию UpdateUserName(id, name string) error, которая принимает два параметра:

id (тип string): уникальный идентификатор пользователя.
name (тип string): новое имя пользователя.
Функция UpdateUserName должна выполнить следующие действия:

Вызовите функцию GetUserByID(id string) (*User, error) для получения пользователя по его идентификатору.
Если функция GetUserByID вернёт ошибку, оберните её и верните из функции.
Если пользователь найден, обновите его поле Name новым значением.
Вызовите функцию SaveUser(user *User) error для сохранения обновлённого пользователя.
Если функция SaveUser вернёт ошибку, также оберните её и верните.

type User struct {
	ID   string
	Name string
}

func UpdateUserName(id, name string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return fmt.Errorf("can't get user by id: %w", err)
	}
	user.Name = name

	err := SaveUser(user)
	if err != nil {
		return fmt.Errorf("can't save user: %w", err)
	}

	return nil
}
*/

