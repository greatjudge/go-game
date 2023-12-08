package main

const absent string = "нет такого"

func LookAround(room int, gm GameMap, player Player) string {
	return gm.Rooms[room].MesFunc("LookAround")(room, gm, player)
}

func Move(room int, toRoomName string, gm GameMap, player Player) (int, string) {
	r := gm.Rooms[room]
	name := r.NameFromLocName(toRoomName)
	idx, nextRoom, ok := gm.GetRoom(name)
	if !ok || !r.IsConnected(idx) {
		return room, "нет пути в " + toRoomName
	}
	if r.IsClosed(idx) {
		return room, "дверь закрыта"
	}
	return idx, nextRoom.MesFunc("Move")(idx, gm, player)
}

func Apply(room int, what string, toWhat string, gm *GameMap, player *Player) string {
	item, ok := player.Items[what]
	if !ok {
		return "нет предмета в инвентаре - " + what
	}
	return item.ApplyFunc(toWhat, room, gm, player)
}

func Take(room int, what string, gm *GameMap, player *Player) string {
	itemLoc, ok := gm.Rooms[room].Items[what]
	if !ok {
		return absent
	}
	mes, ok := player.Take(itemLoc.Item)
	if ok {
		delete(gm.Rooms[room].Items, what)
	}
	return mes
}

func PutOn(room int, what string, gm *GameMap, player *Player) string {
	itemLoc, ok := gm.Rooms[room].Items[what]
	if !ok {
		return absent
	}
	mes, ok := player.PutOn(itemLoc.Item)
	if ok {
		delete(gm.Rooms[room].Items, what)
	}
	return mes
}
