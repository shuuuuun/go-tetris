package main

var shapeList = [][][]int{
  {
    {0, 0, 0, 0},
    {1, 1, 1, 1},
    {0, 0, 0, 0},
    {0, 0, 0, 0},
  },
  {
    {0, 0, 0, 0},
    {0, 1, 1, 1},
    {0, 1, 0, 0},
    {0, 0, 0, 0},
  },
  {
    {0, 0, 0, 0},
    {1, 1, 1, 0},
    {0, 0, 1, 0},
    {0, 0, 0, 0},
  },
  {
    {0, 0, 0, 0},
    {0, 1, 1, 0},
    {0, 1, 1, 0},
    {0, 0, 0, 0},
  },
  {
    {0, 0, 0, 0},
    {1, 1, 0, 0},
    {0, 1, 1, 0},
    {0, 0, 0, 0},
  },
  {
    {0, 0, 0, 0},
    {0, 1, 1, 0},
    {1, 1, 0, 0},
    {0, 0, 0, 0},
  },
  {
    {0, 0, 0, 0},
    {0, 1, 0, 0},
    {1, 1, 1, 0},
    {0, 0, 0, 0},
  },
}

type Block struct {
  shape [][]int
  blockId int
  x int
  y int
}

func NewBlock(blockId int) *Block {
  block := &Block{}
  block.shape = shapeList[blockId]
  block.blockId = blockId
  return block
}

func (b *Block) moveLeft() {
  b.x -= 1
}

func (b *Block) moveRight() {
  b.x += 1
}

func (b *Block) moveDown() {
  b.y += 1
}

func (b *Block) rotate() {
  newShape := make([][]int, number_of_block)
  for y := 0; y < number_of_block; y++ {
    newShape = append(newShape, make([]int, number_of_block))
    for x := 0; x < number_of_block; x++ {
      newShape[y] = append(newShape[y], b.shape[number_of_block - 1 - x][y])
    }
  }
  b.shape = newShape
}
