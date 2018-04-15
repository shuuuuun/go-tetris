package main

const number_of_block int = 4
const cols int = 12
const rows int = 12
const block_size int = 35
const hidden_rows int = number_of_block
const logical_rows int = rows + hidden_rows

type Tetris struct {
  currentBlock *Block
  board [][]int
}

func (tetris *Tetris) newGame() {
  tetris.initBoard()
  tetris.currentBlock = NewBlock(1)
}

func (tetris *Tetris) update() {
  // tetris.currentBlock.moveDown()
  if !tetris.moveBlockDown() {
    tetris.freeze()
    // tetris.clearLines()
    // if tetris.checkGameOver() {
    //   tetris.quitGame()
    //   return false
    // }
    // tetris.frameCount += 1
    // tetris.createCurrentBlock()
    // tetris.createNextBlock()
  }
}

func (tetris *Tetris) initBoard() {
  tetris.board = make([][]int, logical_rows)
  for r := 0; r < logical_rows; r++ {
    tetris.board[r] = make([]int, cols)
  }
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
        tetris.board[boardY][boardX] = tetris.currentBlock.block_id + 1
      } else {
        tetris.board[boardY][boardX] = 0
      }
    }
  }
}

func (tetris *Tetris) moveBlockDown() bool {
  // isValid := tetris.validate(0, 1)
  isValid := tetris.validate(0, 1, tetris.currentBlock)
  if isValid {
    tetris.currentBlock.moveDown()
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
