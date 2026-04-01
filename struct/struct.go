package main

import (
	"fmt"
	"strings"
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



}

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
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

func printInfo(user User) {
	fmt.Printf("Customer %s. Phone: %s. Address: t. %s, %s.\n", user.Name, user.Phone, user.Address.City, user.Address.Street)
	
	isElectronics := false
	for _, item := range user.Cart{
		if item.Product.Category == "Electronics" {
			isElectronics = true
			break			
		}
	}
	if isElectronics {
		fmt.Println("User is a customer of electronics.")
	} else {
		fmt.Println("User isn't a customer of electroniks")
	}

	slice := []string{}
	for _, item := range user.Cart {
		if item.Product.Price > 1000 {
			slice = append(slice, item.Product.Name)
		}
	}
	if len(slice) > 0 {
		fmt.Printf("Items in the basket, with a price of 10,000 or more: %s.\n", strings.Join(slice, ", "))
	} else {
		fmt.Println("Items in the basket with a price of 10,000 or more: none.")
	}

	sum := 0
	for _, item := range user.Cart {
		sum += item.Quantity * item.Product.Price
	}
	fmt.Printf("Total purchase amount: %d\n", sum)
}




// type User struct {
// 	Name Name
// 	BirthYear int
// }

// type Name struct {
// 	First string
// 	Last string
// }

// type User1 struct {
// 	Name struct {
// 		First string
// 		Last string
// 	}
// 	BirthYear int
// }