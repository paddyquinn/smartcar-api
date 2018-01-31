package dao

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/paddyquinn/smartcar-api/models/gm"
  "github.com/paddyquinn/smartcar-api/models/smartcar"
)

const (
  domain = "http://gmapi.azurewebsites.net/%s"
  getVehiclePath = "getVehicleInfoService"
  jsonCapitalized = "JSON"
  jsonContentType = "application/json"
)

// DAO is a Data Access Object to the GM API
type DAO struct {}

func (dao *DAO) GetBatteryRange(id string) (*smartcar.Range, error) {
  return nil, nil
}

func (dao *DAO) GetDoorSecurity(id string) (*smartcar.Doors, error) {
  return nil, nil
}

func (dao *DAO) GetFuelRange(id string) (*smartcar.Range, error) {
  return nil, nil
}

func (dao *DAO) GetVehicle(id string) (*smartcar.Vehicle, error) {
  requestBody := &gm.RequestBody{
    ID: id,
    ResponseType: jsonCapitalized,
  }

  response, err := post(requestBody)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  gmVehicle := &gm.Vehicle{}
  decoder := json.NewDecoder(response.Body)
  err = decoder.Decode(gmVehicle)
  if err != nil {
    return nil, err
  }

  return gmVehicle.ToSmartcar(), nil
}

func (dao *DAO) PushEngineButton(id string) (*smartcar.Status, error) {
  return nil, nil
}

func post(requestBody *gm.RequestBody) (*http.Response, error) {
  jsonBytes, err := json.Marshal(requestBody)
  if err != nil {
    return nil, err
  }

  reader := bytes.NewReader(jsonBytes)
  return http.Post(fmt.Sprintf(domain, getVehiclePath), jsonContentType, reader)
}