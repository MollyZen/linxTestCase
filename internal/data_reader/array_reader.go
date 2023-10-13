package data_reader

import "errors"

type ArrayObjectReader[K any] struct {
	pnt int
	arr []K
}

func NewArrayObjectReader[K any](arr []K) *ArrayObjectReader[K] {
	return &ArrayObjectReader[K]{pnt: -1, arr: arr}
}

func (a *ArrayObjectReader[K]) ReadNext() (K, error) {
	if a.More() {
		a.pnt++
		return a.arr[a.pnt], nil
	}
	var errV K
	return errV, errors.New("no more data")
}

func (a *ArrayObjectReader[K]) More() bool {
	return a.pnt+1 < len(a.arr)
}
