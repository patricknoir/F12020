package packet

/*
Lap Data Packet
The lap data packet gives details of all the cars in the session.

Frequency: Rate as specified in menus
Size: 1190 bytes (Struct updated in Beta 3)
Version: 1
 */

type PitStatus uint8
const (
	PitNone = 0
	Pitting = 1
	InPit	= 2
)

type Sector uint8
const (
	Sector1 = 0
	Sector2 = 1
	Sector3 = 2
)

type LapStatus uint8
const (
	LapValid 	= 0
	LapInvalid	= 1
)

type DriverStatus uint8
const (
	DriverInGarage	= 0
	DriverFlyingLap = 1
	DriverInLap 	= 2
	DriverOutLap	= 3
	DriverOnTrack	= 4
)

type ResultStatus uint8
const (
	ResultInvalid		= 0
	ResultInactive		= 1
	ResultActive		= 2
	ResultFinished		= 3
	ResultDisqualified 	= 4
	ResultNotClassified = 5
	ResultRetired		= 6
)

type LapInfoData struct
{
	LastLapTime 				float32               // Last lap time in seconds
	CurrentLapTime 				float32            	  // Current time around the lap in seconds

//UPDATED in Beta 3:
	Sector1TimeInMS				uint16           	  // Sector 1 time in milliseconds
	Sector2TimeInMS				uint16           	  // Sector 2 time in milliseconds
	BestLapTime					float32               // Best lap time of the session in seconds
	BestLapNum					uint8                 // Lap number best time achieved on
	BestLapSector1TimeInMS  	uint16   			  // Sector 1 time of best lap in the session in milliseconds
	BestLapSector2TimeInMS  	uint16   			  // Sector 2 time of best lap in the session in milliseconds
	BestLapSector3TimeInMS  	uint16   			  // Sector 3 time of best lap in the session in milliseconds
	BestOverallSector1TimeInMS  uint16   			  // Best overall sector 1 time of the session in milliseconds
	BestOverallSector1LapNum 	uint8      			  // Lap number best overall sector 1 time achieved on
	BestOverallSector2TimeInMS 	uint16   			  // Best overall sector 2 time of the session in milliseconds
	BestOverallSector2LapNum 	uint8      			  // Lap number best overall sector 2 time achieved on
	BestOverallSector3TimeInMS  uint16   			  // Best overall sector 3 time of the session in milliseconds
	BestOverallSector3LapNum 	uint8      			  // Lap number best overall sector 3 time achieved on


	LapDistance 				float32               // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance				float32               // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta				float32               // Delta in seconds for safety car
	CarPosition 				uint8    			  // Car race position
	CurrentLapNum 				uint8    			  // Current lap number
	PitStatus					PitStatus             // 0 = none, 1 = pitting, 2 = in pit area
	Sector                    	Sector				  // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid			LapStatus	          // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                 	uint8				  // Accumulated time penalties in seconds to be added
	GridPosition              	uint8				  // Grid position the vehicle started the race in
	DriverStatus              	DriverStatus		  // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus				ResultStatus
}

type LapDataPacket struct {
	Header PacketHeader
	Payload [22]LapInfoData
}