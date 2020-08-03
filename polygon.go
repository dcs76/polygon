package polygon

import (
	"fmt"
	"log"
)

var ErrNoData = fmt.Errorf("polygon: no data")

// Polygon represents and enclosed area
type Polygon struct {
	data []Point
	// these store the min/max of our points
	minLat, maxLat, minLon, maxLon float64
	isMinMaxComputed               bool
}

// Point represent a set of coordinates
type Point struct {
	Lat, Lon float64
}

// NewPolygonAsPoints retuns a Polygon given a variadic list of points.
// If the list is even, the method will automatically close the polygon
// by using the starting point as the last point.
func NewPolygonAsPoints(points ...Point) (*Polygon, error) {
	var mustClose bool
	polygon := Polygon{}

	if len(points) == 0 {
		return nil, ErrNoData
	}

	if len(points)%2 == 0 {
		mustClose = true
	}

	for _, p := range points {
		polygon.data = append(polygon.data, p)
	}

	if mustClose {
		polygon.data = append(polygon.data, polygon.data[0])
	}

	return &polygon, nil
}

// NewPolygonAsSlice retuns a Polygon given a slice of points.
// If slice is even, the method will automatically close the polygon
// by using the starting point as the last point.
func NewPolygonAsSlice(points []Point) (*Polygon, error) {
	return NewPolygonAsPoints(points...)
}

// Contains returns true if the given point lies within the boundary
// of the polygon.
// Based on: https://stackoverflow.com/questions/217578/how-can-i-determine-whether-a-2d-point-is-within-a-polygon
func (p *Polygon) Contains(point Point) bool {
	if p.data == nil {
		return false
	}
	log.Println(len(p.data))
	// this will effciently determine if a point is NOT inside the polygon
	if !p.isMinMaxComputed {
		p.computeMinMax()
	}
	if point.Lat < p.minLat || point.Lat > p.maxLat || point.Lon < p.minLon || point.Lon > p.maxLon {
		return false
	}
	// now check if the point is inside
	var inside bool
	for i, j := 0, len(p.data)-1; i < len(p.data) && j < len(p.data); j = i {
		i++
		if i == len(p.data) {
			break
		}
		if j == len(p.data) {
			break
		}
		if (p.data[i].Lon > point.Lon) != (p.data[j].Lon > point.Lon) &&
			point.Lat < (p.data[j].Lat-p.data[i].Lat)*(point.Lon-p.data[i].Lon)/
				(p.data[j].Lon-p.data[i].Lon)+p.data[i].Lat {
			inside = !inside
			break
		}
	}

	return inside
}

func (p *Polygon) computeMinMax() {
	p.isMinMaxComputed = true
	for i, point := range p.data {
		if i == 0 {
			p.minLat, p.minLon = point.Lat, point.Lon
			p.maxLat, p.maxLon = point.Lat, point.Lon
			continue
		}
		if p.minLat > point.Lat {
			p.minLat = point.Lat
		}
		if p.minLon > point.Lon {
			p.minLon = point.Lon
		}
		if p.maxLat < point.Lat {
			p.maxLat = point.Lat
		}
		if p.maxLon < point.Lon {
			p.maxLon = point.Lon
		}
	}
}
