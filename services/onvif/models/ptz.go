package models

type PtzAxes struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type PtzControl struct {
	Name string  `json:"name"`
	Axes PtzAxes `json:"axes"`
}

type PtzStatusData struct {
	X string `json:"x"`
	Y string `json:"y"`
	Z string `json:"z"`
}
