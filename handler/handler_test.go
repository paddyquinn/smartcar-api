package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paddyquinn/smartcar-api/dao"
	"github.com/paddyquinn/smartcar-api/errors"
	"github.com/paddyquinn/smartcar-api/models/smartcar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandler_GetBatteryRange(t *testing.T) {
	Convey("Given a handler with calls to the GM API mocked out and a call to the battery endpoint", t, func() {
		mock := &dao.Mock{}
		hdlr := &Handler{DAO: mock}
		router := hdlr.SetUpRouter()

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/vehicles/1/battery", nil)

		Convey("if the DAO returns a JSONMarshalError", func() {
			mock.On("GetBatteryRange", "1").Return(nil, errors.NewJSONMarshalError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to marshal the request body JSON\"}")
			})
		})

		Convey("if the DAO returns a PostError", func() {
			mock.On("GetBatteryRange", "1").Return(nil, errors.NewPostError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"POST request to GM API failed\"}")
			})
		})

		Convey("if the DAO returns a DecodeError", func() {
			mock.On("GetBatteryRange", "1").Return(nil, errors.NewDecodeError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to decode the response from the GM API\"}")
			})
		})

		Convey("if the DAO returns a BatteryRange", func() {
			mock.On("GetBatteryRange", "1").Return(&smartcar.Range{Percent: 28.4}, nil)
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 200 with the appropriate response", func() {
				So(recorder.Code, ShouldEqual, http.StatusOK)
				So(recorder.Body.String(), ShouldEqual, "{\"percent\":28.4}")
			})
		})
	})
}

func TestHandler_GetDoorSecurity(t *testing.T) {
	Convey("Given a handler with calls to the GM API mocked out and a call to the doors endpoint", t, func() {
		mock := &dao.Mock{}
		hdlr := &Handler{DAO: mock}
		router := hdlr.SetUpRouter()

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/vehicles/1/doors", nil)

		Convey("if the DAO returns a JSONMarshalError", func() {
			mock.On("GetDoorSecurity", "1").Return(nil, errors.NewJSONMarshalError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to marshal the request body JSON\"}")
			})
		})

		Convey("if the DAO returns a PostError", func() {
			mock.On("GetDoorSecurity", "1").Return(nil, errors.NewPostError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"POST request to GM API failed\"}")
			})
		})

		Convey("if the DAO returns a DecodeError", func() {
			mock.On("GetDoorSecurity", "1").Return(nil, errors.NewDecodeError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to decode the response from the GM API\"}")
			})
		})

		Convey("if the DAO returns a Doors object", func() {
			mock.On("GetDoorSecurity", "1").Return(smartcar.Doors{{Location: "frontLeft", Locked: true}, {Location: "frontRight", Locked: false}}, nil)
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 200 with the appropriate response", func() {
				So(recorder.Code, ShouldEqual, http.StatusOK)
				So(recorder.Body.String(), ShouldEqual, "[{\"location\":\"frontLeft\",\"locked\":true},{\"location\":\"frontRight\",\"locked\":false}]")
			})
		})
	})
}

func TestHandler_GetFuelRange(t *testing.T) {
	Convey("Given a handler with calls to the GM API mocked out and a call to the fuel endpoint", t, func() {
		mock := &dao.Mock{}
		hdlr := &Handler{DAO: mock}
		router := hdlr.SetUpRouter()

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/vehicles/1/fuel", nil)

		Convey("if the DAO returns a JSONMarshalError", func() {
			mock.On("GetFuelRange", "1").Return(nil, errors.NewJSONMarshalError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to marshal the request body JSON\"}")
			})
		})

		Convey("if the DAO returns a PostError", func() {
			mock.On("GetFuelRange", "1").Return(nil, errors.NewPostError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"POST request to GM API failed\"}")
			})
		})

		Convey("if the DAO returns a DecodeError", func() {
			mock.On("GetFuelRange", "1").Return(nil, errors.NewDecodeError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to decode the response from the GM API\"}")
			})
		})

		Convey("if the DAO returns a FuelRange", func() {
			mock.On("GetFuelRange", "1").Return(&smartcar.Range{Percent: 77.7}, nil)
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 200 with the appropriate response", func() {
				So(recorder.Code, ShouldEqual, http.StatusOK)
				So(recorder.Body.String(), ShouldEqual, "{\"percent\":77.7}")
			})
		})
	})
}

func TestHandler_GetVehicle(t *testing.T) {
	Convey("Given a handler with calls to the GM API mocked out and a call to the vehicle endpoint", t, func() {
		mock := &dao.Mock{}
		hdlr := &Handler{DAO: mock}
		router := hdlr.SetUpRouter()

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/vehicles/1", nil)

		Convey("if the DAO returns a JSONMarshalError", func() {
			mock.On("GetVehicle", "1").Return(nil, errors.NewJSONMarshalError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to marshal the request body JSON\"}")
			})
		})

		Convey("if the DAO returns a PostError", func() {
			mock.On("GetVehicle", "1").Return(nil, errors.NewPostError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"POST request to GM API failed\"}")
			})
		})

		Convey("if the DAO returns a DecodeError", func() {
			mock.On("GetVehicle", "1").Return(nil, errors.NewDecodeError())
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 500 with the appropriate error message", func() {
				So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
				So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to decode the response from the GM API\"}")
			})
		})

		Convey("if the DAO returns a Vehicle", func() {
			mock.On("GetVehicle", "1").Return(&smartcar.Vehicle{VIN: "test-vin", Color: "test-color", NumDoors: 2, DriveTrain: "test-drivetrain"}, nil)
			router.ServeHTTP(recorder, req)

			Convey("the status code should be 200 with the appropriate response", func() {
				So(recorder.Code, ShouldEqual, http.StatusOK)
				So(recorder.Body.String(), ShouldEqual, "{\"vin\":\"test-vin\",\"color\":\"test-color\",\"doorCount\":2,\"driveTrain\":\"test-drivetrain\"}")
			})
		})
	})
}

func TestHandler_PushEngineButton(t *testing.T) {
	Convey("Given a handler with calls to the GM API mocked out and a call to the engine endpoint", t, func() {
		mock := &dao.Mock{}
		hdlr := &Handler{DAO: mock}
		router := hdlr.SetUpRouter()

		recorder := httptest.NewRecorder()
		Convey("with a request to start the engine", func() {
			buffer, _ := json.Marshal(&smartcar.EngineRequest{Action: "START"})
			reader := bytes.NewReader(buffer)
			req := httptest.NewRequest("POST", "/vehicles/1/engine", reader)

			Convey("if the DAO returns a JSONMarshalError", func() {
				mock.On("PushEngineButton", "1", smartcar.Start).Return(nil, errors.NewJSONMarshalError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to marshal the request body JSON\"}")
				})
			})

			Convey("if the DAO returns a PostError", func() {
				mock.On("PushEngineButton", "1", smartcar.Start).Return(nil, errors.NewPostError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"POST request to GM API failed\"}")
				})
			})

			Convey("if the DAO returns a DecodeError", func() {
				mock.On("PushEngineButton", "1", smartcar.Start).Return(nil, errors.NewDecodeError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to decode the response from the GM API\"}")
				})
			})

			Convey("if the vehicle is successfully started", func() {
				mock.On("PushEngineButton", "1", smartcar.Start).Return(&smartcar.Status{Value: "success"}, nil)
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 200 with the appropriate response", func() {
					So(recorder.Code, ShouldEqual, http.StatusOK)
					So(recorder.Body.String(), ShouldEqual, "{\"status\":\"success\"}")
				})
			})

			Convey("if the vehicle fails to start", func() {
				mock.On("PushEngineButton", "1", smartcar.Start).Return(&smartcar.Status{Value: "error"}, nil)
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate response", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"status\":\"error\"}")
				})
			})
		})

		Convey("with a request to stop the engine", func() {
			buffer, _ := json.Marshal(&smartcar.EngineRequest{Action: "STOP"})
			reader := bytes.NewReader(buffer)
			req := httptest.NewRequest("POST", "/vehicles/1/engine", reader)

			Convey("if the DAO returns a JSONMarshalError", func() {
				mock.On("PushEngineButton", "1", smartcar.Stop).Return(nil, errors.NewJSONMarshalError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to marshal the request body JSON\"}")
				})
			})

			Convey("if the DAO returns a PostError", func() {
				mock.On("PushEngineButton", "1", smartcar.Stop).Return(nil, errors.NewPostError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"POST request to GM API failed\"}")
				})
			})

			Convey("if the DAO returns a DecodeError", func() {
				mock.On("PushEngineButton", "1", smartcar.Stop).Return(nil, errors.NewDecodeError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"failed to decode the response from the GM API\"}")
				})
			})

			Convey("if the vehicle is successfully stopped", func() {
				mock.On("PushEngineButton", "1", smartcar.Stop).Return(&smartcar.Status{Value: "success"}, nil)
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 200 with the appropriate response", func() {
					So(recorder.Code, ShouldEqual, http.StatusOK)
					So(recorder.Body.String(), ShouldEqual, "{\"status\":\"success\"}")
				})
			})

			Convey("if the vehicle fails to stop", func() {
				mock.On("PushEngineButton", "1", smartcar.Stop).Return(&smartcar.Status{Value: "error"}, nil)
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 500 with the appropriate response", func() {
					So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
					So(recorder.Body.String(), ShouldEqual, "{\"status\":\"error\"}")
				})
			})
		})

		Convey("with an unknown request", func() {
			buffer, _ := json.Marshal(&smartcar.EngineRequest{Action: "UNIDENTIFIED"})
			reader := bytes.NewReader(buffer)
			req := httptest.NewRequest("POST", "/vehicles/1/engine", reader)

			Convey("the DAO will return an UnidentifiedCommandError", func() {
				mock.On("PushEngineButton", "1", smartcar.NoOp).Return(nil, errors.NewUnidentifiedCommandError())
				router.ServeHTTP(recorder, req)

				Convey("the status code should be 400 with the appropriate error message", func() {
					So(recorder.Code, ShouldEqual, http.StatusBadRequest)
					So(recorder.Body.String(), ShouldEqual, "{\"error\":\"could not identify engine command\"}")
				})
			})
		})
	})
}
