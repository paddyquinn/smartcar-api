package dao

import "github.com/paddyquinn/smartcar-api/models/smartcar"

type Interface interface {
  GetBatteryRange(id string) (*smartcar.Range, error)
  GetDoorSecurity(id string) (*smartcar.Doors, error)
  GetFuelRange(id string) (*smartcar.Range, error)
  GetVehicle(id string) (*smartcar.Vehicle, error)
  PushEngineButton(id string) (*smartcar.Status, error)
}