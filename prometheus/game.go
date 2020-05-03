package hangman

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Game a hangman game
type Game struct {
	Word    string `json:"word"`
	Guesses []rune `json:"guesses"`
	Tally   *Tally `json:"tally"`
}

var (
	<<!!YOUR_CODE!!> Specify 2 prometheus counters for good and bad guesses.
)

// NewGame initializes a hangman game
func NewGame(word string) *Game {
	return &Game{Word: word, Tally: NewTally([]rune(word)), Guesses: []rune{}}
}

// Guess a new letter
func (g *Game) Guess(guess rune) {
	g.validateGuess(guess)
}

func (g *Game) validateGuess(guess rune) {
	if g.Tally.Status == Won || g.Tally.Status == Lost {
		return
	}

	if g.alreadyGuessed(guess) {
		g.Tally.Status = AlreadyGuessed
		return
	}

	g.Guesses = append(g.Guesses, guess)

	<<!!YOUR_CODE!!>> Manage your prom counter
	if !g.inWord(guess) {
		g.Tally.TurnsLeft--
	}
	g.Tally.Update([]rune(g.Word), g.Guesses)
}

func (g *Game) alreadyGuessed(guess rune) bool {
	for _, l := range g.Guesses {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) inWord(guess rune) bool {
	for _, l := range g.Word {
		if l == guess {
			return true
		}
	}
	return false
}
