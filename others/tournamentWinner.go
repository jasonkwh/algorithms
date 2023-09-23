package others

const HOME_TEAM_WON = 1
const AWAY_TEAM_WON = 0

func TournamentWinner(competitions [][]string, results []int) string {
	teamsScore := make(map[string]int)

	for i := 0; i < len(competitions); i++ {
		score := scores(competitions[i], results[i])

		if score != "" {
			teamsScore[scores(competitions[i], results[i])] += 3
		}
	}

	return findHighestScore(teamsScore)
}

func findHighestScore(results map[string]int) string {
	team := ""
	score := 0
	for t, s := range results {
		if s >= score {
			score = s
			team = t
		}
	}
	return team
}

func scores(competition []string, result int) string {
	switch result {
	case HOME_TEAM_WON:
		return competition[0]
	default:
		return competition[1]
	}
}
