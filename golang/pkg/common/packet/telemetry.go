package packet

/*
Car Telemetry Packet
This packet details telemetry for all the cars in the race. It details various values that would be recorded on the car such as speed, throttle application, DRS etc.

Frequency: Rate as specified in menus
Size: 1307 bytes (Packet size updated in Beta 3)
Version: 1
 */

type DrsStatus uint8
const (
	DRSOff  = 0
	DRSOn   = 1
)

type GearStatus int8
const (
	Neutral			= 0
	Reverse			= -1
	FirstGear		= 1
	SecondGear		= 2
	ThirdGear		= 3
	ForthGear		= 4
	FifthGear		= 5
	SixthGear		= 6
	SeventhGear		= 7
	EighthGear		= 8
)

type CarTelemetryData struct
{
	Speed						uint16                 // Speed of car in kilometres per hour
	Throttle					float32                // Amount of throttle applied (0.0 to 1.0)
	Steer						float32                // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake						float32                // Amount of brake applied (0.0 to 1.0)
	Clutch						uint8                  // Amount of clutch applied (0 to 100)
	Gear						GearStatus             // Gear selected (1-8, N=0, R=-1)
	EngineRPM					uint16                 // Engine RPM
	Drs							DrsStatus              // 0 = off, 1 = on
	RevLightsPercent			uint8     		       // Rev lights indicator (percentage)
	BrakesTemperature			[4]uint16          	   // Brakes temperature (celsius)
	TyresSurfaceTemperature 	[4]uint8     	       // Tyres surface temperature (celsius)
	TyresInnerTemperature		[4]uint8      		   // Tyres inner temperature (celsius)
	EngineTemperature			uint16             	   // Engine temperature (celsius)
	TyresPressure				[4]float32             // Tyres pressure (PSI)
	SurfaceType					[4]uint8               // Driving surface, see appendices
}

type MFDPanel uint8
const (
	MFDCarSetup		= 0
	MFDPits			= 1
	MFDDamange		= 2
	MFDEngine		= 3
	MFDTemperatures = 4
	MFDClosed		= 255
)

type CarTelemetryDataPacket struct
{
	Header 			 PacketHeader
	CarTelemetryData [22]CarTelemetryData

	ButtonStatus	uint32                      // Bit flags specifying which buttons are being pressed currently - see appendices

// Added in Beta 3:
	MfdPanelIndex   				MFDPanel
	MfdPanelIndexSecondaryPlayer	MFDPanel
	SuggestedGear 					int8       // Suggested gear for the player (1-8) 0 if no gear suggested
}
