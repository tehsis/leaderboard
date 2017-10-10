package leaderboard

import (
	"sync"

	"github.com/go-redis/redis"
)

// RedisRepo allows to save data on redis
type redisRepo struct {
	identifier string
	client     *redis.Client
}

// NewRedisRepo setups a new client
func newRedisRepo(name string, client *redis.Client) redisRepo {
	return redisRepo{
		identifier: name,
		client:     client,
	}
}

// Get gets the score and position of username
func (r redisRepo) get(username string) (uint, uint, error) {
	var wg sync.WaitGroup
	var err error
	var pos int64
	var score float64

	wg.Add(2)

	go func() {
		defer wg.Done()
		pos, err = r.client.ZRevRank(r.identifier, username).Result()

		// Redis order is 0-based
		pos++
	}()

	go func() {
		defer wg.Done()
		score, err = r.client.ZScore(r.identifier, username).Result()
	}()

	wg.Wait()

	if err != nil {
		return 0, 0, err
	}

	return uint(score), uint(pos), nil
}

// Add adds a new score
func (r redisRepo) add(username string, score uint) (uint, uint, error) {
	r.client.ZAdd(r.identifier, redis.Z{
		Score:  float64(score),
		Member: username,
	})

	score, pos, err := r.get(username)

	if err != nil {
		return 0, 0, err
	}

	return score, pos, nil
}

// range gets the users starting at position from until position to
func (r redisRepo) repoRange(from, to uint) ([]Score, error) {

	if to < from {
		panic("from parameter can not be lower than to!")
	}

	// Redis range is 0 based
	from--
	to--

	rank, err := r.client.ZRevRangeWithScores(r.identifier, int64(from), int64(to)).Result()

	if err != nil {
		return nil, err
	}

	ranking := []Score{}

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

	return ranking, nil
}
