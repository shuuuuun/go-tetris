package main

import (
  // tm "github.com/buger/goterm"
  "github.com/nsf/termbox-go"
)

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  pollEvent()
  // tm.Clear() // Clear current screen
  // tetris := Tetris{}
  // tetris.newGame()
  // for {
  //   tetris.update()
  //   tetris.render()
  // }
}

func pollEvent() {
  draw()
  for {
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventKey:
      switch ev.Key {
      case termbox.KeyEsc:
        return
      default:
        draw()
      }
    default:
      draw()
    }
  }
}

func draw() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  drawBorder()

  termbox.Flush()
}

func drawBorder() {
  const color = termbox.ColorDefault
  termbox.SetCell(0, 0, '┏', color, color)
  termbox.SetCell(cols, 0, '┓', color, color)
  termbox.SetCell(0, rows, '┗', color, color)
  termbox.SetCell(cols, rows, '┛', color, color)
}
