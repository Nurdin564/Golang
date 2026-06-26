package main

import (
	"fmt"
	"time"

	
)

func main() {
	coal := 0

	transferPoint := make(chan int, 3)  // каналы синхронизируют отдельные потоки выполнения, горутины. Чтобы они обменивались информацией

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)	
	coal += <-transferPoint

	time.Sleep(3 * time.Second)	
	coal += <-transferPoint

	fmt.Println("Добыли", coal, "угля!")  // Добыли 30 угля!
	fmt.Println("Прошло времени:", time.Since(initTime))  // Прошло времени: 9.0015771s. Это время работы главного потока (main goroutine), он принимал по одной горутине в течении 3 сек, а так горутины закончили за секунду свою работу и продолжили свою логику



	coal1 := 0

	initTime1 := time.Now()

	coal1 += mine1(1)
	coal1 += mine1(2)
	coal1 += mine1(3)

	fmt.Println("Добыли", coal1, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime1))
}

func mine1(n int) int {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")	

	return 10
}




func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	transferPoint <- 10
	fmt.Println("Поход номер", n, "уголь передал!")

	// ... в буферизованном канале, горутина кладет 10 и дальше идет выполнять логику или завершается
}
