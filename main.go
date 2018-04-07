package main

import (
  "fmt"
  // tm "github.com/buger/goterm"
  "github.com/nsf/termbox-go"
)

var tetris = Tetris{}

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  tetris.newGame()
  pollEvent()
  // tm.Clear() // Clear current screen
  // tetris := Tetris{}
  // tetris.newGame()
  // for {
  //   tetris.update()
  //   tetris.render()
  // }
// MAINLOOP:
//   for {
//     fmt.Println("hoge----------------------------------------------------------------------------")
//     switch ev := termbox.PollEvent(); ev.Type {
//     case termbox.EventKey:
//       fmt.Println("poyo----------------------------------------------------------------------------")
//       switch ev.Key {
//       case termbox.KeyEsc:
//         break MAINLOOP
//       default:
//         fmt.Println("fuga----------------------------------------------------------------------------")
//         update()
//       }
//     default:
//       fmt.Println("piyo----------------------------------------------------------------------------")
//       update()
//     }
//     fmt.Println("moge----------------------------------------------------------------------------")
//     update()
//   }
}

func pollEvent() {
  update()
  for {
    fmt.Println("hoge----------------------------------------------------------------------------")
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventKey:
      fmt.Println("poyo----------------------------------------------------------------------------")
      switch ev.Key {
      case termbox.KeyEsc:
        return
      default:
        fmt.Println("fuga----------------------------------------------------------------------------")
        update()
      }
    default:
      fmt.Println("piyo----------------------------------------------------------------------------")
      update()
    }
    fmt.Println("moge----------------------------------------------------------------------------")
    update()
  }
}

func update() {
  tetris.update()
  // tetris.render()
  draw()
}

func draw() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  drawBorder()
  drawBoard()

  termbox.Flush()
}

func drawBoard() {
  const color = termbox.ColorDefault
  for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
      boardX := c
      boardY := r + hidden_rows
      if tetris.board[boardY][boardX] == 0 {
        continue
      }
      termbox.SetCell(c + 1, r + 1, '■', color, color)
    }
  }
}

func drawBorder() {
  const color = termbox.ColorDefault
  termbox.SetCell(0, 0, '┏', color, color)
  termbox.SetCell(cols + 1, 0, '┓', color, color)
  termbox.SetCell(0, rows + 1, '┗', color, color)
  termbox.SetCell(cols + 1, rows + 1, '┛', color, color)
}
