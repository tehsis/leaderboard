package leaderboard

// leaderBoardRepo talks directly to the backend of the leaderboard
type leaderBoardRepo interface {
	// add adds a new score to the repo
	add(name string, points uint) (currentPoints uint, currentPosition uint, err error)
	// get gets the current points and position of user name
	get(name string) (currentPoints uint, currentPosition uint, err error)
	// repoRange gets all scores starting at position from until position to
	repoRange(from uint, to uint) ([]Score, error)
}
