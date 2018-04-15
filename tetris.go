package main

import (
  "time"
  "math/rand"
)

const number_of_block int = 4
const cols int = 12
const rows int = 12
const block_size int = 35
const hidden_rows int = number_of_block
const logical_rows int = rows + hidden_rows

type Tetris struct {
  currentBlock *Block
  nextBlock *Block
  board [][]int
  isPlayng bool
}

func (tetris *Tetris) newGame() {
  tetris.isPlayng = true
  tetris.initBoard()
  tetris.createCurrentBlock()
}

func (tetris *Tetris) quitGame() {
  tetris.isPlayng = false
}

func (tetris *Tetris) update() {
  if !tetris.moveBlockDown() {
    tetris.freeze()
    // tetris.clearLines()
    if tetris.checkGameOver() {
      tetris.quitGame()
      return
    }
    tetris.createCurrentBlock()
    tetris.createNextBlock()
  }
}

func (tetris *Tetris) initBoard() {
  tetris.board = make([][]int, logical_rows)
  for r := 0; r < logical_rows; r++ {
    tetris.board[r] = make([]int, cols)
  }
}

func (tetris *Tetris) createCurrentBlock() {
  if tetris.nextBlock == nil {
    tetris.createNextBlock()
  }
  tetris.currentBlock = tetris.nextBlock
}

func (tetris *Tetris) createNextBlock() {
  rand.Seed(time.Now().UnixNano())
  id := rand.Intn(len(shapeList) - 1)
  tetris.nextBlock = NewBlock(id)
}

func (tetris *Tetris) freeze() {
  for y := 0; y < number_of_block; y++ {
    for x := 0; x < number_of_block; x++ {
      boardX := x + tetris.currentBlock.x
      boardY := y + tetris.currentBlock.y
      if tetris.currentBlock.shape[y][x] == 0 || boardY < 0 {
        continue
      }
      if tetris.currentBlock.shape[y][x] != 0 {
        tetris.board[boardY][boardX] = tetris.currentBlock.blockId + 1
      } else {
        tetris.board[boardY][boardX] = 0
      }
    }
  }
}

func (tetris *Tetris) moveBlockLeft() bool {
  isValid := tetris.validate(-1, 0, tetris.currentBlock)
  if isValid {
    tetris.currentBlock.moveLeft()
  }
  return isValid
}

func (tetris *Tetris) moveBlockRight() bool {
  isValid := tetris.validate(1, 0, tetris.currentBlock)
  if isValid {
    tetris.currentBlock.moveRight()
  }
  return isValid
}

func (tetris *Tetris) moveBlockDown() bool {
  // isValid := tetris.validate(0, 1)
  isValid := tetris.validate(0, 1, tetris.currentBlock)
  if isValid {
    tetris.currentBlock.moveDown()
  }
  return isValid
}

func (tetris *Tetris) rotateBlock() bool {
  rotatedBlock := tetris.currentBlock // copy
  rotatedBlock.rotate()
  isValid := tetris.validate(0, 0, rotatedBlock)
  if isValid {
    tetris.currentBlock = rotatedBlock
  }
  return isValid
}

func (tetris *Tetris) validate(offsetX, offsetY int, block *Block) bool {
  // block = block || tetris.currentBlock
  nextX := block.x + offsetX
  nextY := block.y + offsetY
  if block.shape == nil {
    return false
  }
  for y := 0; y < number_of_block; y++ {
    for x := 0; x < number_of_block; x++ {
      if block.shape[y][x] == 0 {
        continue
      }
      boardX := x + nextX
      boardY := y + nextY
      isOutsideLeftWall := boardX < 0
      isOutsideRightWall := boardX >= cols
      isUnderBottom := boardY >= logical_rows
      isOutsideBoard := boardY >= len(tetris.board) || boardX >= len(tetris.board[boardY])
      isExistsBlock := !isOutsideBoard && tetris.board[boardY][boardX] != 0
      if isOutsideLeftWall || isOutsideRightWall || isUnderBottom || isOutsideBoard || isExistsBlock {
        return false
      }
    }
  }
  return true
}

func (tetris *Tetris) checkGameOver() bool {
  isGameOver := true
  boardY := tetris.currentBlock.y + (number_of_block - 1)
  if boardY >= hidden_rows {
    isGameOver = false
  }
  return isGameOver
}
