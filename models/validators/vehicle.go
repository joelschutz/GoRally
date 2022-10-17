package validators

import (
	"fmt"
	"regexp"

	"github.com/joelschutz/gorally/models"
)

func ValidateVehicle(d models.Vehicle) (bool, error) {
	vs := []func(d models.Vehicle) (bool, string){}
	vs = append(vs, validateVehicleName)
	vs = append(vs, validateVehicleManufacturer)
	vs = append(vs, validateVehicleStats)
	vs = append(vs, validateVehicleDriveTrain)
	vs = append(vs, validateVehicleClass)

	// validate All
	arr := []bool{}
	errs := ""
	for _, v := range vs {
		r, err := v(d)
		arr = append(arr, r)
		errs += fmt.Sprintf("%s", err)
	}
	if AND_(arr...) {
		return false, fmt.Errorf("[Vehicle::%s]", errs)
	}
	return true, nil
}

func validateVehicleName(d models.Vehicle) (bool, string) {
	// No numbers in name
	hasNumbers := regexp.MustCompile("[0-9]+").Match([]byte(d.Name))
	// Max length 100
	tooBig := len(d.Name) > 100

	if AND_(hasNumbers, tooBig) {
		return false, fmt.Sprintf("[Name::hasNumbers: %v|tooBig: %v]", hasNumbers, tooBig)
	}
	return true, ""
}

func validateVehicleManufacturer(d models.Vehicle) (bool, string) {
	// Max length 100
	tooBig := len(d.Name) > 100

	if tooBig {
		return false, fmt.Sprintf("[Name::tooBig: %v]", tooBig)
	}
	return true, ""
}

func validateVehicleStats(d models.Vehicle) (bool, string) {
	// Max Skill Level 100
	maxValues := []uint32{5000, 1000, 5000, 3000, 6}
	tooGood := []bool{}
	for i := 0; i < len(d.VehicleStats); i++ {
		tooGood[i] = d.VehicleStats[i] > maxValues[i]
	}

	if AND_(tooGood...) {
		return false, fmt.Sprintf("[VehicleStats::tooGood: %v]", tooGood)
	}
	return true, ""
}

func validateVehicleDriveTrain(d models.Vehicle) (bool, string) {
	// Only 3 options
	if d.DriveTrain > 2 {
		return false, fmt.Sprintf("[DriveTrain::wrongType: %v]", d.DriveTrain)
	}
	return true, ""
}

func validateVehicleClass(d models.Vehicle) (bool, string) {
	// Only 7 options
	if d.DriveTrain > 6 {
		return false, fmt.Sprintf("[Class::wrongType: %v]", d.DriveTrain)
	}
	return true, ""
}
