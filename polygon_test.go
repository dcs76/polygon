package polygon

import (
	"log"
	"testing"
)

var polygonPoints = []float64{
	29.744586, -95.3630671,
	29.7151455, -95.3151736,
	29.752336, -95.2997241,
	29.7564344, -95.2950034,
	29.7676852, -95.2942309,
	29.7730495, -95.2913127,
	29.7778919, -95.2913985,
	29.7742415, -95.3010974,
	29.7743905, -95.3225551,
	29.7691008, -95.3362021,
	29.7689518, -95.3432403,
	29.7702184, -95.3471885,
	29.7700694, -95.3522525,
	29.7670147, -95.3595481,
	29.7701439, -95.3658137,
	29.7655246, -95.3672729,
	29.7620972, -95.3737102,
	29.7594149, -95.3751693,
	29.7525596, -95.3745685,
	29.7501005, -95.3728518,
	29.744586, -95.3630671,
}

func containsPointTest(point Point) bool {
	points := []Point{}
	for i := 0; i < len(polygonPoints); i += 2 {
		lat, lon := polygonPoints[i], polygonPoints[i+1]
		points = append(points, Point{Lat: lat, Lon: lon})
	}
	polygon, err := NewPolygonAsSlice(points)
	if err != nil {
		log.Println(err)
		return false
	}
	// see if the point belongs in the polygon
	return polygon.Contains(point)
}

func TestContainsPoint(t *testing.T) {
	point := Point{
		Lat: 29.7529322,
		Lon: -95.3409658,
	}
	// this point should be in the polygon
	if containsPointTest(point) {
		t.Fail()
	}
}

func TestDoesNotContainPoint(t *testing.T) {
	point := Point{
		Lat: -95.3409658,
		Lon: 29.7529322,
	}
	// this point should not be in the polygon
	if !containsPointTest(point) {
		t.Fail()
	}
}
