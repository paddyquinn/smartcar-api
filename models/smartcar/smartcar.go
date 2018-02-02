package smartcar

import (
	"github.com/paddyquinn/smartcar-api/util"
	"io"
)

const (
	// Go Enumeration

	// NoOp is the default for this enum
	NoOp = iota

	// Start is the enum value for a start engine request
	Start

	// Stop is the enum value for a stop engine request
	Stop

	// string constants

	start   = "START"
	stop    = "STOP"
	success = "success"
)

// Model is an alias for an interface so that it is clear in a
// function that one of the smartcar responses should be passed
type Model interface{}

// Requests

// EngineRequest holds the POST data sent to the engine endpoint
type EngineRequest struct {
	Action string `json:"action"`
}

// NewEngineRequest creates a new EngineRequest object from the POST data passed
func NewEngineRequest(requestBody io.ReadCloser) (*EngineRequest, error) {
	engineRequest := &EngineRequest{}
	err := util.Decode(requestBody, engineRequest)
	if err != nil {
		return nil, err
	}

	return engineRequest, nil
}

// ToEnum returns the appropriate action to be taken based on the
// POST data, NoOp is the default in case bad POST data is sent
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

// Door contains the location and locked information about a door
type Door struct {
	Location string `json:"location"`
	Locked   bool   `json:"locked"`
}

// Doors is just an array of doors
type Doors []*Door

// Range is a percentage of either battery or fuel remaining
type Range struct {
	Percent float64 `json:"percent"`
}

// Status contains a status message (should always be success or error)
type Status struct {
	Value string `json:"status"`
}

// IsSuccess just indicates whether or not this is a success status
func (s *Status) IsSuccess() bool {
	return s.Value == success
}

// Vehicle represents a smartcar vehicle
type Vehicle struct {
	VIN        string `json:"vin"`
	Color      string `json:"color"`
	NumDoors   int    `json:"doorCount"`
	DriveTrain string `json:"driveTrain"`
}
