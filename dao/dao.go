package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"errors"
	"github.com/paddyquinn/smartcar-api/models/gm"
	"github.com/paddyquinn/smartcar-api/models/smartcar"
	"github.com/paddyquinn/smartcar-api/util"
)

const (
	domain                   = "http://gmapi.azurewebsites.net/%s"
	gmStart                  = "START_VEHICLE"
	gmStop                   = "STOP_VEHICLE"
	jsonCapitalized          = "JSON"
	jsonContentType          = "application/json"
	unidentifiedCommandError = "could not identify engine command"

	// paths
	getDoorSecurityPath  = "getSecurityStatusService"
	getRangePath         = "getEnergyService"
	getVehiclePath       = "getVehicleInfoService"
	pushEngineButtonPath = "actionEngineService"
)

// DAO is a Data Access Object to the GM API
type DAO struct{}

func (dao *DAO) GetBatteryRange(id string) (smartcar.Model, error) {
	requestBody := &gm.RequestBody{
		ID:           id,
		ResponseType: jsonCapitalized,
	}

	response, err := post(getRangePath, requestBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	gmBatteryRange := &gm.BatteryRange{}
	err = util.Decode(response.Body, gmBatteryRange)
	if err != nil {
		return nil, err
	}

	return gmBatteryRange.ToSmartcar(), nil
}

func (dao *DAO) GetDoorSecurity(id string) (smartcar.Model, error) {
	requestBody := &gm.RequestBody{
		ID:           id,
		ResponseType: jsonCapitalized,
	}

	response, err := post(getDoorSecurityPath, requestBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	gmDoorsResponse := &gm.DoorsResponse{}
	err = util.Decode(response.Body, gmDoorsResponse)
	if err != nil {
		return nil, err
	}

	return gmDoorsResponse.ToSmartcar(), nil
}

func (dao *DAO) GetFuelRange(id string) (smartcar.Model, error) {
	requestBody := &gm.RequestBody{
		ID:           id,
		ResponseType: jsonCapitalized,
	}

	response, err := post(getRangePath, requestBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	gmFuelRange := &gm.FuelRange{}
	err = util.Decode(response.Body, gmFuelRange)
	if err != nil {
		return nil, err
	}

	return gmFuelRange.ToSmartcar(), nil
}

// GetVehicle makes a call to the GM API for vehicle info by ID
func (dao *DAO) GetVehicle(id string) (smartcar.Model, error) {
	requestBody := &gm.RequestBody{
		ID:           id,
		ResponseType: jsonCapitalized,
	}

	response, err := post(getVehiclePath, requestBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	gmVehicle := &gm.Vehicle{}
	err = util.Decode(response.Body, gmVehicle)
	if err != nil {
		return nil, err
	}

	return gmVehicle.ToSmartcar(), nil
}

func (dao *DAO) PushEngineButton(id string, cmd int) (*smartcar.Status, error) {
	var command string
	switch cmd {
	case smartcar.Start:
		command = gmStart
	case smartcar.Stop:
		command = gmStop
	default:
		return nil, errors.New(unidentifiedCommandError)
	}

	requestBody := &gm.RequestBody{
		ID:           id,
		Command:      command,
		ResponseType: jsonCapitalized,
	}

	response, err := post(pushEngineButtonPath, requestBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	gmStatus := &gm.ActionResult{}
	err = util.Decode(response.Body, gmStatus)
	if err != nil {
		return nil, err
	}

	return gmStatus.ToSmartcar(), nil
}

func post(path string, requestBody *gm.RequestBody) (*http.Response, error) {
	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonBytes)
	return http.Post(fmt.Sprintf(domain, path), jsonContentType, reader)
}
