package main

import (
  tm "github.com/buger/goterm"
  "time"
)

var number_of_block int = 4
var cols int = 12
var rows int = 12
var block_size int = 35
var hidden_rows int = number_of_block
var logical_rows int = rows + hidden_rows
var currentBlock Block
// var currentBlock *Block

var board [][]int

type Tetris struct {
}

func (tetris *Tetris) update() {
  currentBlock.moveDown()
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
  board = make([][]int, logical_rows)
  for r := 0; r < logical_rows; r++ {
    board[r] = make([]int, cols)
  }
  // tm.Println(board)
}

func (tetris *Tetris) drawBoard() {
  for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
      boardX := c
      boardY := r + hidden_rows
      if board[boardY][boardX] != 0 {
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
        // currentBlock.x
        if board[r][c] == 0 {
          tm.Print(" ")
        } else {
          tm.Print("■")
        }
      }
    }
    tm.Println("")
  }
}
