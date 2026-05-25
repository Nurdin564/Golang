package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	// "log"
	"math"
	// "yourgo/repository"
	// "yourgo/worker"
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


	// user := User{
	// 	ID: 2,
	// 	Firstname: "Pavel",
	// }
	// product := Product{
	// 	ID: 5,
	// 	Title: "Pencil",
	// }
	// fmt.Println(user)  // {UserID-2 Pavel}
	// fmt.Println(product)  // {ProductID-5 Pencil}


	// // repo := repository.NewBookInMemoryRepository()
	// repo1 := repository.NewSecondMemoryRepository()
	
	// wrkr, err := worker.New(repo1)  // dependency injection, передаем что-то воркеру(структуре)
	// if err != nil {
	// 	log.Fatalf("new worker errorL %v", err)
	// }
	// wrkr.CreateBooks()
	// wrkr.PrintBook(5)


	// myDB := MyDB{}
	// myLog := Logger{}
	// run(myDB, myLog)



	village := Village{}

	// Создаем жителей деревни
	resident1 := &Resident{Name: "Алиса", Age: 30, Married: false, Alive: true, Events: []string{}}
	resident2 := &Resident{Name: "Борис", Age: 40, Married: true, Alive: true, Events: []string{}}

	// Создаем животных
	animal1 := &Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	animal2 := &Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	village.AddElement(animal1)
	village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}	

}

/*
Village - структура, которая будет содержать список элементов деревни (жителей и животных). 
Реализуйте для данной структуры методы добавления элементов, обновления всех элементов и вывода информации о всех элементах.

Ваша задача — реализовать симуляцию, в которой будут участвовать два типа элементов: Жители и Животные. 
Каждый элемент должен иметь возможность обновлять свое состояние 
(например, стареть, умирать, заводить семью и т.д.) и выводить информацию о себе.

Составные части программы:
VillageElement - интерфейс, который будет содержать два метода:

Update(): обновляет состояние элемента (добавляет год жизни). 
В этом методе с определенной вероятностью должны происходить различные события. 
События могут как менять свойства объекта (например, смерть или вступление в брак), 
так и быть просто текстовыми (например, "Устроился на работу" или "Покусал прохожего"). 
Все произошедшие за год события должны сохраняться в список Events.

FlushInfo() string: возвращает строку с информацией об элементе и очищает все события (обнуляет срез Events).
*/

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Village struct {
	Elements []VillageElement
}

func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}

func (v *Village) UpdateAll() {
	for _, e := range v.Elements {
		e.Update()
	}
}

func (v Village) ShowAllInfo() string {
	info := ""
	for _, e := range v.Elements {
		info += e.FlushInfo()
	}
	return info
}


/*
Resident - структура, которая будет представлять жителя деревни. Она должна содержать следующие поля:

Name (имя)
Age (возраст)
Married (статус брака)
Alive (жив ли)
Events (список событий за год жизни)
Реализуйте методы для добавления года, изменения статуса брака, умирания, а также методы интерфейса VillageElement.

Реализуйте вспомогательные методы для добавления года, изменения статуса брака, умирания, а также методы интерфейса VillageElement. 
Внутри метода Update вызывайте эти вспомогательные методы, а также добавляйте произвольные текстовые события в список Events 
с любой придуманной вами вероятностью.
*/

type Resident struct {
	Name string
	Age int
	Married bool
	Alive bool
	Events []string
}

func (r *Resident) addYear() {
	r.Age++
}

func (r *Resident) changeMarriedStatus() {
	if r.Married {
		r.Married = false
		r.Events = append(r.Events, "Развод, больше не в браке")
	} else {
		r.Married = true
		r.Events = append(r.Events, "Наконец-то, найден спутник в жизни!!!")
	}
} 

func (r *Resident) die() {
	r.Alive = false
	r.Events = append(r.Events, fmt.Sprintf("Ушел на покой на %d году жизни.", r.Age))
}

func (r *Resident) Update() {
	if !r.Alive {
		return
	}
	r.addYear()
	if rand.IntN(100) < 15 {
		r.changeMarriedStatus()
	}
	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Нашел новую работу.")
	}
	if r.Married && rand.IntN(100) < 25 {
		r.Events = append(r.Events, "Поругался с супругой/ом.")
	}
	if rand.IntN(100) < 5 {
		r.die()
	}
}

func (r *Resident) FlushInfo() string {
	info := fmt.Sprintf("Житель %s умер в возрасте %d.\n", r.Name, r.Age)
	if r.Alive {
		marriedStatus := "холост"
		if r.Married {
			marriedStatus = "в браке"
		}
		events := "нет"
		if len(r.Events) > 0 {
			events = strings.Join(r.Events, "\n")
		}
		info = fmt.Sprintf("Житель %s (возраст: %d), статус: %s.\nСобытия: %s\n", r.Name, r.Age, marriedStatus, events)
	}
	r.Events = []string{}
	return info
}


/*
Animal - структура, которая будет представлять животное. Она должна содержать следующие поля:

Name (имя)
Age (возраст)
Type (тип животного, например, "собака", "кошка")
Alive (жив ли)
Events (список событий за год жизни)
Реализуйте методы для добавления года, умирания, а также методы интерфейса VillageElement.

Реализуйте вспомогательные методы для добавления года, умирания, а также методы интерфейса VillageElement. 
Как и у жителей, в методе Update животного должны случайно происходить события, влияющие на состояние (смерть) 
и не влияющие (например, "Сломал лапу" или "Покусал прохожего").
*/

type Animal struct {
	Name string
	Age int
	Type string
	Alive bool
	Events []string
}

func (a *Animal) addYear() {
	a.Age++
}

func (a *Animal) die() {
	a.Alive = false
	a.Events = append(a.Events, fmt.Sprintf("Умер на %d году жизни.", a.Age))
}

func (a *Animal) Update() {
	if !a.Alive {
		return
	}
	a.addYear()
	if rand.IntN(100) < 7 {
		a.Events = append(a.Events, "Сломал лапу.")
	}
	if a.Type == "кошка" && rand.IntN(100) < 25 {
		a.Events = append(a.Events, "Убежала из дома.")
	}
	if a.Type == "собака" && rand.IntN(100) < 25 {
		a.Events = append(a.Events, "Покусал прохожего.")
	}
	if rand.IntN(100) < 7 {
		a.die()
	}
}

func (a *Animal) FlushInfo() string {
	info := fmt.Sprintf("Житель %s умер в возрасте %d.\n", a.Name, a.Age)
	if a.Alive {
		events := "нет"
		if len(a.Events) > 0 {
			events = strings.Join(a.Events, "\n")
		}
		info = fmt.Sprintf("Животное %s (возраст: %d),\nСобытия: %s\n", a.Name, a.Age, events)
	}
	a.Events = []string{}
	return info
}




type DB interface {  // Встраивание интерфейсов (дробление, для состава кастомных интерфесов. Чтобы использовать лишь то что нужно)
	Writer
	Read() string
}

type Writer interface {
	Write(string)
}

type Logger struct{}

func (l Logger) Write(str string) {
	fmt.Printf("Wrote this log: %q\n", str)
}


type MyDB struct{}

func (m MyDB) Write(str string) {
	fmt.Printf("Wrote this line: %q\n", str)
}

func (m MyDB) Read() string {
	return "Returned line that we need"
}


func run(db DB, logger Writer) {
	db.Write("Helloo")
	fmt.Println(db.Read())

	logger.Write("i have entered at 12:15")
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