package zipcode

type Code struct {
	Code   string  `json:"code,omitempty"`
	Lat    float64 `json:"lat,omitempty"`
	Long   float64 `json:"long,omitempty"`
	City   string  `json:"city,omitempty"`
	State  string  `json:"state,omitempty"`
	County string  `json:"county,omitempty"`
}
