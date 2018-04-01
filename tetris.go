package main

import (
  tm "github.com/buger/goterm"
  // "github.com/nsf/termbox-go"
  "time"
)

var number_of_block int = 4
var cols int = 12
var rows int = 12
var block_size int = 35
var hidden_rows int = number_of_block
var logical_rows int = rows + hidden_rows

type Tetris struct {
  currentBlock Block
  // currentBlock *Block
  board [][]int
}

func (tetris *Tetris) newGame() {
  tetris.initBoard()
  // currentBlock = new(Block)
  tetris.currentBlock = Block{x: 1, y: 5}
}

func (tetris *Tetris) update() {
  tetris.currentBlock.moveDown()
}

func (tetris *Tetris) render() {
  // By moving cursor to top-left position we ensure that console output
  // will be overwritten each time, instead of adding new.
  tm.MoveCursor(1,1)

  tm.Println("Current Time:", time.Now().Format(time.RFC1123))
  tetris.drawBorder()
  // drawBoard()

  tm.Flush() // Call it every time at the end of rendering
  time.Sleep(time.Second)
}

func (tetris *Tetris) initBoard() {
  tetris.board = make([][]int, logical_rows)
  for r := 0; r < logical_rows; r++ {
    tetris.board[r] = make([]int, cols)
  }
  // tm.Println(tetris.board)
}

func (tetris *Tetris) drawBoard() {
  for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
      boardX := c
      boardY := r + hidden_rows
      if tetris.board[boardY][boardX] != 0 {
        continue
      }
      // drawBox()
      // tm.Print(" ")
    }
  }
}

func (tetris *Tetris) drawBorder() {
  for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
      // tm.Print("┌ ┏ ┓ ┐└ ┗ ┛ ┘│ ┃─ ━")
      if r == 0 && c == 0 {
        tm.Print("┌")
      } else if r == 0 && (c == cols - 1) {
        tm.Print("┐")
      } else if (r == rows - 1) && c == 0{
        tm.Print("└")
      } else if (r == rows - 1) && (c == cols - 1) {
        tm.Print("┘")
      } else if r == 0 || (r == rows - 1) {
        tm.Print("-")
      } else if c == 0 || (c == cols - 1) {
        tm.Print("|")
      } else {
        // tm.Print(" ")
        // tm.Print("■")
        is_block := tetris.currentBlock.x == c && tetris.currentBlock.y == r
        // tm.Print(c)
        // tm.Print(is_block)
        if tetris.board[r][c] == 0 && !is_block {
          tm.Print(" ")
        } else {
          tm.Print("■")
        }
      }
    }
    tm.Println("")
  }
}
