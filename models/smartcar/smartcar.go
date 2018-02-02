package smartcar

import (
	"github.com/paddyquinn/smartcar-api/util"
	"io"
)

const (
	NoOp = iota
	Start
	Stop

	start   = "START"
	stop    = "STOP"
	success = "success"
)

type Model interface{}

// Requests

type EngineRequest struct {
	Action string `json:"action"`
}

func NewEngineRequest(requestBody io.ReadCloser) (*EngineRequest, error) {
	engineRequest := &EngineRequest{}
	err := util.Decode(requestBody, engineRequest)
	if err != nil {
		return nil, err
	}

	return engineRequest, nil
}

func (req *EngineRequest) ToEnum() int {
	switch req.Action {
	case start:
		return Start
	case stop:
		return Stop
	default:
		return NoOp
	}
}

// Responses

type Door struct {
	Location string `json:"location"`
	Locked   bool   `json:"locked"`
}

type Doors []*Door

type Range struct {
	Percent float64 `json:"percent"`
}

type Status struct {
	Value string `json:"status"`
}

func (s *Status) IsSuccess() bool {
	return s.Value == success
}

type Vehicle struct {
	VIN        string `json:"vin"`
	Color      string `json:"color"`
	NumDoors   int    `json:"doorCount"`
	DriveTrain string `json:"driveTrain"`
}
