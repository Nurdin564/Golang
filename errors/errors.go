package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
)

func main() {
	// result, err := calculate(6, 2)
	// if err != nil {
	// 	log.Fatalf("calculate: %v", err)
	// }
	// fmt.Println("Result of working:", result)


	if err := someFunc(); err != nil {
		log.Fatalf("Error: %s", err)
	}
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

