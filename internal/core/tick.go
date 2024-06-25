package core

func Tick(game *Game, player *Player) {
	for {
		x, y := Input(player)
		if MakeMove(game, player, x, y) {
			return
		}
	}
}
