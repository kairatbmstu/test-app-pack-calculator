// Package model defines data structures related to pack calculations.
package model

// Pack represents a pack with its size, number, and dynamic programming coordinates.
type Pack struct {
	Size     int      `json:"size"`
	Num      int      `json:"num"`
	DpCoords DpCoords `json:"-"`
}

// DpCoords represents the coordinates used in dynamic programming.
type DpCoords struct {
	I int `json:"-"` // Vertical coordinate
	J int `json:"-"` // Horizontal coordinate
}
