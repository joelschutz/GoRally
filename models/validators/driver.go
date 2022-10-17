package validators

import (
	"fmt"
	"regexp"

	"github.com/joelschutz/gorally/models"
)

func ValidateDriver(d models.Driver) (bool, error) {
	vs := []func(d models.Driver) (bool, string){}
	vs = append(vs, validateDriverName)
	vs = append(vs, validateDriverAge)
	vs = append(vs, validateDriverTerrainSkills)
	vs = append(vs, validateDriverDrivingStyle)
	vs = append(vs, validateDriverVehicleSkills)

	// validate All
	arr := []bool{}
	errs := ""
	for _, v := range vs {
		r, err := v(d)
		arr = append(arr, r)
		errs += fmt.Sprintf("%s", err)
	}
	if AND_(arr...) {
		return false, fmt.Errorf("[Driver::%s]", errs)
	}
	return true, nil
}

func validateDriverName(d models.Driver) (bool, string) {
	// No numbers in name
	hasNumbers := regexp.MustCompile("[0-9]+").Match([]byte(d.Name))
	// Max length 100
	tooBig := len(d.Name) > 100

	if AND_(hasNumbers, tooBig) {
		return false, fmt.Sprintf("[Name::hasNumbers: %v|tooBig: %v]", hasNumbers, tooBig)
	}
	return true, ""
}

func validateDriverAge(d models.Driver) (bool, string) {
	// Max age 69 - No old people
	tooOld := d.Age > 69

	if tooOld {
		return false, fmt.Sprintf("[Age::tooOld: %v]", tooOld)
	}
	return true, ""
}

func validateDriverTerrainSkills(d models.Driver) (bool, string) {
	// Max Skill Level 100
	tooGood := []bool{}
	for i := 0; i < len(d.TerrainSkills); i++ {
		tooGood[i] = d.TerrainSkills[i] > 100
	}

	if AND_(tooGood...) {
		return false, fmt.Sprintf("[TerrainSkills::tooGood: %v]", tooGood)
	}
	return true, ""
}

func validateDriverDrivingStyle(d models.Driver) (bool, string) {
	// Max Skill Level 100
	tooGood := []bool{}
	for i := 0; i < len(d.DrivingStyle); i++ {
		tooGood[i] = d.DrivingStyle[i] > 100
	}

	if AND_(tooGood...) {
		return false, fmt.Sprintf("[DrivingStyle::tooGood: %v]", tooGood)
	}
	return true, ""
}

func validateDriverVehicleSkills(d models.Driver) (bool, string) {
	// Max Skill Level 100
	tooGood := []bool{}
	for i := 0; i < len(d.VehicleSkills); i++ {
		tooGood[i] = d.VehicleSkills[i] > 100
	}

	if e := AND_(tooGood...); e {
		return false, fmt.Sprintf("[VehicleSkills::tooGood: %v]", e)
	}
	return true, ""
}
