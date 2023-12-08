package main

type Player struct {
	hasBag       bool
	Items        map[string]Item
	MissionNames []string
	Name2Mission map[string]Mission
}

func NewPlayer(missions []Mission) Player {
	names, name2Mision := NamesAndName2Mission(missions)
	return Player{
		hasBag:       false,
		Items:        make(map[string]Item),
		MissionNames: names,
		Name2Mission: name2Mision,
	}
}

func NamesAndName2Mission(missions []Mission) ([]string, map[string]Mission) {
	names := make([]string, len(missions))
	name2Mission := make(map[string]Mission, len(missions))
	for _, m := range missions {
		names = append(names, m.Name)
		name2Mission[m.Name] = m
	}
	return names, name2Mission
}

func (p *Player) PutOn(item Item) (string, bool) {
	if item.IsBag && !p.hasBag {
		p.hasBag = true
		return "вы надели: " + item.Name, true
	}
	return "нельзя надеть", false
}

func PackBagMission(p *Player) {
	mission, ok := player.Name2Mission["PackBag"]
	if ok && p.hasBag {
		_, allIn := player.Items["конспекты"]
		if allIn {
			mission.Completed = true
			player.Name2Mission["PackBag"] = mission
		}
	}
}

func (p *Player) Take(item Item) (string, bool) {
	if p.hasBag {
		p.Items[item.Name] = item
		PackBagMission(p)
		return "предмет добавлен в инвентарь: " + item.Name, true
	}
	return "некуда класть", false
}

type Mission struct {
	Name        string
	Description string
	Completed   bool
}

func NewMission(description string) Mission {
	return Mission{
		Description: description,
		Completed:   false,
	}
}
