package gm

import "github.com/paddyquinn/smartcar-api/models/smartcar"

const trueCapitalized = "True"

type RequestBody struct {
  ID string `json:"id"`
  Command string `json:"command,omitempty"`
  ResponseType string `json:"responseType"`
}

type Value struct {
  Value string `json:"value"`
}

type Vehicle struct {
  Data *VehicleData `json:"data"`
}

func (vehicle *Vehicle) ToSmartcar() *smartcar.Vehicle {
  data := vehicle.Data
  if data == nil {
    return nil
  }

  var numDoors int

  if data.TwoDoorCoupe.Value == trueCapitalized {
    numDoors = 2
  }

  if data.FourDoorSedan.Value == trueCapitalized {
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