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
	repo   leaderBoardRepo
	scores []Score
}

// NewRedisLeaderBoard buils a leaderboard using a redis repo
func NewRedisLeaderBoard(redisClient *redis.Client) LeaderBoard {
	repo := newRedisRepo(redisClient)

	return NewLeaderBoard(repo)
}

// NewLeaderBoard builds a leaderboard using a custom repo
func NewLeaderBoard(repo leaderBoardRepo) LeaderBoard {
	return LeaderBoard{
		repo: repo,
	}
}

// Set adds a new score to the leaderboard returning its position
func (l *LeaderBoard) Set(n string, s uint) (currentScore uint) {
	_, pos := l.repo.add(n, s)

	return pos
}

// Get returns the score recorded for n and the position in the leaderboard
func (l *LeaderBoard) Get(n string) (currentScore uint, currentPosition uint) {
	return l.repo.get(n)
}

// GetTop returns the n best scores in the leaderboard
func (l *LeaderBoard) GetTop(n uint) []Score {
	return l.repo.repoRange(1, n)
}
