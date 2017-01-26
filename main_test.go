package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameAllZero(t *testing.T) {
	game := &Game{}
	rollMany(game, 20, 0)

	assert.Equal(t, 0, game.score())
}

func TestGameAllOne(t *testing.T) {
	game := &Game{}
	rollMany(game, 20, 1)

	assert.Equal(t, 20, game.score())
}

func TestOneSpare(t *testing.T) {
	game := &Game{}
	rollSpare(game)
	game.roll(3)

	rollMany(game, 17, 0)

	assert.Equal(t, 16, game.score())
}

func TestOneStrike(t *testing.T) {
	game := &Game{}
	rollStrike(game)
	game.roll(3)
	game.roll(4)
	rollMany(game, 16, 0)

	assert.Equal(t, 24, game.score())
}

func TestPerfectGame(t *testing.T) {
	game := &Game{}
	rollMany(game, 12, 10)
	assert.Equal(t, 300, game.score())
}

func rollMany(game *Game, n int, pins int) {
	for i := 0; i < n; i++ {
		game.roll(pins)
	}
}

func rollSpare(game *Game) {
	game.roll(5)
	game.roll(5) //spare
}

func rollStrike(game *Game) {
	game.roll(10) //strike
}
