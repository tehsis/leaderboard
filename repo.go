package leaderboard

type leaderBoardRepo interface {
	Add(string, uint) (uint, uint)
	Get(string) (uint, uint)
	Range(uint, uint) []Score
}
