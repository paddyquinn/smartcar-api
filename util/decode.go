package util

import (
  "encoding/json"
  "io"
)

func Decode(body io.ReadCloser, model interface{}) error {
  decoder := json.NewDecoder(body)
  err := decoder.Decode(model)
  if err != nil {
    return err
  }

  return nil
}