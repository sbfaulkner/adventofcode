package advent

import "testing"

func TestMemoryGame(t *testing.T) {
	game := NewMemoryGame(0, 3, 6)

	turns := []uint{0, 3, 3, 1, 0, 4, 0}

	for i, want := range turns {
		got := game.Play(1)
		if got != want {
			t.Errorf("Turn %d: got %d, want %d", i+4, got, want)
		}
	}
}

func TestMemoryGame2020(t *testing.T) {
	testCases := []struct {
		game *MemoryGame
		want uint
	}{
		{NewMemoryGame(0, 3, 6), 436},
		{NewMemoryGame(1, 3, 2), 1},
		{NewMemoryGame(2, 1, 3), 10},
		{NewMemoryGame(1, 2, 3), 27},
		{NewMemoryGame(2, 3, 1), 78},
		{NewMemoryGame(3, 2, 1), 438},
		{NewMemoryGame(3, 1, 2), 1836},
	}

	for i, tc := range testCases {
		got := tc.game.Play(2020 - tc.game.turns)
		if got != tc.want {
			t.Errorf("Game %d: got %d, want %d", i+1, got, tc.want)
		}
	}
}

func TestMemoryGame30000000(t *testing.T) {
	testCases := []struct {
		game *MemoryGame
		want uint
	}{
		{NewMemoryGame(0, 3, 6), 175594},
		{NewMemoryGame(1, 3, 2), 2578},
		{NewMemoryGame(2, 1, 3), 3544142},
		{NewMemoryGame(1, 2, 3), 261214},
		{NewMemoryGame(2, 3, 1), 6895259},
		{NewMemoryGame(3, 2, 1), 18},
		{NewMemoryGame(3, 1, 2), 362},
	}

	for i, tc := range testCases {
		got := tc.game.Play(30000000 - tc.game.turns)
		if got != tc.want {
			t.Errorf("Game %d: got %d, want %d", i+1, got, tc.want)
		}
	}
}

func BenchmarkMemoryGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		game := NewMemoryGame(0, 3, 6)
		game.Play(1)
	}
}

func BenchmarkMemoryGame2020(b *testing.B) {
	for i := 0; i < b.N; i++ {
		game := NewMemoryGame(0, 3, 6)
		game.Play(2020-3)
	}
}

func BenchmarkMemoryGame30000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		game := NewMemoryGame(0, 3, 6)
		game.Play(30000000-3)
	}
}
