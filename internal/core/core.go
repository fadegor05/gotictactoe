package core

import "fmt"

func Init() {
	p1 := CreatePlayer(1, "X")
	p2 := CreatePlayer(2, "O")
	game := CreateGame([]*Player{&p1, &p2})
	Loop(&game)
}

func Loop(game *Game) {
	for {
		PrintBoard(game)
		fmt.Println("Now", game.players[0].icon, "turn!")
		Tick(game, game.players[0])
		player, isWin := CheckWin(game.players, game.board)
		if isWin {
			fmt.Println(player.icon, "Won")
			break
		}
		if CheckTie(game.board) {
			fmt.Println("Tie")
			break
		}
		SwapPlayers(game)
	}
}

func SwapPlayers(game *Game) {
	game.players = append(game.players[1:], game.players[0])
}
