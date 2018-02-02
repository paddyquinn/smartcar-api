package dao

import (
	"github.com/paddyquinn/smartcar-api/models/smartcar"
	"github.com/stretchr/testify/mock"
)

// Mock represents a mock DAO used for testing
type Mock struct {
	mock.Mock
}

// GetBatteryRange mocks out the call to the GM API for the battery percentage of a vehicle
func (mock *Mock) GetBatteryRange(id string) (smartcar.Model, error) {
	args := mock.Called(id)
	return args.Get(0), args.Error(1)
}

// GetDoorSecurity mocks out the call to the GM API for the security status of each door of a vehicle
func (mock *Mock) GetDoorSecurity(id string) (smartcar.Model, error) {
	args := mock.Called(id)
	return args.Get(0), args.Error(1)
}

// GetFuelRange mocks out the call to the GM API for the fuel percentage of a vehicle
func (mock *Mock) GetFuelRange(id string) (smartcar.Model, error) {
	args := mock.Called(id)
	return args.Get(0), args.Error(1)
}

// GetVehicle mocks out the call to the GM API for information on a vehicle
func (mock *Mock) GetVehicle(id string) (smartcar.Model, error) {
	args := mock.Called(id)
	return args.Get(0), args.Error(1)
}

// PushEngineButton mocks out the call to the GM API to start or stop a vehicle's engine
func (mock *Mock) PushEngineButton(id string, cmd int) (*smartcar.Status, error) {
	args := mock.Called(id, cmd)
	status, ok := args.Get(0).(*smartcar.Status)
	if !ok {
		return nil, args.Error(1)
	}
	return status, args.Error(1)
}
