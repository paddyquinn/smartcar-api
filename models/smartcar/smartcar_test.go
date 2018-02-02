package smartcar

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEngineRequest_ToEnum(t *testing.T) {
	Convey("Given an engine request with the start action", t, func() {
		engineRequest := EngineRequest{Action: start}

		Convey("a call to ToEnum returns the start enum value", func() {
			So(engineRequest.ToEnum(), ShouldEqual, Start)
		})
	})

	Convey("Given an engine request with the stop action", t, func() {
		engineRequest := EngineRequest{Action: stop}

		Convey("a call to ToEnum returns the stop enum value", func() {
			So(engineRequest.ToEnum(), ShouldEqual, Stop)
		})
	})

	Convey("Given an engine request with an unknown action", t, func() {
		engineRequest := EngineRequest{Action: "UNKNOWN"}

		Convey("a call to ToEnum returns the no op enum value", func() {
			So(engineRequest.ToEnum(), ShouldEqual, NoOp)
		})
	})
}

func TestStatus_IsSuccess(t *testing.T) {
	Convey("Given a success status", t, func() {
		status := &Status{Value: "success"}

		Convey("IsSuccess should return true", func() {
			So(status.IsSuccess(), ShouldBeTrue)
		})
	})

	Convey("Given an error status", t, func() {
		status := &Status{Value: "error"}

		Convey("IsSuccess should return false", func() {
			So(status.IsSuccess(), ShouldBeFalse)
		})
	})
}
