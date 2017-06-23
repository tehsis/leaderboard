package leaderboard

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/redis.v5"
)

func TestAdd(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "db:6379",
	})

	repo := newRedisRepo(uuid.NewV4().String(), client)

	_, posTehsis, err := repo.add("tehsis", 10)

	if posTehsis != 1 {
		t.Error("Expected position 1 and is ", posTehsis)
	}

	_, posTehsis, err = repo.get("tehsis")

	if err != nil {
		t.Error("Error should be nil", err)
	}

	if posTehsis != 1 {
		t.Error("Expected position 1 and is ", posTehsis)
	}
}

func TestGet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "db:6379",
	})

	repo := newRedisRepo(uuid.NewV4().String(), client)

	score, _, err := repo.add("tehsis", 10)

	if err != nil {
		t.Error("Error should be nil", err)
	}

	score, _, err = repo.get("tehsis")

	if err != nil {
		t.Error("Error should be nil", err)
	}

	if score != 10 {
		t.Error("Expected score 10 and got", score)
	}
}

func TestRange(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "db:6379",
	})

	repo := newRedisRepo(uuid.NewV4().String(), client)

	for i := 0; i < 10; i++ {
		repo.add(uuid.NewV4().String(), uint(i))
	}

	top10, err := repo.repoRange(1, 10)

	if err != nil {
		t.Error("Error should be nil", err)
	}

	for index, score := range top10 {
		if score.Points != uint(9-index) {
			t.Error("Expected score ", string(uint(index)))
		}

		if score.Username == "unknown" {
			t.Error("Unknown user found")
		}
	}
}
