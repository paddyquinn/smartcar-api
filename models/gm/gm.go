package gm

import (
  "strconv"

  "github.com/paddyquinn/smartcar-api/models/smartcar"
)

type Model interface {
  ToSmartcar() smartcar.Model
}

type BatteryLevelData struct {
  BatteryLevel *Value `json:"batteryLevel"`
}

type BatteryRange struct {
  Data *BatteryLevelData `json:"data"`
}

func (batteryRange *BatteryRange) ToSmartcar() smartcar.Model {
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

type DoorsData struct {
  Doors *Values `json:"doors"`
}

type DoorsResponse struct {
  Data *DoorsData `json:"data"`
}

func (doorsResponse *DoorsResponse) ToSmartcar() smartcar.Model {
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
      Locked: isLocked,
    }
  }

  return doors
}

type FuelLevelData struct {
  FuelLevel *Value `json:"tankLevel"`
}

type FuelRange struct {
  Data *FuelLevelData `json:"data"`
}

func (fuelRange *FuelRange) ToSmartcar() smartcar.Model {
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

type LocationLocked struct {
  Location *Value `json:"location"`
  Locked *Value `json:"locked"`
}

type RequestBody struct {
  ID string `json:"id"`
  Command string `json:"command,omitempty"`
  ResponseType string `json:"responseType"`
}

type Value struct {
  Value string `json:"value"`
}

type Values struct {
  Values []*LocationLocked `json:"values"`
}

type Vehicle struct {
  Data *VehicleData `json:"data"`
}

func (vehicle *Vehicle) ToSmartcar() smartcar.Model {
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
    VIN: data.VIN.Value,
    Color: data.Color.Value,
    NumDoors: numDoors,
    DriveTrain: data.DriveTrain.Value,
  }
}

type VehicleData struct {
  VIN *Value `json:"vin"`
  Color *Value `json:"color"`
  FourDoorSedan *Value `json:"fourDoorSedan"`
  TwoDoorCoupe *Value `json:"twoDoorCoupe"`
  DriveTrain *Value `json:"driveTrain"`
}