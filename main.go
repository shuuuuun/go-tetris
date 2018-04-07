package main

import (
  "fmt"
  "time"
  "github.com/nsf/termbox-go"
)

var tetris = Tetris{}

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  keyCh := make(chan termbox.Key)
  timerCh := make(chan bool)

  go keyEventLoop(keyCh)
  go timerLoop(timerCh)

  tetris.newGame()

  mainLoop(keyCh, timerCh)

  // tetris := Tetris{}
  // tetris.newGame()
  // for {
  //   tetris.update()
  //   tetris.render()
  // }
}

func keyEventLoop(kch chan termbox.Key) {
  for {
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventKey:
      kch <- ev.Key
    default:
    }
  }
}

func timerLoop(tch chan bool) {
  _timeSpan := 100
  for {
    tch <- true
    time.Sleep(time.Duration(_timeSpan) * time.Millisecond)
  }
}

func mainLoop(keyCh chan termbox.Key, timerCh chan bool) {
  for {
    select {
    case key := <-keyCh:
      switch key {
      case termbox.KeyEsc, termbox.KeyCtrlC:
        return
      default:
        break
      }
    case <-timerCh:
      fmt.Println("update ----------------------------------------------------------------------------")
      update()
      break
    default:
      break
    }
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
