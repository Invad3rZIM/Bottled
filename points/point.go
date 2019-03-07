package points

import (
	"math"
)

const (
	pi float64 = .017453292519943295 //pi / 100
)

type Point struct {
	BottleID int

	Enabled bool

	Lat  float64
	Long float64

	Distance float64
	Age      int
}

func NewPoint(id int, bid int, enabled bool, lat float64, long float64) *Point {
	p := Point{
		BottleID: bid,
		Enabled:  enabled,
		Lat:      lat,
		Long:     long,
	}

	return &p
}

/* https://stackoverflow.com/questions/27928/calculate-distance-between-two-latitude-longitude-points-haversine-formula
function distance(lat1, lon1, lat2, lon2) {
  var p = 0.017453292519943295;    // Math.PI / 180
  var c = Math.cos;
  var a = 0.5 - c((lat2 - lat1) * p)/2 +
          c(lat1 * p) * c(lat2 * p) *
          (1 - c((lon2 - lon1) * p))/2;

  return 12742 * Math.asin(Math.sqrt(a)); // 2 * R; R = 6371 km
}
*/

func (p Point) CalcDistance(q Point) float64 {
	a := 0.5 - math.Cos((q.Lat-p.Lat)*pi)/2 + math.Cos(p.Lat*pi)*math.Cos(q.Lat*pi)*(1-math.Cos((q.Long-p.Long)*pi))/2
	a = 12742 * math.Asin(math.Sqrt(a))

	return a
}
