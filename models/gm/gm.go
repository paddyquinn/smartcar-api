package gm

import (
	"strconv"

	"github.com/paddyquinn/smartcar-api/models/smartcar"
)

const (
	err      = "error"
	executed = "EXECUTED"
	success  = "success"
)

// ActionResult is the top level response to GM's action engine service
type ActionResult struct {
	Status *Status `json:"actionResult"`
}

// ToSmartcar transforms an ActionResult to a smartcar compliant response
func (ar *ActionResult) ToSmartcar() *smartcar.Status {
	arStatus := ar.Status
	if ar.Status == nil {
		return nil
	}

	status := err
	if arStatus.Value == executed {
		status = success
	}

	return &smartcar.Status{
		Value: status,
	}
}

// BatteryLevelData is the data object nested within GM's get energy service response
type BatteryLevelData struct {
	BatteryLevel *Value `json:"batteryLevel"`
}

// BatteryRange is the top level response to GM's get energy service
type BatteryRange struct {
	Data *BatteryLevelData `json:"data"`
}

// ToSmartcar transforms a BatteryRange into a smartcar compliant response
func (batteryRange *BatteryRange) ToSmartcar() *smartcar.Range {
	data := batteryRange.Data
	if data == nil {
		return nil
	}

	// Ignore error (occurs when string is "null"), defaults to 0
	percent, _ := strconv.ParseFloat(data.BatteryLevel.Value, 64)
	return &smartcar.Range{
		Percent: percent,
	}
}

// DoorsData is the data object nested within GM's security status service response
type DoorsData struct {
	Doors *Values `json:"doors"`
}

// DoorsResponse is the top level response to GM's security status service
type DoorsResponse struct {
	Data *DoorsData `json:"data"`
}

// ToSmartcar transforms a DoorsResponse into a smartcar compliant response
func (doorsResponse *DoorsResponse) ToSmartcar() smartcar.Doors {
	data := doorsResponse.Data
	if data == nil {
		return nil
	}

	doors := make(smartcar.Doors, len(data.Doors.Values))
	for idx, door := range data.Doors.Values {
		// Ignore error (occurs when string value is not accepted true/false format), defaults to false
		isLocked, _ := strconv.ParseBool(door.Locked.Value)
		doors[idx] = &smartcar.Door{
			Location: door.Location.Value,
			Locked:   isLocked,
		}
	}

	return doors
}

// FuelLevelData is the data object nested within GM's get energy service response
type FuelLevelData struct {
	FuelLevel *Value `json:"tankLevel"`
}

// FuelRange is the top level response to GM's get energy service
type FuelRange struct {
	Data *FuelLevelData `json:"data"`
}

// ToSmartcar transforms a FuelRange into a smartcar compliant response
func (fuelRange *FuelRange) ToSmartcar() *smartcar.Range {
	data := fuelRange.Data
	if data == nil {
		return nil
	}

	// Ignore error (occurs when string is "null"), defaults to 0
	percent, _ := strconv.ParseFloat(data.FuelLevel.Value, 64)
	return &smartcar.Range{
		Percent: percent,
	}
}

// LocationLocked contains the data about each door's security status
type LocationLocked struct {
	Location *Value `json:"location"`
	Locked   *Value `json:"locked"`
}

// RequestBody represents the body of the POST requests sent to GM
type RequestBody struct {
	ID           string `json:"id"`
	Command      string `json:"command,omitempty"`
	ResponseType string `json:"responseType"`
}

// Status is just a string value from GM's action engine service (should always be EXECUTED or FAILED)
type Status struct {
	Value string `json:"status"`
}

// Value is just a string value, necessary to marshal GM's responses into Go structs
type Value struct {
	Value string `json:"value"`
}

// Values is an array of information about the doors of a vehicle
type Values struct {
	Values []*LocationLocked `json:"values"`
}

// Vehicle is the top level response to GM's get vehicle service
type Vehicle struct {
	Data *VehicleData `json:"data"`
}

// ToSmartcar transforms a Vehicle into a smartcar compliant response
func (vehicle *Vehicle) ToSmartcar() *smartcar.Vehicle {
	data := vehicle.Data
	if data == nil {
		return nil
	}

	var numDoors int

	// Ignore error (occurs when string value is not accepted true/false format), defaults to false
	isTwoDoorCoupe, _ := strconv.ParseBool(data.TwoDoorCoupe.Value)
	if isTwoDoorCoupe {
		numDoors = 2
	}

	// Ignore error (occurs when string value is not accepted true/false format), defaults to false
	isFourDoorSedan, _ := strconv.ParseBool(data.FourDoorSedan.Value)
	if isFourDoorSedan {
		numDoors = 4
	}

	return &smartcar.Vehicle{
		VIN:        data.VIN.Value,
		Color:      data.Color.Value,
		NumDoors:   numDoors,
		DriveTrain: data.DriveTrain.Value,
	}
}

// VehicleData is the data object nested within GM's vehicle service response
type VehicleData struct {
	VIN           *Value `json:"vin"`
	Color         *Value `json:"color"`
	FourDoorSedan *Value `json:"fourDoorSedan"`
	TwoDoorCoupe  *Value `json:"twoDoorCoupe"`
	DriveTrain    *Value `json:"driveTrain"`
}
