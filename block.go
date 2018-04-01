package main

type Block struct {
  shape [][]int
  block_id int
  x int
  y int
}

func (b Block) moveDown() {
  b.y += 1
}
