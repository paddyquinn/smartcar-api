package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paddyquinn/smartcar-api/dao"
	"github.com/paddyquinn/smartcar-api/errors"
	"github.com/paddyquinn/smartcar-api/models/smartcar"
)

const id = "id"

// Handler is a struct that handles requests from our web server
type Handler struct {
	DAO dao.Interface
}

// GetBatteryRange gets the battery percentage for a vehicle by id
func (hdlr *Handler) GetBatteryRange(c *gin.Context) {
	hdlr.handleGet(c, hdlr.DAO.GetBatteryRange)
}

// GetDoorSecurity gets the locked status of each door
func (hdlr *Handler) GetDoorSecurity(c *gin.Context) {
	hdlr.handleGet(c, hdlr.DAO.GetDoorSecurity)
}

// GetFuelRange gets the fuel percentage for a vehicle by id
func (hdlr *Handler) GetFuelRange(c *gin.Context) {
	hdlr.handleGet(c, hdlr.DAO.GetFuelRange)
}

// GetVehicle gets a vehicle by id
func (hdlr *Handler) GetVehicle(c *gin.Context) {
	hdlr.handleGet(c, hdlr.DAO.GetVehicle)
}

// PushEngineButton attempts to start or stop a vehicle by id based on the POST data passed
func (hdlr *Handler) PushEngineButton(c *gin.Context) {
	engineRequest, err := smartcar.NewEngineRequest(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	status, err := hdlr.DAO.PushEngineButton(c.Param(id), engineRequest.ToEnum())
	if err != nil {
		statusCode := http.StatusInternalServerError

		// Check for an UnidentifiedCommandError, in which case a user
		// error occurred and a 400 response code should be returned
		_, ok := err.(*errors.UnidentifiedCommandError)
		if ok {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, err)
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
