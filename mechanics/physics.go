package mechanics

import (
	"math"

	"github.com/joelschutz/gorally/models"
)

func GetBasicSplinePoint(s models.Spline2D, t float64) models.Point2D {
	// point indexes
	var p0, p1, p2, p3 int
	if !s.Looped() {
		p1 = int(t) + 1
		p2 = p1 + 1
		p3 = p2 + 1
		p0 = p1 - 1
	} else {
		p1 = int(t)
		p2 = (p1 + 1) % s.Size()
		p3 = (p2 + 1) % s.Size()
		if p1 >= 1 {
			p0 = p1 - 1
		} else {
			p0 = s.Size() - 1
		}
	}
	// fmt.Println(p1, p3)

	// cache of t squared and cubic
	t = t - float64(int(t))
	tt := t * t
	ttt := tt * t

	q1 := -ttt + 2*tt - t
	q2 := 3*ttt - 5*tt + 2
	q3 := -3*ttt + 4*tt + t
	q4 := ttt - tt

	return models.NewBasicPoint2D(
		(s.GetNode(p0).X()*q1+s.GetNode(p1).X()*q2+s.GetNode(p2).X()*q3+s.GetNode(p3).X()*q4)/2,
		(s.GetNode(p0).Y()*q1+s.GetNode(p1).Y()*q2+s.GetNode(p2).Y()*q3+s.GetNode(p3).Y()*q4)/2,
	)
}

func GetBasicSplineGradient(s models.Spline2D, t float64) models.Point2D {
	// point indexes
	var p0, p1, p2, p3 int
	if !s.Looped() {
		p1 = int(t) + 1
		p2 = p1 + 1
		p3 = p2 + 1
		p0 = p1 - 1
	} else {
		p1 = int(t)
		p2 = (p1 + 1) % s.Size()
		p3 = (p2 + 1) % s.Size()
		if p1 >= 1 {
			p0 = p1 - 1
		} else {
			p0 = s.Size() - 1
		}
	}
	// fmt.Println(p1, p3)

	// cache of t squared and cubic
	t = t - float64(int(t))
	tt := t * t

	q1 := -3*tt + 4*t - 1
	q2 := 9*tt - 10*t
	q3 := -9*tt + 8*t + 1
	q4 := 3*tt - 2*t

	return models.NewBasicPoint2D(
		(s.GetNode(p0).X()*q1+s.GetNode(p1).X()*q2+s.GetNode(p2).X()*q3+s.GetNode(p3).X()*q4)/2,
		(s.GetNode(p0).Y()*q1+s.GetNode(p1).Y()*q2+s.GetNode(p2).Y()*q3+s.GetNode(p3).Y()*q4)/2,
	)
}

func IsPointInSplineBoarders(s models.BorderedSpline2D, p models.Point2D, resolution float64) bool {
	return IsPointInSplineSectionBoarders(s, p, resolution, 0., float64(s.Size())-3.)
}

func IsPointInSplineSectionBoarders(s models.BorderedSpline2D, p models.Point2D, resolution, offset, limit float64) bool {
	// Loop tru all valid t in the slice.
	// The resolution offset is just for now, the index can not be grater than the spline size in gradient
	for i := offset; i < limit; i += resolution {
		// South adge of area
		sg := GetBasicSplineGradient(s, i)
		sp := GetBasicSplinePoint(s, i)
		sr := math.Atan2(-sg.Y(), sg.X())

		s1 := models.NewBasicPoint2D(
			((s.Width(int(i))/2)*math.Sin(sr) + sp.X()),
			((s.Width(int(i))/2)*math.Cos(sr) + sp.Y()),
		)
		s2 := models.NewBasicPoint2D(
			(-(s.Width(int(i))/2)*math.Sin(sr) + sp.X()),
			(-(s.Width(int(i))/2)*math.Cos(sr) + sp.Y()),
		)
		// North adge of area
		ng := GetBasicSplineGradient(s, i+resolution)
		np := GetBasicSplinePoint(s, i+resolution)
		nr := math.Atan2(-ng.Y(), ng.X())

		n1 := models.NewBasicPoint2D(
			((s.Width(int(i+resolution))/2)*math.Sin(nr) + np.X()),
			((s.Width(int(i+resolution))/2)*math.Cos(nr) + np.Y()),
		)
		n2 := models.NewBasicPoint2D(
			(-(s.Width(int(i))/2)*math.Sin(nr) + np.X()),
			(-(s.Width(int(i))/2)*math.Cos(nr) + np.Y()),
		)

		// This square is a estimated section of the spline
		if IsPointInSquare(models.Square2D{s1, s2, n1, n2}, p) {
			return true
		} else {
			// This line is the north edge, the exiting side for the section
			if !IsPointAboveLine(models.Line2D{n1, n2}, p) {
				return false
			}
		}

	}
	return false
}

func IsPointAboveSection(s models.BorderedSpline2D, p models.Point2D, position float64) bool {
	// North adge of area
	ng := GetBasicSplineGradient(s, position)
	np := GetBasicSplinePoint(s, position)
	nr := math.Atan2(-ng.Y(), ng.X())

	n1 := models.NewBasicPoint2D(
		((s.Width(int(position))/2)*math.Sin(nr) + np.X()),
		((s.Width(int(position))/2)*math.Cos(nr) + np.Y()),
	)
	n2 := models.NewBasicPoint2D(
		(-(s.Width(int(position))/2)*math.Sin(nr) + np.X()),
		(-(s.Width(int(position))/2)*math.Cos(nr) + np.Y()),
	)
	// This line is the north edge, the exiting side for the section

	return IsPointAboveLine(models.Line2D{n1, n2}, p)
}

func IsPointInSquare(s models.Square2D, p models.Point2D) bool {
	// Find area foi each triangle
	a := models.TriangleArea(models.Traingle2D{s[0], s[1], s[2]}) + models.TriangleArea(models.Traingle2D{s[0], s[2], s[3]})

	// Find the areas of the segments
	a0 := models.TriangleArea(models.Traingle2D{p, s[0], s[1]})
	a1 := models.TriangleArea(models.Traingle2D{p, s[1], s[2]})
	a2 := models.TriangleArea(models.Traingle2D{p, s[2], s[3]})
	a3 := models.TriangleArea(models.Traingle2D{p, s[3], s[0]})

	// returns if equals
	return a == (a0 + a1 + a2 + a3)
}

func IsPointAboveLine(l models.Line2D, p models.Point2D) bool {
	// Calcs the tangent of line
	teta := math.Atan2(l[1].Y()-l[0].Y(), l[1].X()-l[0].X())

	//Checks if translated p.y is over the line
	return math.Cos(teta)*p.Y() > 0
}

func DistanceBetweenTwoPoints(p1, p2 models.Point2D) float64 {
	return math.Sqrt(math.Pow(p2.X()-p1.X(),2)+math.Pow(p2.Y()-p1.Y(),2))
}
