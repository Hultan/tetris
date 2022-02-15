package randomizer

import (
	"fmt"
	"math/rand"
	"time"
)

type Randomizer struct {
	queueSize      int
	tetrominoCount int
	next           []int
}

func NewRandomizer(tetrominoCount, queueSize int) *Randomizer {
	if tetrominoCount <= 2 {
		panic("tetrominoCount must be above 2")
	}
	if queueSize < 1 || queueSize > 100 {
		panic("queueSize must be between 1 and 100")
	}
	t := &Randomizer{}
	t.queueSize = queueSize
	t.tetrominoCount = tetrominoCount

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < queueSize; i++ {
		t.next = append(t.next, rand.Intn(tetrominoCount))
	}

	return t
}

func (t *Randomizer) Next() int {
	n := t.next[0]

	for i := 0; i < t.queueSize-1; i++ {
		t.next[i] = t.next[i+1]
	}
	t.next[t.queueSize-1] = rand.Intn(t.tetrominoCount)

	return n
}

func (t *Randomizer) Queue() []int {
	return t.next
}

func (t *Randomizer) Print() {
	for i := 0; i < t.queueSize; i++ {
		fmt.Printf("%d", t.next[i])
	}
	fmt.Println()
}
