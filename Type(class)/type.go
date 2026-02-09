package main

import(
	"fmt"
)

type Direction int

const(
	_ Direction = iota
	North
	West
	South
	East
)

func main(){
	action(North)
	action(4)

}

func action(d Direction) {
	fmt.Println("Действие в направлении:", d)
}