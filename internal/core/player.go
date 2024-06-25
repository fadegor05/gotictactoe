package core

type Player struct {
	id   int
	icon string
}

func CreatePlayer(id int, icon string) Player {
	return Player{id, icon}
}
