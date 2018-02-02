package handler

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/smartcar-api/dao"
  "github.com/paddyquinn/smartcar-api/models/smartcar"
)

const id = "id"

type Handler struct {
  DAO dao.Interface
}

func (hdlr *Handler) GetBatteryRange(c *gin.Context) {
  hdlr.handleGet(c, hdlr.DAO.GetBatteryRange)
}

func (hdlr *Handler) GetDoorSecurity(c *gin.Context) {
  hdlr.handleGet(c, hdlr.DAO.GetDoorSecurity)
}

func (hdlr *Handler) GetFuelRange(c *gin.Context) {
  hdlr.handleGet(c, hdlr.DAO.GetFuelRange)
}

func (hdlr *Handler) GetVehicle(c *gin.Context) {
  hdlr.handleGet(c, hdlr.DAO.GetVehicle)
}

func (hdlr *Handler) PushEngineButton(c *gin.Context) {
  engineRequest, err := smartcar.NewEngineRequest(c.Request.Body)
  if err != nil {
    // TODO: rethink this with sane error codes
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  status, err := hdlr.DAO.PushEngineButton(c.Param(id), engineRequest.ToEnum())
  if err != nil {
    // TODO: rethink this with sane error codes
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  statusCode := http.StatusInternalServerError
  if status.IsSuccess() {
    statusCode = http.StatusOK
  }

  c.JSON(statusCode, status)
}

func (hdlr *Handler) handleGet(c *gin.Context, daoFunc func(string) (smartcar.Model, error)) {
  model, err := daoFunc(c.Param(id))
  if err != nil {
    // TODO: rethink this with sane error codes
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, model)
}