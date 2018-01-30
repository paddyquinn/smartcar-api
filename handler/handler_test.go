package handler

import (
  "testing"

  "github.com/paddyquinn/smartcar-api/dao"
)

func TestHandler_GetBatteryRange(t *testing.T) {
  hdlr := &Handler{DAO: &dao.Mock{}}
}

func TestHandler_GetDoorSecurity(t *testing.T) {
  hdlr := &Handler{DAO: &dao.Mock{}}
}

func TestHandler_GetFuelRange(t *testing.T) {
  hdlr := &Handler{DAO: &dao.Mock{}}
}

func TestHandler_GetVehicle(t *testing.T) {
  hdlr := &Handler{DAO: &dao.Mock{}}
}

func TestHandler_PushEngineButton(t *testing.T) {
  hdlr := &Handler{DAO: &dao.Mock{}}
}