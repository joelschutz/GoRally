package validators

import (
	"fmt"
	"math"
	"regexp"

	"github.com/joelschutz/gorally/models"
)

func ValidateTrack(d models.Track) (bool, error) {
	vs := []func(d models.Track) (bool, string){}
	vs = append(vs, validateTrackName)
	vs = append(vs, validateTrackCountry)
	vs = append(vs, validateTrackSegments)

	// validate All
	arr := []bool{}
	errs := ""
	for _, v := range vs {
		r, err := v(d)
		arr = append(arr, r)
		errs += fmt.Sprintf("%s", err)
	}
	if AND_(arr...) {
		return false, fmt.Errorf("[Track::%s]", errs)
	}
	return true, nil
}

func validateTrackName(d models.Track) (bool, string) {
	// Max length 100
	tooBig := len(d.Name) > 100

	if tooBig {
		return false, fmt.Sprintf("[Name::tooBig: %v]", tooBig)
	}
	return true, ""
}

func validateTrackCountry(d models.Track) (bool, string) {
	// No numbers in name
	hasNumbers := regexp.MustCompile("[0-9]+").Match([]byte(d.Country))
	// Max length 100
	tooBig := len(d.Country) > 100

	if AND_(hasNumbers, tooBig) {
		return false, fmt.Sprintf("[Country::hasNumbers: %v|tooBig: %v]", hasNumbers, tooBig)
	}
	return true, ""
}

func validateTrackSegments(d models.Track) (bool, string) {
	tooTight := []bool{}
	wrongDirection := []bool{}
	wrongLevel := []bool{}
	wrongTerrain := []bool{}
	for i := 0; i < len(d.Segments); i++ {
		tooTight[i] = validateSegmentLength(d.Segments[i])
		wrongDirection[i] = validateSegmentDirection(d.Segments[i])
		wrongTerrain[i] = validateSegmentTerrain(d.Segments[i])
		wrongLevel[i] = validateSegmentCorner(d.Segments[i])
	}

	if AND_(AND_(tooTight...), AND_(wrongDirection...), AND_(wrongLevel...), AND_(wrongTerrain...)) {
		return false, fmt.Sprintf("[Segments::tooTight: %v|wrongDirection: %v|wrongLevel: %v]", tooTight, wrongDirection, wrongLevel)
	}
	return true, ""
}

func validateSegmentLength(d models.Segmnent) bool {
	// No tighter than a hairpin
	teoricalPerimeter := 2 * math.Pi * 100 / float64(d.Corner)
	return (d.Length > teoricalPerimeter/2)
}

func validateSegmentDirection(d models.Segmnent) bool {
	// Only 3 options
	return (d.Direction > 2)
}

func validateSegmentCorner(d models.Segmnent) bool {
	// Only 9 options
	return (d.Corner > 8)
}

func validateSegmentTerrain(d models.Segmnent) bool {
	// Only 6 options
	return (d.Terrain > 5)
}
