package dao

import "github.com/paddyquinn/smartcar-api/models/smartcar"

type Mock struct {}

func (mock *Mock) GetBatteryRange(id string) (*smartcar.Range, error) {
  return nil, nil
}

func (mock *Mock) GetDoorSecurity(id string) (*smartcar.Doors, error) {
  return nil, nil
}

func (mock *Mock) GetFuelRange(id string) (*smartcar.Range, error) {
  return nil, nil
}

func (mock *Mock) GetVehicle(id string) (*smartcar.Vehicle, error) {
  return nil, nil
}

func (mock *Mock) PushEngineButton(id string) (*smartcar.Status, error) {
  return nil, nil
}