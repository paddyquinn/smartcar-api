package dao

import "github.com/paddyquinn/smartcar-api/models/smartcar"

type Interface interface {
  GetBatteryRange(id string) (smartcar.Model, error)
  GetDoorSecurity(id string) (smartcar.Model, error)
  GetFuelRange(id string) (smartcar.Model, error)
  GetVehicle(id string) (smartcar.Model, error)
  PushEngineButton(id string, cmd int) (*smartcar.Status, error)
}