package packet

import (
	"encoding/json"
	"fmt"
	"github.com/patricknoir/F12020/pkg/common/strutil"
)

/*
Event Packet
This packet gives details of events that happen during the course of a session.

Frequency: When the event occurs
Size: 35 bytes (Packet size updated in Beta 3)
Version: 1
 */

type EventDetails struct {
	FastestLap struct {
		VehicleIdx uint8
		LapTime	   float32
	}
	Retirement struct {
		VehicleIdx uint8
	}
	TeamMateInPits struct {
		VehicleIdx uint8
	}
	RaceWinner struct {
		VehicleIdx uint8
	}
	Penalty struct {
		PenaltyType			uint8          // Penalty type – see Appendices
		InfringementType	uint8          // Infringement type – see Appendices
		VehicleIdx		    uint8          // Vehicle index of the car the penalty is applied to
		OtherVehicleIdx		uint8          // Vehicle index of the other car involved
		Time				uint8          // Time gained, or time spent doing action in seconds
		LapNum				uint8          // Lap the penalty occurred on
		PlacesGained		uint8          // Number of places gained by this
	}
	SpeedTrap struct {
		VehicleIdx 	uint8
		Speed 		float32				   // Top speed achieved in kilometres per hour
	}
}

type EventStringCode [4]byte
func (esc EventStringCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(strutil.ToString(esc[:]))
}
const (
	SessionStarted		= "SSTA"
	SessionEnded		= "SEND"
	FastestLap			= "FTLP"
	Retirement			= "RTMT"
	DRSEnabled			= "DRSE"
	DRSDisabled			= "DRSD"
	TeamMateInPits		= "TMPT"
	ChequeredFlag		= "CHQF"
	RaceWinner			= "RCWN"
	PenaltyIssued		= "PENA"
	SpeedTrapTriggered	= "SPTP"
)

func (esc EventStringCode) ToString() string {
	return fmt.Sprintf("%s", esc)
}

type EventDataPacket struct {
	Header 					PacketHeader
	EventStringCode			EventStringCode	   // Event string code, see below
	EventDetails 			EventDetails
}