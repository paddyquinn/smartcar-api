package dao

import "github.com/paddyquinn/smartcar-api/models"

type Interface interface {
  GetBatteryRange(id int) *models.Range
  GetDoorSecurity(id int) *models.Doors
  GetFuelRange(id int) *models.Range
  GetVehicle(id int) *models.Vehicle
  PushEngineButton(id int) *models.Status
}