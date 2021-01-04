package advent

// MemoryGame is the current state for the elves memory game
type MemoryGame struct {
	index map[uint]uint
	last  uint
	turns uint
}

// NewMemoryGame starts a new game
func NewMemoryGame(nums ...uint) *MemoryGame {
	g := MemoryGame{index: make(map[uint]uint)}

	for _, n := range nums {
		g.add(n)
	}

	return &g
}

func (g *MemoryGame) add(n uint) {
	g.index[g.last] = g.turns
	g.last = n
	g.turns++
}

// Play next turn of MemoryGame
func (g *MemoryGame) Play(n uint) uint {
	var p uint

	for ; n > 0; n-- {
		p = g.play()
	}

	return p
}

func (g *MemoryGame) play() uint {
	var p uint

	i := g.index[g.last]
	if i > 0 {
		p = g.turns - i
	}

	g.add(p)

	return p
}
