package main

import (
  tm "github.com/buger/goterm"
)

func main() {
  tetris := Tetris{}
  tm.Clear() // Clear current screen
  tetris.initBoard()
  // currentBlock = new(Block)
  tetris.currentBlock = Block{x: 0, y: 0}
  for {
    tetris.update()
    tetris.render()
  }
}
