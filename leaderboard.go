package leaderboard

import (
	"gopkg.in/redis.v5"
)

// Score represents the points obtained by username
type Score struct {
	Username string
	Points   uint
}

// Board represents usernames that obtained scores ordered
type Board interface {
	Set(username string, points uint) (currentPosition uint)
	Get(username string) (currentScore uint, currentPosition uint)
	GetTop(n int) Score
}

// LeaderBoard is a list of scores
type LeaderBoard struct {
	name   string
	repo   leaderBoardRepo
	scores []Score
}

// NewRedisLeaderBoard buils a leaderboard using a redis repo
func NewRedisLeaderBoard(name string, redisClient *redis.Client) LeaderBoard {
	repo := newRedisRepo(name, redisClient)

	return NewLeaderBoard(name, repo)
}

// NewLeaderBoard builds a leaderboard using a custom repo
func NewLeaderBoard(name string, repo leaderBoardRepo) LeaderBoard {
	return LeaderBoard{
		name: name,
		repo: repo,
	}
}

// Set adds a new score to the leaderboard returning its position
func (l *LeaderBoard) Set(n string, s uint) (currentScore uint, err error) {
	_, pos, err := l.repo.add(n, s)

	return pos, err
}

// Get returns the score recorded for n and the position in the leaderboard
func (l *LeaderBoard) Get(n string) (currentScore uint, currentPosition uint, err error) {
	return l.repo.get(n)
}

// GetTop returns the n best scores in the leaderboard
func (l *LeaderBoard) GetTop(n uint) ([]Score, error) {
	return l.repo.repoRange(1, n)
}
