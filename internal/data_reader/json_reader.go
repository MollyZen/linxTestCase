package data_reader

import (
	"encoding/json"
	"errors"
	"io"
)

type JsonReader[K any] struct {
	unmarshaler *json.Decoder
	obj         K
}

func NewJSONReader[K any](rd io.Reader) *JsonReader[K] {
	var tmp K
	unmarshaler := json.NewDecoder(rd)
	// skipping first square bracket
	_, _ = unmarshaler.Token()
	jsonReader := &JsonReader[K]{unmarshaler, tmp}
	return jsonReader
}

func (j *JsonReader[K]) ReadNext() (K, error) {
	var errRes K
	// parsing json input
	// reading opening bracket
	var err error
	// reading contents
	if j.unmarshaler.More() {
		if err := j.unmarshaler.Decode(&j.obj); err != nil {
			return errRes, err
		} else {
			return j.obj, nil
		}
	} else {
		// reading closing bracket
		_, err = j.unmarshaler.Token()
		if err != nil {
			return errRes, err
		} else {
			return errRes, errors.New("There are no more objects in JSON array")
		}
	}
}

func (j *JsonReader[K]) More() bool {
	if j.unmarshaler.More() {
		return true
	} else {
		//reading last bracket
		_, _ = j.unmarshaler.Token()
		return false
	}
}
