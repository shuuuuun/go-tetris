package main

const number_of_block int = 4
const cols int = 12
const rows int = 12
const block_size int = 35
const hidden_rows int = number_of_block
const logical_rows int = rows + hidden_rows

type Tetris struct {
  currentBlock *Block
  board [][]int
}

func (tetris *Tetris) newGame() {
  tetris.initBoard()
  tetris.currentBlock = NewBlock(1)
}

func (tetris *Tetris) update() {
  tetris.currentBlock.moveDown()
}

func (tetris *Tetris) initBoard() {
  tetris.board = make([][]int, logical_rows)
  for r := 0; r < logical_rows; r++ {
    tetris.board[r] = make([]int, cols)
  }
}
