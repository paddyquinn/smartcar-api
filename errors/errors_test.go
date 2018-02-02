package errors

import (
  "testing"

  . "github.com/smartystreets/goconvey/convey"
)

func TestNewDecodeError(t *testing.T) {
  Convey("Creating a new decode error", t, func() {
    err := NewDecodeError()

    Convey("creates a new error with the correct error message", func() {
      So(err.Error(), ShouldEqual, decodeError)
    })
  })
}

func TestNewJSONMarshalError(t *testing.T) {
  Convey("Creating a new JSON marshal error", t, func() {
    err := NewJSONMarshalError()

    Convey("creates a new error with the correct error message", func() {
      So(err.Error(), ShouldEqual, jsonMarshalError)
    })
  })
}

func TestNewPostError(t *testing.T) {
  Convey("Creating a new POST error", t, func() {
    err := NewPostError()

    Convey("creates a new error with the correct error message", func() {
      So(err.Error(), ShouldEqual, postError)
    })
  })
}

func TestNewUnidentifiedCommandError(t *testing.T) {
  Convey("Creating a new unidentified command error", t, func() {
    err := NewUnidentifiedCommandError()

    Convey("creates a new error with the correct error message", func() {
      So(err.Error(), ShouldEqual, unidentifiedCommandError)
    })
  })
}
