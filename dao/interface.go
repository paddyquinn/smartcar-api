package dao

import "github.com/paddyquinn/smartcar-api/models/smartcar"

// Interface is an interface that a DAO must implement (used so a mock DAO can be used in the handler for testing)
type Interface interface {
	GetBatteryRange(id string) (smartcar.Model, error)
	GetDoorSecurity(id string) (smartcar.Model, error)
	GetFuelRange(id string) (smartcar.Model, error)
	GetVehicle(id string) (smartcar.Model, error)
	PushEngineButton(id string, cmd int) (*smartcar.Status, error)
}
