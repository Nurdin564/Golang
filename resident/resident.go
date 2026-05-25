package resident

import "fmt"

/*
Resident - структура, которая будет представлять жителя деревни. Она должна содержать следующие поля:

Name (имя)
Age (возраст)
Married (статус брака)
Alive (жив ли)
Events (список событий за год жизни)
Реализуйте методы для добавления года, изменения статуса брака, умирания, а также методы интерфейса VillageElement.

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

type Resident struct {
	Name string
	Age int
	Married bool
	Alive bool
	Events []string
}

func (r *Resident) AddYear(year int) string {
	return fmt.Sprintf("(возраст: %d),", year+1)
}

func (r *Resident) ChangeMarried(marry bool) string {
	if marry {
		return "статус: холост."
	} else {
		return "статус: в браке."
	}
}

func (r *Resident) ChangeAlive(alive bool) error {
	if alive {
		alive = false
	}
	return nil
}

func (r *Resident) FlushInfo() string {
	r.Events = []string{}
	return fmt.Sprintf("Житель %s %d %b\nСобытия:Нет\n", r.Name, r.Age, r.Married)
}