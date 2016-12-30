package leaderboard_test

import (
	"testing"

	"github.com/tehsis/leaderboard"
)

func TestLeaderBoardSet(t *testing.T) {
	lb := leaderboard.NewRedisLeaderBoard("localhost:6379")

	pos := lb.Set("tehsis", 20)

	if pos != 1 {
		t.Error("Position should be 1 but is ", pos)
	}

	pos = lb.Set("lenny", 50)

	if pos != 1 {
		t.Error("Position should be 1 but is ", pos)
	}

	pos = lb.Set("carl", 10)

	if pos != 3 {
		t.Error("Position should be 1 but is ", pos)
	}
}

func TestLeaderboardGet(t *testing.T) {
	lb := leaderboard.NewRedisLeaderBoard("localhost:6379")

	lb.Set("tehsis", 20)

	score, pos := lb.Get("tehsis")

	if score != 20 && pos != 1 {
		t.Error("Position should be 1 and score 20")
	}

	lb.Set("lenny", 30)

	score, pos = lb.Get("lenny")

	if score != 30 && pos != 1 {
		t.Error("Position should be 1 and score 30")
	}

	score, pos = lb.Get("tehsis")

	if score != 20 && pos != 2 {
		t.Error("Position should be 2 and score 20")
	}

}

func TestLeaderboardGetTop(t *testing.T) {
	lb := leaderboard.NewRedisLeaderBoard("localhost:6379")
	lb.Set("tehsis", 20)
	lb.Set("lenny", 30)
	lb.Set("carl", 10)

	top3 := lb.GetTop(3)

	for index, score := range top3 {
		if index == 0 && score.Username != "lenny" {
			t.Error("username should be lenny and is", score.Username, index)
		}

		if index == 1 && score.Username != "tehsis" {
			t.Error("username should be tehsis and is", score.Username, index)
		}

		if index == 2 && score.Username != "carl" {
			t.Error("username should be carl and is", score.Username, index)
		}
	}

}
