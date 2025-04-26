package zipcode

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

type Metric string

const (
	Kilometers = Metric("K")
	Miles      = Metric("N")
)

func New() (a *Engine) {
	a = new(Engine)
	a.init()
	return
}

type Engine struct {
	db map[int]*Code
	sync.RWMutex
}

func (a *Engine) init() {
	a.db = make(map[int]*Code)
	for _, x := range list {
		id, _ := strconv.Atoi(x.Code)
		if id == 0 {
			continue
		}
		a.db[id] = &x
	}
}

func (a *Engine) GetString(code string) (res *Code) {
	id, _ := strconv.Atoi(code)
	if id == 0 {
		return
	}
	return a.Get(id)
}

func (a *Engine) Get(code int) (res *Code) {
	a.RLock()
	res = a.db[code]
	a.RUnlock()
	return
}

// pseudo drive distance (factor 1.5)
func (a *Engine) DriveMiles(from, to int) (miles int) {
	res := a.DirectMiles(from, to)
	if res == 0 {
		return
	}

	miles = int(float64(res) * 1.5)
	return
}

// pseudo drive distance (factor 1.5)
func (a *Engine) DriveMilesWithFactor(from, to int, factor float64) (miles int) {
	res := a.DirectMiles(from, to)
	if res == 0 {
		return
	}

	miles = int(float64(res) * factor)
	return
}

// A > B direct miles on map (without roads etc)
func (a *Engine) DirectMiles(from, to int) (miles int) {
	a.RLock()
	p1 := a.db[from]
	if p1 == nil {
		return
	}
	p2 := a.db[to]
	if p2 == nil {
		return
	}

	return int(Distance(p1.Lat, p1.Long, p2.Lat, p2.Long, Miles))
}

// A > B direct miles on map (without roads etc)
func (a *Engine) DirectKilimeters(from, to int) (km int) {
	a.RLock()
	p1 := a.db[from]
	if p1 == nil {
		return
	}
	p2 := a.db[to]
	if p2 == nil {
		return
	}

	return int(Distance(p1.Lat, p1.Long, p2.Lat, p2.Long, Kilometers))

}

// A > B drive km on map (without roads etc)
func (a *Engine) DriveKilimeters(from, to int) (res int) {
	r := a.DirectKilimeters(from, to)
	res = int(float64(r) * 1.5)
	return
}

// math distance
func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, metric Metric) float64 {
	fmt.Println(lat1, lng1, lat2, lng2)
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	switch metric {
	case Kilometers:
		dist = dist * 1.609344
	case Miles:
		dist = dist * 0.8684
	}

	return dist
}
