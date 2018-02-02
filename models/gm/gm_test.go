package gm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActionResult_ToSmartcar(t *testing.T) {
	Convey("When transforming a GM action result to a smartcar status", t, func() {
		Convey("if the status object is nil", func() {
			ar := &ActionResult{}
			smartcarStatus := ar.ToSmartcar()

			Convey("nil is returned", func() {
				So(smartcarStatus, ShouldBeNil)
			})
		})

		Convey("if the status object is not nil", func() {
			Convey("and the status is not EXECUTED", func() {
				ar := &ActionResult{
					Status: &Status{
						Value: "FAILED",
					},
				}
				smartcarStatus := ar.ToSmartcar()

				Convey("an error status is returned", func() {
					So(smartcarStatus, ShouldNotBeNil)
					So(smartcarStatus.Value, ShouldEqual, err)
				})
			})

			Convey("and the status is EXECUTED", func() {
				ar := &ActionResult{
					Status: &Status{
						Value: executed,
					},
				}
				smartcarStatus := ar.ToSmartcar()

				Convey("a success status is returned", func() {
					So(smartcarStatus, ShouldNotBeNil)
					So(smartcarStatus.Value, ShouldEqual, success)
				})
			})
		})
	})
}

func TestBatteryRange_ToSmartcar(t *testing.T) {
	Convey("When transforming a GM battery range to a smartcar range", t, func() {
		Convey("if the data object is nil", func() {
			br := &BatteryRange{}
			smartcarRange := br.ToSmartcar()

			Convey("nil is returned", func() {
				So(smartcarRange, ShouldBeNil)
			})
		})

		Convey("if the data object is not nil", func() {
			Convey("and the value is null", func() {
				br := &BatteryRange{
					Data: &BatteryLevelData{
						BatteryLevel: &Value{
							Value: "null",
						},
					},
				}
				smartcarRange := br.ToSmartcar()

				Convey("the range defaults to 0", func() {
					So(smartcarRange, ShouldNotBeNil)
					So(smartcarRange.Percent, ShouldEqual, 0)
				})
			})

			Convey("and the value is not null", func() {
				br := &BatteryRange{
					Data: &BatteryLevelData{
						BatteryLevel: &Value{
							Value: "53.9",
						},
					},
				}
				smartcarRange := br.ToSmartcar()

				Convey("the range is successfully changed to a floating point number", func() {
					So(smartcarRange, ShouldNotBeNil)
					So(smartcarRange.Percent, ShouldEqual, 53.9)
				})
			})
		})
	})
}

func TestDoorsResponse_ToSmartcar(t *testing.T) {
	Convey("When transforming a GM security status response to a smartcar doors object", t, func() {
		Convey("if the data object is nil", func() {
			dr := &DoorsResponse{}
			smartcarDoors := dr.ToSmartcar()

			Convey("nil is returned", func() {
				So(smartcarDoors, ShouldBeNil)
			})
		})

		Convey("if the data object is not nil", func() {
			dr := &DoorsResponse{
				Data: &DoorsData{
					Doors: &Values{
						Values: []*LocationLocked{
							{
								Location: &Value{
									Value: "test-location1",
								},
								Locked: &Value{
									Value: "True",
								},
							},
							{
								Location: &Value{
									Value: "test-location2",
								},
								Locked: &Value{
									Value: "False",
								},
							},
						},
					},
				},
			}
			smartcarDoors := dr.ToSmartcar()

			Convey("the data is successfully transformed", func() {
				So(smartcarDoors, ShouldNotBeNil)
				So(len(smartcarDoors), ShouldEqual, 2)
				So(smartcarDoors[0].Location, ShouldEqual, "test-location1")
				So(smartcarDoors[0].Locked, ShouldBeTrue)
				So(smartcarDoors[1].Location, ShouldEqual, "test-location2")
				So(smartcarDoors[1].Locked, ShouldBeFalse)
			})
		})
	})
}

func TestFuelRange_ToSmartcar(t *testing.T) {
	Convey("When transforming a GM fuel range to a smartcar range", t, func() {
		Convey("if the data object is nil", func() {
			fr := &FuelRange{}
			smartcarRange := fr.ToSmartcar()

			Convey("nil is returned", func() {
				So(smartcarRange, ShouldBeNil)
			})
		})

		Convey("if the data object is not nil", func() {
			Convey("and the value is null", func() {
				fr := &FuelRange{
					Data: &FuelLevelData{
						FuelLevel: &Value{
							Value: "null",
						},
					},
				}
				smartcarRange := fr.ToSmartcar()

				Convey("the range defaults to 0", func() {
					So(smartcarRange, ShouldNotBeNil)
					So(smartcarRange.Percent, ShouldEqual, 0)
				})
			})

			Convey("and the value is not null", func() {
				fr := &FuelRange{
					Data: &FuelLevelData{
						FuelLevel: &Value{
							Value: "37.6",
						},
					},
				}
				smartcarRange := fr.ToSmartcar()

				Convey("the range is successfully changed to a floating point number", func() {
					So(smartcarRange, ShouldNotBeNil)
					So(smartcarRange.Percent, ShouldEqual, 37.6)
				})
			})
		})
	})
}

func TestVehicle_ToSmartcar(t *testing.T) {
	Convey("When transforming a GM vehicle to a smartcar vehicle", t, func() {
		Convey("if the data object is nil", func() {
			vehicle := &Vehicle{}
			smartcarVehicle := vehicle.ToSmartcar()

			Convey("nil is returned", func() {
				So(smartcarVehicle, ShouldBeNil)
			})
		})

		Convey("if the data object is not nil", func() {
			Convey("and the vehicle is a two door coupe", func() {
				vehicle := &Vehicle{
					Data: &VehicleData{
						VIN: &Value{
							Value: "test-vin",
						},
						Color: &Value{
							Value: "test-color",
						},
						TwoDoorCoupe: &Value{
							Value: "True",
						},
						FourDoorSedan: &Value{
							Value: "False",
						},
						DriveTrain: &Value{
							Value: "test-drivetrain",
						},
					},
				}
				smartcarVehicle := vehicle.ToSmartcar()

				Convey("the data is successfully transformed to smartcar data", func() {
					So(smartcarVehicle, ShouldNotBeNil)
					So(smartcarVehicle.VIN, ShouldEqual, "test-vin")
					So(smartcarVehicle.Color, ShouldEqual, "test-color")
					So(smartcarVehicle.NumDoors, ShouldEqual, 2)
					So(smartcarVehicle.DriveTrain, ShouldEqual, "test-drivetrain")
				})
			})

			Convey("and the vehicle is a four door sedan", func() {
				vehicle := &Vehicle{
					Data: &VehicleData{
						VIN: &Value{
							Value: "test-vin",
						},
						Color: &Value{
							Value: "test-color",
						},
						TwoDoorCoupe: &Value{
							Value: "False",
						},
						FourDoorSedan: &Value{
							Value: "True",
						},
						DriveTrain: &Value{
							Value: "test-drivetrain",
						},
					},
				}
				smartcarVehicle := vehicle.ToSmartcar()

				Convey("the data is successfully transformed to smartcar data", func() {
					So(smartcarVehicle, ShouldNotBeNil)
					So(smartcarVehicle.VIN, ShouldEqual, "test-vin")
					So(smartcarVehicle.Color, ShouldEqual, "test-color")
					So(smartcarVehicle.NumDoors, ShouldEqual, 4)
					So(smartcarVehicle.DriveTrain, ShouldEqual, "test-drivetrain")
				})
			})
		})
	})
}
