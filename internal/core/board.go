package core

type Board struct {
	Field [3][3]int
}

func CreateBoard() *Board {
	var field [3][3]int
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			field[y][x] = 0
		}
	}
	return &Board{field}
}

func GetCell(board *Board, x, y int) int {
	return board.Field[y][x]
}

func SetCell(board *Board, x, y int, value int) {
	board.Field[y][x] = value
}

type Cell struct {
	x, y int
}

func GetWinPositions() [][]Cell {
	var positions [][]Cell

	diagonal := make([]Cell, 3)
	antiDiagonal := make([]Cell, 3)
	for i := 0; i < 3; i++ {
		row := make([]Cell, 3)
		col := make([]Cell, 3)
		for j := 0; j < 3; j++ {
			row[j] = Cell{i, j}
			col[j] = Cell{j, i}
		}
		positions = append(positions, row)
		positions = append(positions, col)
		diagonal[i] = Cell{i, i}
		antiDiagonal[i] = Cell{2 - i, i}
	}
	positions = append(positions, diagonal)
	positions = append(positions, antiDiagonal)
	return positions
}
