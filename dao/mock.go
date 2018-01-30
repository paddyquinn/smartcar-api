package dao

import "github.com/paddyquinn/smartcar-api/models"

type Mock struct {}

func (mock *Mock) GetBatteryRange(id int) *models.Range {

}

func (mock *Mock) GetDoorSecurity(id int) *models.Doors {

}

func (mock *Mock) GetFuelRange(id int) *models.Range {

}

func (mock *Mock) GetVehicle(id int) *models.Vehicle {

}

func (mock *Mock) PushEngineButton(id int) *models.Status {

}