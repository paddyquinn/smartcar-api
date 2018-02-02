package dao

import "github.com/paddyquinn/smartcar-api/models/smartcar"

// Mock represents a mock DAO used for testing
type Mock struct{}

// GetBatteryRange mocks out the call to the GM API for the battery percentage of a vehicle
func (mock *Mock) GetBatteryRange(id string) (smartcar.Model, error) {
	return nil, nil
}

// GetDoorSecurity mocks out the call to the GM API for the security status of each door of a vehicle
func (mock *Mock) GetDoorSecurity(id string) (smartcar.Model, error) {
	return nil, nil
}

// GetFuelRange mocks out the call to the GM API for the fuel percentage of a vehicle
func (mock *Mock) GetFuelRange(id string) (smartcar.Model, error) {
	return nil, nil
}

// GetVehicle mocks out the call to the GM API for information on a vehicle
func (mock *Mock) GetVehicle(id string) (smartcar.Model, error) {
	return nil, nil
}

// PushEngineButton mocks out the call to the GM API to start or stop a vehicle's engine
func (mock *Mock) PushEngineButton(id string, cmd int) (*smartcar.Status, error) {
	return nil, nil
}
