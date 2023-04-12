package model

type Road struct {
	ID          uint
	Distance    float64
	Floor       int
	NodeStartID int
	NodeEndID   int
}

type Result struct {
	Seq        int     `json:seq`
	Node       int     `json:node`
	Edge       int     `json:edge`
	Cost       float64 `json:cost`
	NodeDetail Node
}
