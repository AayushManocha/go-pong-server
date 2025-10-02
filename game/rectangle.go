package game

import "fmt"

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
}

func (r Rectangle) toJSON() string {
	return fmt.Sprintf(`{"X": "%d", "Y": "%d"}`, r.X, r.Y)
}

func (r Rectangle) String() string {
	fmt.Sprintf("%d", 5)
	return fmt.Sprintf("Rectangle{X: %d, Y: %d}", r.X, r.Y)
}
