package mechanics

import "github.com/joelschutz/gorally/models"

func CalcEventResults(e models.Event) (er models.EventResults) {
	for _, comp := range e.Competitors {
		for i, t := range e.Tracks {
			er[comp][i] = calcCompetitorResult(comp, t)
		}
	}
	return er
}

func calcCompetitorResult(c models.Competitor, t models.Track) models.TrackResult {
	return CalcTrackTime(t, c.Driver, c.Vehicle)
}

func calcEventRanking(er models.EventResults) (r models.Ranking) {
	for comp, tr := range er {
		r[comp] = sumTrackTimes(tr)
	}
	return r
}

func calcTrackRanking(er models.EventResults, trackIndex int) (r models.Ranking) {
	for comp, tr := range er {
		r[comp] = tr[trackIndex].TotalTime
	}
	return r
}

func sumTrackTimes(tr []models.TrackResult) (sum float64) {
	for _, v := range tr {
		sum += v.TotalTime
	}
	return sum
}
