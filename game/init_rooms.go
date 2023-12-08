package main

import "strings"

func CanGoMes(room int, gm GameMap) string {
	NextRooms := gm.Rooms[room].NextRoomsMessage(gm)
	if NextRooms != "" {
		return "можно пройти - " + NextRooms
	}
	return ""
}

func MissionMes(p Player) string {
	descriptions := make([]string, 0, len(p.MissionNames))
	for _, name := range p.MissionNames {
		mission, ok := p.Name2Mission[name]
		if ok && !mission.Completed {
			descriptions = append(descriptions, mission.Description)
		}
	}
	var missionMess string
	switch {
	case len(descriptions) == 2:
		missionMess = descriptions[0] + " и " + descriptions[1]
	case len(descriptions) > 0:
		missionMess = strings.Join(descriptions, ", ")
	}
	if missionMess != "" {
		return "надо " + missionMess
	}
	return "`"
}

func kitchenLookAroundMess(room int, gm GameMap, player Player) string {
	result := "ты находишься на кухне, "
	ItemsMes := gm.Rooms[room].ItemsLocaionMessage()
	if ItemsMes == "" {
		result += "ничего интересного"
	} else {
		result += ItemsMes
	}

	missMess := MissionMes(player)
	if missMess != "" {
		result += ", " + missMess + "."
	} else {
		result += "."
	}

	cgm := CanGoMes(room, gm)
	if cgm != "" {
		result += " " + cgm
	}
	return result
}

func MyRoomLookAroundMess(room int, gm GameMap, player Player) string {
	var result string
	ItemsMes := gm.Rooms[room].ItemsLocaionMessage()
	if ItemsMes == "" {
		result = "пустая комната."
	} else {
		result = ItemsMes + "."
	}
	cgm := CanGoMes(room, gm)
	if cgm != "" {
		result += " " + cgm
	}
	return result
}

func InitGameMap() GameMap {
	room0 := NewRoom(
		"кухня",
		map[string]MessageFunc{
			"LookAround": kitchenLookAroundMess,
			"Move": func(room int, gm GameMap, player Player) string {
				return "кухня, ничего интересного. " + CanGoMes(room, gm)
			},
		},
		[]ItemLocation{
			{NewItem("чай"), "на столе"},
		},
		[]RoomStatus{
			{1, false},
		},
		make(map[string]string),
	)
	room1 := NewRoom(
		"коридор",
		map[string]MessageFunc{
			"LookAround": func(room int, gm GameMap, player Player) string {
				return CanGoMes(room, gm)
			},
			"Move": func(room int, gm GameMap, player Player) string {
				return "ничего интересного. " + CanGoMes(room, gm)
			},
		},
		make([]ItemLocation, 0),
		[]RoomStatus{
			{0, false}, {3, false}, {2, true},
		},
		make(map[string]string),
	)
	room2 := NewRoom(
		"улица",
		map[string]MessageFunc{
			"Move": func(room int, gm GameMap, player Player) string {
				return "на улице весна. " + CanGoMes(room, gm)
			},
		},
		make([]ItemLocation, 0),
		[]RoomStatus{
			{1, true},
		},
		map[string]string{
			"домой": "коридор",
		},
	)
	room3 := NewRoom(
		"комната",
		map[string]MessageFunc{
			"LookAround": MyRoomLookAroundMess,
			"Move": func(room int, gm GameMap, player Player) string {
				return "ты в своей комнате. " + CanGoMes(room, gm)
			},
		},
		[]ItemLocation{
			{NewKey("ключи", 1, 2), "на столе"},
			{NewItem("конспекты"), "на столе"},
			{NewBag("рюкзак"), "на стуле"},
		},
		[]RoomStatus{
			{1, false},
		},
		make(map[string]string),
	)
	return NewGameMap([]Room{room0, room1, room2, room3})
}
