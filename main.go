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

  drawBox(0, 0)
  drawBox(2, 1)

  termbox.Flush()
}

func drawBox(x, y int) {
  const color = termbox.ColorDefault
  termbox.SetCell(x, y, '┏', color, color)
  termbox.SetCell(x+1, y, '┓', color, color)
  termbox.SetCell(x, y+1, '┗', color, color)
  termbox.SetCell(x+1, y+1, '┛', color, color)
}
