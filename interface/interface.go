package main

import (
	"errors"
	"fmt"
	"math"

	// "golang.org/x/text/width"
)

type Dog struct{}

func (Dog) Speak() string {
	return "Gaf"
}

func (Dog) Growl() string {
	return "RRRRrrrRRR"
}

type Cat struct{}

func (Cat) Speak() string {
	return "MMyyaaauuu"
}

type Bird struct{}

func (Bird) Speak() string {
	return "KKarrrr"
}


type Speaker interface {
	Speak() string
}

func makeSound(s Speaker) {  // polimorphism, duck typing
	fmt.Println(s.Speak())
}

func main() {
	// Interface can be nil!


	// Dog := Dog{}
	// Cat := Cat{}
	// Bird := Bird{}

	// makeSound(Dog)
	// makeSound(Cat)
	// makeSound(Bird)


	// warrior, _ := NewWarrior("Артур")
	// mage, _ := NewMage("Мерлин")
	// archer, _ := NewArcher("Робин")

	// characters := []Character{
	// 	warrior,
	// 	mage,
	// 	archer,
	// }

	// Fight(characters)


	// rectangle := Rectangle{width: 10, height: 5}
	// circle := Circle{radius: 7}

	// printShapeInfo("rectangle", rectangle)
	// printShapeInfo("circle", circle)

	// s := []any{
	// 	Circle{radius: 4},
	// 	Rectangle{width: 12, height: 3},
	// 	"sgdfhdgss",
	// 	nil,
	// 	345,
	// }
	// printAreas(s)


	// var p *Person  // nil
	// fmt.Println(p.Struct()) // Person
	// p.SayHello() // never achieve


	// fmt.Println(South)  // 4 or South. Package fmt has interface Stringer and inside of it  String() string{} - same as our func. It's automatikly activate if it see the same method. So if you delete your method it will print 4 
	// fmt.Println(South.String())  // South
	// action(South)  // action South
	// actionStr(South)  // action South


	user := User{
		ID: 2,
		Firstname: "Pavel",
	}
	product := Product{
		ID: 5,
		Title: "Pencil",
	}
	fmt.Println(user)  // {UserID-2 Pavel}
	fmt.Println(product)  // {ProductID-5 Pencil}
}

type UserId int

func (u UserId) String() string {
	return fmt.Sprintf("UserID-%d", u)  // if i want to print ID with previous unique words
}

type User struct {
	ID UserId
	Firstname string
}


type ProductID int

func (p ProductID) String() string {
	return fmt.Sprintf("ProductID-%d", p)  // if i want to print ID with previous unique words
}

type Product struct {
	ID ProductID
	Title string
}



type Direction int

const(
	_ Direction = iota  // skip 0, start from 1
	North  // 2
	West   // 3
	South  // 4
	East   // 5
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return ""
	}
}

func action(d Direction) {
	fmt.Println("action", d)
}

func actionStr(d fmt.Stringer) {
	fmt.Println("action", d)
}


type Person struct {
	Firsname string
	Lastname string
	Age int
}

func (p *Person) SayHello() { 
	fmt.Printf("%s greets to you!\n", p.Firsname)  // doesn't work cause you need to dereference  (*p).Firstname  but you can't dereference the nil
}

func (*Person) Struct() string { // works when you have pointer and doesn't have object. * - allows you dereference
	return "Person"
}
 

func printAreas(s []any) {
	for _, v := range s {
		if t, ok := v.(Shaper); ok {
			fmt.Printf("Shape: Area=%.2f, Perimeter=%.2f\n", t.Area(), t.Perimeter())
		} else if t, ok := v.(string); ok {
			fmt.Printf("String: %s\n", t)
		} else {
			fmt.Println("Not a shape")
		}
	}
}


type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}


type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}


type Shaper interface {
	Area() float64
	Perimeter() float64
}

func printShapeInfo(title string, s Shaper) {
	fmt.Printf("Figure %q, area is %.2f, perimeter is %.2f\n", title, s.Area(), s.Perimeter())
}


/*
Необходимо создать игровую систему с тремя типами персонажей: Warrior, Mage и Archer. 
Каждый персонаж должен иметь метод Attack() string, который возвращает строку, описывающую действие персонажа:

Для Warrior метод Attack должен возвращать строку: "Воин [ИМЯ] бьет мечом."
Для Mage метод Attack должен возвращать строку: "Маг [ИМЯ] колдует огненный шар."
Для Archer метод Attack должен возвращать строку: "Лучник [ИМЯ] выпускает град стрел."
Для каждого типа персонажа необходимо создать функцию-конструктор, которая будет возвращать объект нужной структуры и ошибку, 
пример для воина: NewWarrior(name string) (*Warrior, error). 
Ошибка должна возвращаться с текстом empty name в случае, если переданное имя было пустым.

Необходимо реализовать функцию Fight, которая принимает слайс интерфейсов типа Character (этот тип вам нужно создать).
Функция должна вывести в консоль строку, которую возвращает метод Attack для каждого переданного в слайсе персонажа.

Примечание
Вам необходимо реализовать структуры, методы для структур , функцию Fight и интерфейс Character. 
Функция main уже реализована, она создает персонажей и вызывает функцию Fight.
Вывод должен быть на новой строке для каждого персонажа.
Функция main уже реализована, она вызывает функции NewWarrior, NewMage и NewArcher для создания персонажей 
и после передает слайс Character в функцию Fight.
*/

type Character interface {
	Attack() string
}

func Fight(chars []Character) {
	for _, c := range chars {
		fmt.Print(c.Attack())
	}
}


type Warrior struct {
	Name string
}

func NewWarrior(name string) (*Warrior, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Warrior{
		Name: name,
	}, nil
}

func (w *Warrior) Attack() string {
	return fmt.Sprintf("Воин %s бьет мечом.\n", w.Name)
}


type Mage struct{
	Name string
}

func NewMage(name string) (*Mage, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Mage{
		Name: name,
	}, nil
}

func (m *Mage) Attack() string {
	return fmt.Sprintf("Маг %s колдует огненный шар.\n", m.Name)
}


type Archer struct{
	Name string
}

func NewArcher(name string) (*Archer, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Archer{
		Name: name,
	}, nil
}

func (a *Archer) Attack() string {
	return fmt.Sprintf("Лучник %s выпускает град стрел.\n", a.Name)
}