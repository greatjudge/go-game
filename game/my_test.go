package main

import (
	"testing"
)

func TestLookAround(t *testing.T) {
	initGame()

	mes := LookAround(0, gm, player)
	answer := "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	if answer != mes {
		t.Error("cmd:", "LookAround",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer)
	}

	mes = LookAround(3, gm, player)
	answer = "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"
	if answer != mes {
		t.Error("cmd:", "LookAround",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer)
	}
}

func TestMove(t *testing.T) {
	initGame()

	room := 0
	room, mes := Move(room, "коридор", gm, player)
	answer := "ничего интересного. можно пройти - кухня, комната, улица"
	if answer != mes || room != 1 {
		t.Error("cmd:", "Move",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\tnew room: ", room)
	}

	room, mes = Move(1, "кухня", gm, player)
	answer = "кухня, ничего интересного. можно пройти - коридор"
	if answer != mes || room != 0 {
		t.Error("cmd:", "Move",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\tnew room: ", room)
	}

	room, mes = Move(1, "комната", gm, player)
	answer = "ты в своей комнате. можно пройти - коридор"
	if answer != mes || room != 3 {
		t.Error("cmd:", "Move",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\tnew room: ", room)
	}

	room, mes = Move(1, "улица", gm, player)
	answer = "дверь закрыта"
	if answer != mes || room != 1 {
		t.Error("cmd:", "Move",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\tnew room: ", room)
	}

	room, mes = Move(0, "комната", gm, player)
	answer = "нет пути в комната"
	if answer != mes || room != 0 {
		t.Error("cmd:", "Move",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\tnew room: ", room)
	}
}

func TestPutOn(t *testing.T) {
	initGame()

	mes := PutOn(3, "рюкзак", &gm, &player)
	answer := "вы надели: рюкзак"

	_, ok := gm.Rooms[3].Items["рюкзак"]
	if answer != mes || !player.hasBag || ok {
		t.Error("cmd:", "PutOn",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\thasBag: ", player.hasBag,
			"\n\titem in room: ", ok)
	}
}

func TestTake(t *testing.T) {
	initGame()

	mes := Take(3, "ключи", &gm, &player)
	answer := "некуда класть"
	_, roomOk := gm.Rooms[3].Items["ключи"]
	_, plOk := player.Items["ключи"]

	if answer != mes || !roomOk || plOk {
		t.Error("cmd:", "Take",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\titem in room: ", roomOk,
			"\n\titem in player: ", plOk)
	}

	PutOn(3, "рюкзак", &gm, &player)

	mes = Take(3, "ключи", &gm, &player)
	answer = "предмет добавлен в инвентарь: ключи"
	_, roomOk = gm.Rooms[3].Items["ключи"]
	_, plOk = player.Items["ключи"]

	if answer != mes || roomOk || !plOk {
		t.Error("cmd:", "Take",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\titem in room: ", roomOk,
			"\n\titem in player: ", plOk)
	}

	mes = Take(3, "ключи", &gm, &player)
	answer = "нет такого"
	if answer != mes {
		t.Error("cmd:", "Take",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer)
	}

	mes = Take(3, "телефон", &gm, &player)
	answer = "нет такого"
	if answer != mes {
		t.Error("cmd:", "Take",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer)
	}
}

func TestApply(t *testing.T) {
	initGame()

	PutOn(3, "рюкзак", &gm, &player)
	Take(3, "ключи", &gm, &player)

	mes := Apply(1, "ключи", "дверь", &gm, &player)
	answer := "дверь открыта"
	if mes != answer || gm.Rooms[1].IsClosed(2) || gm.Rooms[2].IsClosed(1) {
		t.Error("cmd:", "Apply",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer,
			"\n\t closed 1 2: ", gm.Rooms[1].IsClosed(2),
			"\n\t closed 2 1: ", gm.Rooms[2].IsClosed(1),
		)
	}

	mes = Apply(1, "телефон", "шкаф", &gm, &player)
	answer = "нет предмета в инвентаре - телефон"
	if mes != answer {
		t.Error("cmd:", "Apply",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer)
	}

	mes = Apply(1, "ключи", "шкаф", &gm, &player)
	answer = "не к чему применить"
	if mes != answer {
		t.Error("cmd:", "Apply",
			"\n\tresult:  ", mes,
			"\n\texpected:", answer)
	}
}
