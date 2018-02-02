package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paddyquinn/smartcar-api/errors"
	"github.com/paddyquinn/smartcar-api/models/gm"
	"github.com/paddyquinn/smartcar-api/models/smartcar"
	"github.com/paddyquinn/smartcar-api/util"
)

const (
	domain          = "http://gmapi.azurewebsites.net/%s"
	gmStart         = "START_VEHICLE"
	gmStop          = "STOP_VEHICLE"
	jsonCapitalized = "JSON"
	jsonContentType = "application/json"

	// paths
	getDoorSecurityPath  = "getSecurityStatusService"
	getRangePath         = "getEnergyService"
	getVehiclePath       = "getVehicleInfoService"
	pushEngineButtonPath = "actionEngineService"
)

// DAO is a Data Access Object to the GM API
type DAO struct{}

// GetBatteryRange retrieves the battery percentage from the GM API and converts it to a smartcar compliant response
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
		return nil, errors.NewDecodeError()
	}

	return gmBatteryRange.ToSmartcar(), nil
}

// GetDoorSecurity gets the security status for each door of a GM vehicle
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
		return nil, errors.NewDecodeError()
	}

	return gmDoorsResponse.ToSmartcar(), nil
}

// GetFuelRange retrieves the fuel percentage from the GM API and converts it to a smartcar compliant response
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
		return nil, errors.NewDecodeError()
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
		return nil, errors.NewDecodeError()
	}

	return gmVehicle.ToSmartcar(), nil
}

// PushEngineButton attempts to start or stop the engine of a GM vehicle
func (dao *DAO) PushEngineButton(id string, cmd int) (*smartcar.Status, error) {
	var command string
	switch cmd {
	case smartcar.Start:
		command = gmStart
	case smartcar.Stop:
		command = gmStop
	default:
		return nil, errors.NewUnidentifiedCommandError()
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
		return nil, errors.NewDecodeError()
	}

	return gmStatus.ToSmartcar(), nil
}

func post(path string, requestBody *gm.RequestBody) (*http.Response, error) {
	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, errors.NewJSONMarshalError()
	}

	reader := bytes.NewReader(jsonBytes)
	rsp, err := http.Post(fmt.Sprintf(domain, path), jsonContentType, reader)
	if err != nil {
		return nil, errors.NewPostError()
	}

	return rsp, nil
}
