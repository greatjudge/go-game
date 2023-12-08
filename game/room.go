package main

import (
	"sort"
	"strings"
)

type MessageFunc func(room int, gm GameMap, player Player) string

type ItemLocation struct {
	Item     Item
	Location string
}

type RoomStatus struct {
	Room     int
	IsClosed bool
}

type Room struct {
	Name                   string
	ActionName2MessageFunc map[string]MessageFunc
	Items                  map[string]ItemLocation
	AdjRoomsOrdered        []int
	AdjRooms               map[int]bool
	LocName2RoomName       map[string]string
	RoomName2LocName       map[string]string
}

func NewRoom(name string,
	actionName2MessageFunc map[string]MessageFunc,
	items []ItemLocation,
	adjRoomsOrdered []RoomStatus,
	locName2RoomName map[string]string) Room {
	return Room{
		Name:                   name,
		ActionName2MessageFunc: actionName2MessageFunc,
		Items:                  FormItems(items),
		AdjRoomsOrdered:        FormAdjRoomsOrdered(adjRoomsOrdered),
		AdjRooms:               FormAdjRooms(adjRoomsOrdered),
		LocName2RoomName:       locName2RoomName,
		RoomName2LocName:       FlipMap(locName2RoomName),
	}
}

func FormItems(itemLocs []ItemLocation) map[string]ItemLocation {
	items := make(map[string]ItemLocation, len(itemLocs))
	for _, il := range itemLocs {
		items[il.Item.Name] = il
	}
	return items
}

func FormAdjRoomsOrdered(adjRoomStatus []RoomStatus) []int {
	adjRooms := make([]int, len(adjRoomStatus))
	for i, rs := range adjRoomStatus {
		adjRooms[i] = rs.Room
	}
	return adjRooms
}

func FormAdjRooms(adjRoomsOrdered []RoomStatus) map[int]bool {
	AdjRooms := make(map[int]bool, len(adjRoomsOrdered))
	for _, rs := range adjRoomsOrdered {
		AdjRooms[rs.Room] = rs.IsClosed
	}
	return AdjRooms
}

func FlipMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for key, val := range in {
		out[val] = key
	}
	return out
}

func (room Room) location2Item() (out map[string][]Item) {
	out = make(map[string][]Item)
	for _, val := range room.Items {
		out[val.Location] = append(out[val.Location], val.Item)
	}
	return
}

func itemNames(items []Item) string {
	names := make([]string, len(items))
	for i, item := range items {
		names[i] = item.Name
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

func (room Room) ItemsLocaionMessage() (result string) {
	location2Item := room.location2Item()
	locations := make([]string, 0, len(location2Item))
	for where, items := range location2Item {
		names := itemNames(items)
		if names != "" {
			locations = append(locations, where+": "+names)
		}
	}
	sort.Strings(locations)
	return strings.Join(locations, ", ")
}

func (room Room) NameFromLocName(locName string) string {
	name, ok := room.LocName2RoomName[locName]
	if ok {
		return name
	}
	return locName
}

func (room Room) LocNameFromName(name string) string {
	locName, ok := room.RoomName2LocName[name]
	if ok {
		return locName
	}
	return name
}

func (room Room) NextRoomsMessage(gm GameMap) string {
	names := make([]string, 0, len(room.AdjRoomsOrdered))
	for _, r := range room.AdjRoomsOrdered {
		names = append(names, room.LocNameFromName(gm.Rooms[r].Name))
	}
	return strings.Join(names, ", ")
}

func (room Room) IsConnected(otherRoom int) bool {
	_, ok := room.AdjRooms[otherRoom]
	return ok
}

func (room Room) IsClosed(otherRoom int) bool {
	closed := room.AdjRooms[otherRoom]
	return closed
}

func (room Room) MesFunc(name string) MessageFunc {
	mesFunc, ok := room.ActionName2MessageFunc[name]
	if ok {
		return mesFunc
	}
	return func(int, GameMap, Player) string { return "" }
}

type GameMap struct {
	Rooms    []Room
	Name2Idx map[string]int
}

func NewGameMap(rooms []Room) GameMap {
	return GameMap{
		Rooms:    rooms,
		Name2Idx: FormName2Idx(rooms),
	}
}

func FormName2Idx(rooms []Room) map[string]int {
	name2Idx := make(map[string]int, len(rooms))
	for i, r := range rooms {
		name2Idx[r.Name] = i
	}
	return name2Idx
}

func (gm GameMap) GetRoom(name string) (int, Room, bool) {
	i, ok := gm.Name2Idx[name]
	if ok {
		return i, gm.Rooms[i], ok
	}
	return 0, Room{}, ok
}

func (gm *GameMap) Open(room1 int, room2 int) {
	gm.Rooms[room1].AdjRooms[room2] = false
	gm.Rooms[room2].AdjRooms[room1] = false
}
