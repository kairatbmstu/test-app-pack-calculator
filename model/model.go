package model

type Pack struct {
	Size     int `json:"size"`
	Num      int `json:"num"`
	DpCoords DpCoords
}

type DpCoords struct {
	I int
	J int
}
