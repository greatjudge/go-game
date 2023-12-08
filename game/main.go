package main

import (
	"strings"
)

// import "fmt"

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/

var gm GameMap
var room int
var player Player

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все комнаты
		если что-то было - оно корректно перезатирается
	*/
	room = 0
	gm = InitGameMap()
	player = NewPlayer([]Mission{
		{Name: "PackBag", Description: "собрать рюкзак"},
		{Name: "GoUniversity", Description: "идти в универ"},
	})
}

func handleCommand(command string) string {
	// $команда $параметр1 $параметр2 $параметр3
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	comm := strings.Split(command, " ")
	var mes string
	switch comm[0] {
	case "осмотреться":
		mes = LookAround(room, gm, player)
	case "идти":
		room, mes = Move(room, comm[1], gm, player)
	case "надеть":
		mes = PutOn(room, comm[1], &gm, &player)
	case "взять":
		mes = Take(room, comm[1], &gm, &player)
	case "применить":
		mes = Apply(room, comm[1], comm[2], &gm, &player)
	default:
		mes = "неизвестная команда"
	}
	return mes
}
