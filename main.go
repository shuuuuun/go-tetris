package main

import (
  "fmt"
  "time"
  "github.com/nsf/termbox-go"
)

var tetris = Tetris{}
var startTime = time.Now()

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

  mainLoop(keyCh, timerCh)
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
  _timeSpan := 1000
  for {
    tch <- true
    time.Sleep(time.Duration(_timeSpan) * time.Millisecond)
  }
}

func mainLoop(keyCh chan termbox.Key, timerCh chan bool) {
  tetris.newGame()
  for {
    select {
    case key := <-keyCh:
      switch key {
      case termbox.KeyArrowLeft:
        tetris.moveBlockLeft()
        break
      case termbox.KeyArrowRight:
        tetris.moveBlockRight()
        break
      case termbox.KeyArrowDown:
        tetris.moveBlockDown()
        break
      case termbox.KeyArrowUp:
        tetris.rotateBlock()
        break
      case termbox.KeyEsc, termbox.KeyCtrlC:
        return
      default:
        break
      }
    case <-timerCh:
      if !tetris.isPlayng {
        return
      }
      update()
      break
    default:
      draw()
      break
    }
  }
}

func update() {
  // fmt.Println("Current Time:", time.Now().Format(time.RFC1123))
  // fmt.Println("Time:", time.Now().Sub(startTime).Seconds())

  tetris.update()
}

func draw() {
  // termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)

  elapsedTime := time.Now().Sub(startTime).Seconds()
  drawText(0, rows+2, fmt.Sprintln("Elapsed Time:", elapsedTime))
  drawText(0, rows+3, fmt.Sprintln("block id:", tetris.currentBlock.blockId))
  drawText(0, rows+4, fmt.Sprintln("isPlayng:", tetris.isPlayng))
  // drawText(0, rows+5, fmt.Sprintln("board:", len(tetris.board), len(tetris.board[0])))
  // drawText(0, rows+7, fmt.Sprintln("board:", tetris.board))

  // drawBorder()
  drawBoard()
  drawCurrentBlock()

  termbox.Flush()
}

func drawCurrentBlock() {
  block := tetris.currentBlock
  // const color = termbox.ColorDefault
  const color = termbox.ColorWhite
  for y := 0; y < number_of_block; y++ {
    for x := 0; x < number_of_block; x++ {
      if block.shape[y][x] == 0 {
        continue
      }
      drawX := x + block.x
      drawY := y + block.y - hidden_rows
      if drawY < 0 {
        continue
      }
      // termbox.SetCell(drawX + 1, drawY + 1, '■', color, color)
      termbox.SetCell(drawX + 1, drawY + 1, ' ', color, color)
    }
  }
}

func drawBoard() {
  // const color = termbox.ColorDefault
  const color = termbox.ColorWhite
  for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
      boardX := c
      boardY := r + hidden_rows
      if tetris.board[boardY][boardX] == 0 {
        termbox.SetCell(c + 1, r + 1, ' ', termbox.ColorBlack, termbox.ColorBlack)
        continue
      }
      // termbox.SetCell(c + 1, r + 1, '■', color, color)
      termbox.SetCell(c + 1, r + 1, ' ', color, color)
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

func drawText(x, y int, str string) {
  color := termbox.ColorDefault
  runes := []rune(str)

  for i := 0; i < len(runes); i += 1 {
    termbox.SetCell(x+i, y, runes[i], color, color)
  }
}
