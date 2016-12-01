package leaderboard

import "fmt"

// Score represents the points obtained by username
type Score struct {
	Username string
	Points   uint
}

// Board represents usernames that obtained scores ordered
type Board interface {
	Set(string, uint) uint
	Get(string) (uint, uint)
	GetTop(int) Score
}

// LeaderBoard is a list of scores
type LeaderBoard struct {
	scores []Score
}

// New creates a new leaderboard
func (l *LeaderBoard) New() {
	l.scores = make([]Score, 0)
}

// Set adds a new score to the leaderboard returning its position
func (l *LeaderBoard) Set(n string, s uint) uint {
	l.scores = append(l.scores, Score{
		Username: n,
		Points:   s,
	})
	fmt.Printf("Scores %v points: %v, username: %v", l, n, s)
	return 1
}

// Get returns the score recorded for n and the position in the leaderboard
func (l *LeaderBoard) Get(n string) (uint, uint) {
	var r Score
	var position uint

	for i, s := range l.scores {
		if s.Username == n {
			r = s
			position = uint(i)
		}
	}

	return position, r.Points
}

// GetTop returns the n best scores in the leaderboard
func (l *LeaderBoard) GetTop(n uint) []Score {
	if len(l.scores) < int(n) {
		return l.scores[:]
	}

	return l.scores[:n]
}
