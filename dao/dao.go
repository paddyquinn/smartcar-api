package dao

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "net/http"

  "github.com/paddyquinn/smartcar-api/models/gm"
  "github.com/paddyquinn/smartcar-api/models/smartcar"
)

const (
  domain = "http://gmapi.azurewebsites.net/%s"
  jsonCapitalized = "JSON"
  jsonContentType = "application/json"

  // paths
  getDoorSecurityPath = "getSecurityStatusService"
  getRangePath = "getEnergyService"
  getVehiclePath = "getVehicleInfoService"
)

// DAO is a Data Access Object to the GM API
type DAO struct {}

func (dao *DAO) GetBatteryRange(id string) (smartcar.Model, error) {
  requestBody := &gm.RequestBody{
    ID: id,
    ResponseType: jsonCapitalized,
  }

  response, err := post(getRangePath, requestBody)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  gmBatteryRange := &gm.BatteryRange{}
  return decodeResponse(response.Body, gmBatteryRange)
}

func (dao *DAO) GetDoorSecurity(id string) (smartcar.Model, error) {
  requestBody := &gm.RequestBody{
    ID: id,
    ResponseType: jsonCapitalized,
  }

  response, err := post(getDoorSecurityPath, requestBody)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  gmDoorsResponse := &gm.DoorsResponse{}
  return decodeResponse(response.Body, gmDoorsResponse)
}

func (dao *DAO) GetFuelRange(id string) (smartcar.Model, error) {
  requestBody := &gm.RequestBody{
    ID: id,
    ResponseType: jsonCapitalized,
  }

  response, err := post(getRangePath, requestBody)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  gmFuelRange := &gm.FuelRange{}
  return decodeResponse(response.Body, gmFuelRange)
}

// GetVehicle makes a call to the GM API for vehicle info by ID
func (dao *DAO) GetVehicle(id string) (smartcar.Model, error) {
  requestBody := &gm.RequestBody{
    ID: id,
    ResponseType: jsonCapitalized,
  }

  response, err := post(getVehiclePath, requestBody)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  gmVehicle := &gm.Vehicle{}
  return decodeResponse(response.Body, gmVehicle)
}

func (dao *DAO) PushEngineButton(id string) (smartcar.Model, error) {
  return nil, nil
}

func decodeResponse(body io.ReadCloser, model gm.Model) (smartcar.Model, error){
  decoder := json.NewDecoder(body)
  err := decoder.Decode(model)
  if err != nil {
    return nil, err
  }

  return model.ToSmartcar(), nil
}

func post(path string, requestBody *gm.RequestBody) (*http.Response, error) {
  jsonBytes, err := json.Marshal(requestBody)
  if err != nil {
    return nil, err
  }

  reader := bytes.NewReader(jsonBytes)
  return http.Post(fmt.Sprintf(domain, path), jsonContentType, reader)
}