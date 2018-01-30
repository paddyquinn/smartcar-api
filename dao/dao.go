package dao

import "github.com/paddyquinn/smartcar-api/models"

type DAO struct {}

func (dao *DAO) GetBatteryRange(id int) *models.Range {

}

func (dao *DAO) GetDoorSecurity(id int) *models.Doors {

}

func (dao *DAO) GetFuelRange(id int) *models.Range {

}

func (dao *DAO) GetVehicle(id int) *models.Vehicle {

}

func (dao *DAO) PushEngineButton(id int) *models.Status {

}