package models

type Door struct {
  Location string `json:"location"`
  Locked bool `json:"locked"`
}

type Doors []*Door

type Range struct {
  Percent int `json:"percent"`
}

type Status struct {
  Value string `json:"status"`
}

type Vehicle struct {
  VIN string `json:"vin"`
  Color string `json:"color"`
  NumDoors int `json:"doorCount"`
  DriveTrain string `json:"driveTrain"`
}