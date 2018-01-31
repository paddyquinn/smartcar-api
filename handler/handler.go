package handler

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/smartcar-api/dao"
)

const id = "id"

type Handler struct {
  DAO dao.Interface
}

func (hdlr *Handler) GetBatteryRange(c *gin.Context) {

}

func (hdlr *Handler) GetDoorSecurity(c *gin.Context) {

}

func (hdlr *Handler) GetFuelRange(c *gin.Context) {

}

func (hdlr *Handler) GetVehicle(c *gin.Context) {
  vehicle, err := hdlr.DAO.GetVehicle(c.Param(id))
  if err != nil {
    // TODO: handle error
  }

  c.JSON(http.StatusOK, vehicle)
}

func (hdlr *Handler) PushEngineButton(c *gin.Context) {

}