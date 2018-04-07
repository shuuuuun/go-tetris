package main

import (
  // "os"
  "fmt"
  "time"
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

  keyCh := make(chan termbox.Key)
  timerCh := make(chan bool)

  go keyEventLoop(keyCh)
  go timerLoop(timerCh)

  tetris.newGame()
  
  mainLoop(keyCh, timerCh)
  // event_queue := make(chan termbox.Event)
  // go func() {
  //   for {
  //     event_queue <- termbox.PollEvent()
  //   }
  // }()
  // for {
  //   fmt.Println("update ----------------------------------------------------------------------------")
  //   update()
  // }
  // pollEvent()
  // tm.Clear() // Clear current screen
  // tetris := Tetris{}
  // tetris.newGame()
  // for {
  //   tetris.update()
  //   tetris.render()
  // }
// MAINLOOP:
//   for {
//     // fmt.Println("hoge----------------------------------------------------------------------------")
//     select {
//     case key := <-keyCh:
//     // ev := <-event_queue
//     // switch ev.Type {
//     // case termbox.EventKey:
//       fmt.Println("poyo----------------------------------------------------------------------------")
//       // switch ev.Key {
//       switch key {
//       case termbox.KeyEsc, termbox.KeyCtrlC:
//         break MAINLOOP
//       default:
//         break
//         // fmt.Println("fuga----------------------------------------------------------------------------")
//         // update()
//       }
//     case <-timerCh:
//       fmt.Println("update ----------------------------------------------------------------------------")
//       update()
//       break
//     default:
//       // fmt.Println("default ----------------------------------------------------------------------------")
//       // update()
//       break
//     }
//     // fmt.Println("moge----------------------------------------------------------------------------")
//     // update()
//   }
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
    // fmt.Println("hoge----------------------------------------------------------------------------")
    select {
    case key := <-keyCh:
    // ev := <-event_queue
    // switch ev.Type {
    // case termbox.EventKey:
      fmt.Println("poyo----------------------------------------------------------------------------")
      // switch ev.Key {
      switch key {
      case termbox.KeyEsc, termbox.KeyCtrlC:
        return
      default:
        break
        // fmt.Println("fuga----------------------------------------------------------------------------")
        // update()
      }
    case <-timerCh:
      fmt.Println("update ----------------------------------------------------------------------------")
      update()
      break
    default:
      // fmt.Println("default ----------------------------------------------------------------------------")
      // update()
      break
    }
    // fmt.Println("moge----------------------------------------------------------------------------")
    // update()
  }
}

func pollEvent() {
  // update()
  for {
    fmt.Println("hoge----------------------------------------------------------------------------")
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventKey:
      // fmt.Println("poyo----------------------------------------------------------------------------")
      switch ev.Key {
      case termbox.KeyEsc:
        fmt.Println("poyo----------------------------------------------------------------------------")
        return
        // os.Exit(0)
      default:
        fmt.Println("fuga----------------------------------------------------------------------------")
        update()
      }
    default:
      fmt.Println("piyo----------------------------------------------------------------------------")
      update()
    }
    // fmt.Println("moge----------------------------------------------------------------------------")
    // update()
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
