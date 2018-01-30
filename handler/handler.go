package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/smartcar-api/dao"
)

type Handler struct {
  DAO dao.Interface
}

func (handler *Handler) GetBatteryRange(c *gin.Context) {

}

func (handler *Handler) GetDoorSecurity(c *gin.Context) {

}

func (handler *Handler) GetFuelRange(c *gin.Context) {

}

func (handler *Handler) GetVehicle(c *gin.Context) {

}

func (handler *Handler) PushEngineButton(c *gin.Context) {

}