package main

import (
  tm "github.com/buger/goterm"
)

func main() {
  tm.Clear() // Clear current screen
  tetris := Tetris{}
  tetris.newGame()
  for {
    tetris.update()
    tetris.render()
  }
}
