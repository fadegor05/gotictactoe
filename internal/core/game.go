package core

type Game struct {
	board   *Board
	players []*Player
	icons   map[int]string
}

var winPositions = GetWinPositions()

func CreateGame(players []*Player) Game {
	var icons = map[int]string{0: " "}
	for _, player := range players {
		icons[player.id] = player.icon
	}
	return Game{CreateBoard(), players, icons}
}

func MakeMove(game *Game, p *Player, x, y int) bool {
	if GetCell(game.board, x, y) != 0 {
		return false
	}
	SetCell(game.board, x, y, p.id)
	return true
}

func CheckWin(players []*Player, board *Board) (player *Player, isWin bool) {
	for _, player := range players {
		wins := 0
		for _, item := range winPositions {
			matches := 0
			for _, positions := range item {
				if GetCell(board, positions.x, positions.y) == player.id {
					matches++
				}
			}
			if matches == 3 {
				wins++
			}
		}
		if wins >= 1 {
			return player, true
		}
	}
	return players[0], false
}

func CheckTie(board *Board) bool {
	filled := 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if GetCell(board, x, y) != 0 {
				filled++
			}
		}
	}
	return filled == 9
}
