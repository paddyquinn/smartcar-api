package util

import (
	"encoding/json"
	"io"
)

// Decode marshals a stream into the model
func Decode(body io.ReadCloser, model interface{}) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(model)
	if err != nil {
		return err
	}

	return nil
}
