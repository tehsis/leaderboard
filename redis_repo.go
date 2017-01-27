package leaderboard

import (
	uuid "github.com/satori/go.uuid"

	redis "gopkg.in/redis.v5"
)

// RedisRepo allows to save data on redis
type redisRepo struct {
	identifier string
	client     *redis.Client
}

// NewRedisRepo setups a new client
func newRedisRepo(client *redis.Client) redisRepo {
	return redisRepo{
		identifier: uuid.NewV4().String(),
		client:     client,
	}
}

// Get gets the score and position of username
func (r redisRepo) get(username string) (uint, uint) {
	pos, err := r.client.ZRevRank(r.identifier, username).Result()

	score, err := r.client.ZScore(r.identifier, username).Result()

	if err != nil {
		return 0, 0
	}

	// Redis order is 0-based
	pos++

	return uint(score), uint(pos)
}

// Add adds a new score
func (r redisRepo) add(username string, score uint) (uint, uint) {
	r.client.ZAdd(r.identifier, redis.Z{
		Score:  float64(score),
		Member: username,
	})

	score, pos := r.get(username)

	return score, pos
}

// range gets the users starting at position from until position to
func (r redisRepo) repoRange(from, to uint) []Score {

	if to < from {
		panic("from parameter can not be lower than to!")
	}

	// Redis range is 0 based
	from--
	to--

	rank, _ := r.client.ZRevRangeWithScores(r.identifier, int64(from), int64(to)).Result()

	ranking := make([]Score, 0)

	for _, r := range rank {
		username, ok := r.Member.(string)

		if !ok {
			username = "unknown"
		}

		newScore := Score{
			Username: username,
			Points:   uint(r.Score),
		}

		ranking = append(ranking, newScore)
	}

	return ranking
}
