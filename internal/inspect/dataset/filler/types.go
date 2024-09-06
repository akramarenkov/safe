package filler

// Fill arguments of dataset item by values.
type Filler interface {
	Fill(args []int64) (bool, error)
}
