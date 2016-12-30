package leaderboard_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tehsis/leaderboard"
	"gopkg.in/redis.v5"
)

func TestAdd(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := leaderboard.NewRedisRepo(client)

	_, posTehsis := repo.Add("tehsis", 10)

	if posTehsis != 1 {
		t.Error("Expected position 1 and is ", posTehsis)
	}

	_, posTehsis = repo.Get("tehsis")

	if posTehsis != 1 {
		t.Error("Expected position 1 and is ", posTehsis)
	}
}

func TestGet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := leaderboard.NewRedisRepo(client)

	score, _ := repo.Add("tehsis", 10)

	score, _ = repo.Get("tehsis")

	if score != 10 {
		t.Error("Expected score 10 and got", score)
	}
}

func TestRange(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := leaderboard.NewRedisRepo(client)

	for i := 0; i < 10; i++ {
		repo.Add(uuid.NewV4().String(), uint(i))
	}

	top10 := repo.Range(1, 10)

	for index, score := range top10 {
		if score.Points != uint(9-index) {
			t.Error("Expected score ", string(uint(index)))
		}

		if score.Username == "unknown" {
			t.Error("Unknown user found")
		}
	}
}
