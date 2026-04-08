package main

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"strings"
	// "yourgo/model"
	"time"
	"unicode"
	"fmt"
	"math/rand"
)

func main() {
	// user := User{
	// 	FirstName: "Nigga",
	// 	LastName: "Shit",
	// 	BirthYear: 1984,
	// }

	// fmt.Println(user)  // {Nigga Shit 1984}
	// fmt.Printf("%+v\n", user)  // {FirstName:Nigga LastName:Shit BirthYear:1984}
	// fmt.Println(user.FirstName)  // Nigga
	// fmt.Println(user.BirthYear) // 1984

	// var user1 User
	// fmt.Printf("%+v\n", user1)  // {FirstName: LastName: BirthYear:0}
	// user2 := User{}
	// fmt.Printf("%+v\n", user2) // {FirstName: LastName: BirthYear:0}
	// user3 := &User{}
	// fmt.Printf("%+v\n", user3)  // &{FirstName: LastName: BirthYear:0}

	// name := Name{
	// 	First: "Giga",
	// 	Last: "Niga",
	// }

	// user := User {
	// 	Name: name,
	// 	BirthYear: 1990,
	// }
	// fmt.Printf("%+v\n", user)

	// user.Name = Name {
	// 	First: "Gangster",
	// 	Last: "is Gay",
	// }
	// fmt.Println(user.BirthYear)  // 1990
	// fmt.Println(user.Name.First)  // Gangster
	// fmt.Println(user.Name.Last)  // is Gay

	// user.Name.First = "Pididi"
	// fmt.Printf("%+v\n", user)  // {Name:{First:Pididi Last:is Gay} BirthYear:1990}

	// user1 := User1 {
	// 	Name: struct{
	// 		First string
	// 		Last string
	// 	}{
	// 		First: "Epstein",
	// 		Last: "Jeffry",
	// 	},
	// 	BirthYear: 1888,
	// }

	// user1.Name = struct{
	// 	First string
	// 	Last string
	// }{
	// 	First: "Arthur",
	// 	Last: "Celestial",
	// }
	// fmt.Printf("%+v\n", user1)  // {Name:{First:Arthur Last:Celestial} BirthYear:1888}

	// user := User{
	// 	ID:    1,
	// 	Name:  "Иван Петров",
	// 	Email: "ivan.petrov@example.com",
	// 	Phone: "+7 999 123-45-67",
	// 	Address: Address{
	// 		Street:     "Улица Ленина",
	// 		City:       "Москва",
	// 		PostalCode: "101000",
	// 	},
	// 	Cart: []CartItem{
	// 		{
	// 			Product: Product{
	// 				ID:          1,
	// 				Name:        "Ноутбук",
	// 				Description: "Мощный ноутбук для работы и игр",
	// 				Price:       59990,
	// 				Category:    "Электроника",
	// 				Brand:       "Brand A",
	// 				Rating:      4.5,
	// 				Reviews:     120,
	// 			},
	// 			Quantity: 1,
	// 		},
	// 		{
	// 			Product: Product{
	// 				ID:          2,
	// 				Name:        "Смартфон",
	// 				Description: "Современный смартфон с отличной камерой",
	// 				Price:       29990,
	// 				Category:    "Электроника",
	// 				Brand:       "Brand B",
	// 				Rating:      4.7,
	// 				Reviews:     200,
	// 			},
	// 			Quantity: 2,
	// 		},
	// 		{
	// 			Product: Product{
	// 				ID:          3,
	// 				Name:        "Наушники",
	// 				Description: "Беспроводные наушники с шумоподавлением",
	// 				Price:       7990,
	// 				Category:    "Аудио",
	// 				Brand:       "Brand C",
	// 				Rating:      4.3,
	// 				Reviews:     80,
	// 			},
	// 			Quantity: 1,
	// 		},
	// 	},
	// }

	// printInfo(user)

	// user1 := User{
	// 	Name: Name{
	// 		First: "Pavel",
	// 		Last: "Tarasov",
	// 	},
	// 	BirthYear: 1919,
	// }
	// user1.Greet()

	// newName := Name{
	// 	First: "Ivan",
	// 	Last: "Ivanov",
	// }
	// user1.ChangeName(newName)
	// user1.Greet()


	// Создаём пользователя
	// user := User{
	// 	FirstName: "Алексей",
	// 	LastName:  "Смирнов",
	// 	BirthYear: 1987,
	// }

	// // Добавляем любимые языки
	// _ = user.AddFavoriteLanguage("Go")
	// _ = user.AddFavoriteLanguage("Python")

	// // Пытаемся добавить дубликат
	// err := user.AddFavoriteLanguage("go")
	// if err != nil {
	// 	fmt.Println("Ошибка:", err)
	// }

	// // Генерируем секретное имя
	// fmt.Println("Секретное имя:", user.SecretIdentity())

	// // Проверяем возраст
	// fmt.Println("Возраст:", user.Age())

	// // Проверяем любимый язык
	// fmt.Println("Любимый язык Go?", user.IsProgrammingLanguageFavorite("Go"))
	// fmt.Println("Любимый язык Java?", user.IsProgrammingLanguageFavorite("Java"))

	// // Случайный язык
	// lang, _ := user.RandomFavoriteLanguage()
	// fmt.Println("Случайный любимый язык:", lang)

	// // Печатаем профиль
	// fmt.Println("\nПрофиль пользователя:")
	// fmt.Println(user.GenerateProfile())

	// // Обновляем имя
	// err = user.UpdateName("Иван", "Петров")
	// if err != nil {
	// 	fmt.Println("Ошибка обновления имени:", err)
	// }

	// // Удаляем язык
	// err = user.RemoveFavoriteLanguage("Python")
	// if err != nil {
	// 	fmt.Println("Ошибка удаления:", err)
	// }

	// // Итоговый профиль
	// fmt.Println("\nОбновлённый профиль:")
	// fmt.Println(user.GenerateProfile())


	// student, err := NewStudent("Nurdin", 20, 5, "nnnn11@1")
	// if err != nil {
	// 	fmt.Printf("Incorrect data: %v\n", err)
	// 	return  // если ошибка есть → мы сразу выходим из main
	// }
	// fmt.Printf("New student has created %+v\n", *student)


	// // p1 := model.Person{FirstName: "Nurdin", Age: 35}  // {FirstName:Nurdin LastName: Age:35} lastname is empty
	// p1, err := model.NewPerson("Nurdin", "Isamailov", 20)
	// if err != nil {
	// 	log.Fatalf("unable to create person: %v", err)
	// }
	// fmt.Printf("%+v\n", p1)
	// // p2 := model.Person{FirstName: "Pavel", Age: 35}  // можем не заметить если структура поменялась
	// p2, err := model.NewPerson("Pavel", "", 29)
	// if err != nil {
	// 	log.Fatalf("unable to create person: %v", err)
	// }
	// fmt.Printf("%+v\n", p2)

	
	car := Car{
		Engine: Engine{
			Started: false,
			HorsePower: 150,
		},
		Model: "Toyota",
	}

	if err := car.Drive(); err != nil {
		log.Fatalf("drive car: %v", err)
	}

	if err := car.Start(); err != nil {
		log.Fatalf("start car: %v", err)
	}
	fmt.Printf("Car is started, model %s, horsepower %d\n", car.Model, car.Engine.HorsePower)
	fmt.Printf("%+v\n", car)

}

type Engine struct {
	Started bool
	HorsePower int
}

func (e *Engine) Start() error {  // dont forget pointer when you change value in the struct
	if e.Started {
		return errors.New("already started")
	}
	e.Started = true
	return nil
}

type Car struct {
	Engine Engine
	Model string

}

func (c *Car) Start() error {  // Работаем не с копией. Если внутри метода меняется хотя бы одно поле (даже вложенное) → используй pointer receiver
	if err := c.Engine.Start(); err != nil {
		return fmt.Errorf("engine start: %w\n", err)
	}
	// стартуем другие сервисы
	return nil
}

func (c Car) Drive() error {
	if !c.Engine.Started {
		return errors.New("engine not started")
	}
	// ...
	return nil
}

/*
Создайте структуру Student, которая будет представлять студента с полями Name, Age, Grade и Email. 
Реализуйте функцию-конструктор NewStudent, которая будет проверять следующие условия:
	Поле Name не должно быть пустым.
	Поле Age должно быть в определенном диапазоне.
	Поле Grade должно быть в определенном диапазоне.
	Если Age больше константы GreatAge, то Grade должен быть не ниже MinGradeAfterGreatAge.
	Поле Email должно содержать символ @.
	Если какое-либо условие не выполняется - возвращаем подходящую ошибку, указатель на структуру при этом, должен быть nil.
*/
const (
	MinAge                = 15
	MaxAge                = 80
	MinGrade              = 1
	MaxGrade              = 5
	GreatAge              = 30
	MinGradeAfterGreatAge = 3
)

var (
	ErrEmptyName         = errors.New("name cannot be empty")
	ErrTooYoung          = errors.New("too young")
	ErrTooOld            = errors.New("too old")
	ErrGradeOutOfRange   = errors.New("grade out of range")
	ErrTooLowGradeForAge = errors.New("too low grade for age")
	ErrIncorrectEmail    = errors.New("incorrect email")
)

type Student struct {
	Name  string
	Age   int
	Grade int
	Email string
}

func NewStudent(name string, age, grade int, email string) (*Student, error) {
	if name == "" {return nil, ErrEmptyName}
	if age < MinAge {return nil, ErrTooYoung}
	if age > MaxAge {return nil, ErrTooOld}
	if grade < MinGrade || grade > MaxGrade {return nil, ErrGradeOutOfRange}
	if age > GreatAge && grade < MinGradeAfterGreatAge {return nil, ErrTooLowGradeForAge}
	if !strings.Contains(email, "@") {return nil, ErrIncorrectEmail}

	student := &Student{
		Name: name,
		Age: age,
		Grade: grade,
		Email: email,
	}
	return student, nil
}



/*
Необходимо реализовать структуру User с определенными методами.
Структура

	FirstName (string) - имя
	LastName (string) - фамилия
	BirthYear (int) - год рождения
	FavoriteLanguages ([]string) - любимые языки программирования

Методы

	SecretIdentity() string: генерирует секретное имя.
	Секретное имя должно быть составлено из первых букв имени и фамилии, за которыми следует случайное число от 1 до 100.
	Например, для пользователя с именем "Алексей" и фамилией "Смирнов" секретное имя может быть "АС42".

	Age() int: возвращает текущий возраст пользователя на основе года рождения
	(текущий год можно получить с помощью time.Now().Year(), мы это пока не проходили).

	AddFavoriteLanguage(language string) error: добавляет язык программирования в слайс FavoriteLanguages,
	если добавляемый язык уже присутствует в слайсе, то необходимо вернуть ошибку с сообщением duplicate.
	Если было передано пустое имя, нужно вернуть ошибку empty language name.

	RemoveFavoriteLanguage(language string) error: удаляет язык программирования из слайса FavoriteLanguages,
	если языка нет, возвращает ошибку not found.

	IsProgrammingLanguageFavorite(language string) bool: проверяет, является ли указанный язык программирования любимым.

	RandomFavoriteLanguage() (string, error): возвращает случайный язык из списка любимых языков программирования,
	если любимых языков нет в слайсе FavoriteLanguages, необходимо вернуть ошибку no options.

	GenerateProfile() string: возвращает строку с полным профилем пользователя в верном формате (пример):
	Имя: Павел.
	Фамилия: Тарасов.
	Возраст: 35.
	Список любимых языков программирования: [Язык1, Язык2].

	UpdateName(firstName, lastName string) error: обновляет имя и фамилию пользователя,
	если были переданы пустые имя или фамилия, необходимо вернуть ошибку empty data,
	также необходимо вернуть ошибку invalid data, если имя или фамилия написаны с маленькой буквы.
*/
type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User) SecretIdentity() string {
	random := rand.Intn(100) + 1
	secret := u.FirstName[:1] + u.LastName[:1] + strconv.Itoa(random)
	return secret
}

func (u User) Age() int {
	currentDate := time.Now().Year()
	age := currentDate - u.BirthYear
	return age
}

func (u *User) AddFavoriteLanguage(language string) error {
	if language == "" {
		return errors.New("empty language name")
	}
	if slices.Contains(u.FavoriteLanguages, language) {
		return errors.New("duplicate")
	}	

	u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	return nil
}

func (u *User) RemoveFavoriteLanguage(language string) error {
	for i, v := range u.FavoriteLanguages {
		if v == language {
			u.FavoriteLanguages = append(u.FavoriteLanguages[:i], u.FavoriteLanguages[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (u User) IsProgrammingLanguageFavorite(language string) bool {
	return slices.Contains(u.FavoriteLanguages, language)
}

func (u User) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) == 0 {
		return "", errors.New("no options")
	}
	randomIndex := rand.Intn(len(u.FavoriteLanguages))
	return u.FavoriteLanguages[randomIndex], nil
}

func (u User) GenerateProfile() string {
		return fmt.Sprintf(
		"Name: %s.\nSurname: %s.\nAge: %d.\nList of favorite programming languages: %v.",
		u.FirstName,
		u.LastName,
		u.Age(),
		u.FavoriteLanguages,
	)
}

func (u *User) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("empty data")
	}

	if !unicode.IsUpper(rune(firstName[0])) || !unicode.IsUpper(rune(lastName[0])) {
		return errors.New("invalid data")
	}

	u.FirstName = firstName
	u.LastName = lastName
	return nil
}

/*
Необходимо разработать структуру данных студента и реализовать методы для работы с данной структурой.
Структура должна содержать информацию об имени студента и его оценках,
а также предоставлять функциональность для вычисления средней оценки и вывода информации о студенте.

Структура Student
	Name (строка) - имя студента.
	Grades (срез целых чисел) - список оценок студента.
Методы
	AverageGrade() float64: метод, который вычисляет и возвращает среднюю оценку студента,
	с округлением по правилам математики до одного числа в дробной части.
	Средняя оценка рассчитывается как сумма всех оценок, деленная на количество оценок.
	Если у студента нет оценок, метод должен возвращать 0.
Info(): метод, который возвращает информацию о студенте в следующем формате (без переноса строки):
	Студент [ИМЯ], средняя оценка: [ОЦЕНКА].
*/
// type Student struct {
// 	Name string
// 	Grades []int
// }

// func (s Student) AverageGrade() float64 {
// 	if len(s.Grades) == 0 {
// 		return 0
// 	}

// 	sum := 0
// 	for _, v := range s.Grades {
// 		sum += v
// 	}

// 	return float64(sum / len(s.Grades))
// }

// func (s Student) Info() {
// 	fmt.Printf("Student %s, average grade: %.1f", s.Name, s.AverageGrade())
// }

// func (u User) Greet() {
// 	fmt.Printf("I am %s %s, %d years of born\n", u.Name.First, u.Name.Last, u.BirthYear)
// }

// func (u *User) ChangeName(name Name) {
// 	u.Name = name
// }

/*
Функция printInfo должна выводить следующую информацию, каждое сообщение на новой строке:
Информация о пользователе.
Формат данных:
	Покупатель [ИМЯ]. Телефон: [ТЕЛЕФОН]. Адрес: г. [ГОРОД], [УЛИЦА].
Покупатель электроники.
	Вывести, есть ли в корзине пользователя что-нибудь из категории "Электроника" в формате:
	Пользователь [является/не является] покупателем электроники.
Товары с высокой ценой.
	Вывести названия товаров в корзине, цена которых выше или равна 10000 в формате:
	Товары в корзине, где цена 10000 и более: [НАЗВАНИЯ_ТОВАРОВ].
	Если подходящих товаров не окажется, то вместо названия товаров необходимо вывести:
	Товары в корзине, где цена 10000 и более: отсутствуют.
Общая сумма покупки.
	Вывести общую сумму товаров в корзине, которая рассчитывается как сумма Price * Quantity для каждого элемента в Cart в формате:
	Общая сумма покупки: [СУММА] руб.
*/
// User представляет пользователя
// type User struct {
// 	ID      int
// 	Name    string
// 	Email   string
// 	Phone   string
// 	Address Address
// 	Cart    []CartItem
// }

// // Address представляет адрес пользователя
// type Address struct {
// 	Street     string
// 	City       string
// 	PostalCode string
// }

// // CartItem представляет элемент в корзине
// type CartItem struct {
// 	Product  Product
// 	Quantity int
// }

// // Product представляет продукт в корзине
// type Product struct {
// 	ID          int
// 	Name        string
// 	Description string
// 	Price       int
// 	Category    string
// 	Brand       string
// 	Rating      float64
// 	Reviews     int
// }

// func printInfo(user User) {
// 	fmt.Printf("Customer %s. Phone: %s. Address: t. %s, %s.\n", user.Name, user.Phone, user.Address.City, user.Address.Street)

// 	isElectronics := false
// 	for _, item := range user.Cart{
// 		if item.Product.Category == "Electronics" {
// 			isElectronics = true
// 			break
// 		}
// 	}
// 	if isElectronics {
// 		fmt.Println("User is a customer of electronics.")
// 	} else {
// 		fmt.Println("User isn't a customer of electroniks")
// 	}

// 	slice := []string{}
// 	for _, item := range user.Cart {
// 		if item.Product.Price > 1000 {
// 			slice = append(slice, item.Product.Name)
// 		}
// 	}
// 	if len(slice) > 0 {
// 		fmt.Printf("Items in the basket, with a price of 10,000 or more: %s.\n", strings.Join(slice, ", "))
// 	} else {
// 		fmt.Println("Items in the basket with a price of 10,000 or more: none.")
// 	}

// 	sum := 0
// 	for _, item := range user.Cart {
// 		sum += item.Quantity * item.Product.Price
// 	}
// 	fmt.Printf("Total purchase amount: %d\n", sum)
// }

// type User struct {
// 	Name Name
// 	BirthYear int
// }

type Name struct {
	First string
	Last  string
}

// type User1 struct {
// 	Name struct {
// 		First string
// 		Last string
// 	}
// 	BirthYear int
// }
