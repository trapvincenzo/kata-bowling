package main

type Frame struct {
	state string
	score int
}

type Game struct {
	frames      []Frame
	rolls       [21]int
	currentRoll int
}

func (g *Game) roll(pins int) {
	g.rolls[g.currentRoll] = pins

	g.currentRoll++
}

func (g *Game) score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame++ {
		if isStrike(frameIndex, g) {
			score += 10 + strikeBonus(frameIndex, g)
			frameIndex++
		} else if isSpare(frameIndex, g) {
			score += 10 + spareBonus(frameIndex, g)
			frameIndex += 2
		} else {
			score += sumOfBallsInFrame(frameIndex, g)
			frameIndex += 2
		}

	}

	return score
}

func isStrike(frameIndex int, g *Game) bool {
	return g.rolls[frameIndex] == 10
}

func isSpare(frameIndex int, g *Game) bool {
	return sumOfBallsInFrame(frameIndex, g) == 10
}

func spareBonus(frameIndex int, g *Game) int {
	return g.rolls[frameIndex+2]
}

func strikeBonus(frameIndex int, g *Game) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func sumOfBallsInFrame(frameIndex int, g *Game) int {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1]
}

func main() {
	panic("Run test instead!")
}
