package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cols = map[rune]int{
		'a': 0,
		'b': 1,
		'c': 2,
	}
	rows = map[rune]int{
		'1': 0,
		'2': 1,
		'3': 2,
	}
)

func serializeInput(str string) (x, y int, err error) {
	str = strings.ToLower(str)
	if len(str) != 2 {
		return 0, 0, os.ErrInvalid
	}
	var data [2]int
	for i, v := range str {
		val1, ok1 := cols[v]
		val2, ok2 := rows[v]
		if ok1 && i == 0 {
			data[i] = val1
		} else if ok2 && i == 1 {
			data[i] = val2
		} else {
			return 0, 0, os.ErrInvalid
		}
	}
	return data[0], data[1], nil
}

func Input(player *Player) (x, y int) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		x, y, err := serializeInput(scanner.Text())
		if err == nil {
			return x, y
		}
		fmt.Println("Error")
	}
}

func PrintBoard(game *Game) {
	fmt.Println("  A B C")
	for i, v := range game.board.Field {
		fmt.Print(i + 1)
		fmt.Print(" ")
		for _, k := range v {
			fmt.Print(game.icons[k], " ")
		}
		fmt.Print("\n")
	}
}
