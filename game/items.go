package main

type Item struct {
	Name      string
	IsBag     bool
	ApplyFunc func(to_what string, room int, gm *GameMap, player *Player) string
}

func NewItem(name string) Item {
	return Item{
		Name:  name,
		IsBag: false,
		ApplyFunc: func(to_what string, room int, gm *GameMap, player *Player) string {
			return "нельзя применить"
		},
	}
}

func NewBag(name string) Item {
	return Item{
		Name:  name,
		IsBag: true,
		ApplyFunc: func(to_what string, room int, gm *GameMap, player *Player) string {
			return "нельзя применить"
		},
	}
}

func NewKey(name string, room1 int, room2 int) Item {
	return Item{
		Name:  name,
		IsBag: false,
		ApplyFunc: func(to_what string, room int, gm *GameMap, player *Player) string {
			nothingToApplyTo := "не к чему применить"
			if to_what != "дверь" {
				return nothingToApplyTo
			}
			var nextRoom int
			switch room {
			case room1:
				nextRoom = room2
			case room2:
				nextRoom = room1
			default:
				return nothingToApplyTo
			}
			if !gm.Rooms[room].IsConnected(nextRoom) {
				return nothingToApplyTo
			}
			gm.Open(room, nextRoom)
			return "дверь открыта"
		},
	}
}
