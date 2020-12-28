package packet

/*
Car Status Packet
This packet details car statuses for all the cars in the race. It includes values such as the damage readings on the car.

Frequency: Rate as specified in menus
Size: 1344 bytes (Packet updated in Beta 3)
Version: 1
*/

type TractionControlStatus uint8
const (
	TractionOff 	= 0
	TractionPartial = 1
	TractionFull	= 2
)

type ABSStatus uint8
const (
	ABSOff = 0
	ABSOn  = 1
)

type FuelStatus uint8
const (
	LeanFuel	 = 0
	StandardFuel = 1
	RichFuel	 = 2
	MaxFuel		 = 3
)

type PitLimiterStatus uint8
const (
	PitLimiterOff = 0
	PitLimiterOn  = 1
)

type TyreCompoundType uint8
const (
	F1_C5		 = 16
	F1_C4 		 = 17
	F1_C3 		 = 18
	F1_C2 		 = 19
	F1_C1 		 = 20
	F1_Inter 	 = 7
	F1_Wet	 	 = 8
	Classic_Dry  = 9
	Classic_Wet  = 10
	F2_SuperSoft = 11
	F2_Soft		 = 12
	F2_Medium	 = 13
	F2_Hard		 = 14
	F2_Wet		 = 15
)

type ERSMode uint8
const (
	ERS_None		= 0
	ERS_Medium		= 1
	ERS_Overtake	= 2
	ERS_Hotlap		= 3
)

type CarStatusData struct
{
	TractionControl			TractionControlStatus          // 0 (off) - 2 (high)
	AntiLockBrakes			ABSStatus          			   // 0 (off) - 1 (on)
	FuelMix 				FuelStatus                     // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias			uint8       		           // Front brake bias (percentage)
	PitLimiterStatus		PitLimiterStatus	           // Pit limiter status - 0 = off, 1 = on
	FuelInTank				float32			               // Current fuel mass
	FuelCapacity			float32			               // Fuel capacity
	FuelRemainingLaps		float32			               // Fuel remaining in terms of laps (value on MFD)
	MaxRPM					uint16                         // Cars max RPM, point of rev limiter
	IdleRPM					uint16                         // Cars idle RPM
	MaxGears                uint8       				   // Maximum number of gears
	DrsAllowed 				uint8            			   // 0 = not allowed, 1 = allowed, -1 = unknown


// Added in Beta3:
	DrsActivationDistance	uint16      			       // 0 = DRS not available, non-zero - DRS will be available in [X] metres
	TyresWear				[4]uint8       	               // Tyre wear percentage
	ActualTyreCompound		TyreCompoundType
	VisualTyreCompound		uint8       				   // F1 visual (can be different from actual compound)
	TyresAgeLaps 			uint8       				   // Age in laps of the current set of tyres
	TyresDamage				[4]uint8       				   // Tyre damage (percentage)
	FrontLeftWingDamage		uint8       				   // Front left wing damage (percentage)
	FrontRightWingDamage	uint8       				   // Front right wing damage (percentage)
	RearWingDamage			uint8       				   // Rear wing damage (percentage)

// Added Beta 3:
	DrsFault				uint8       				   // Indicator for DRS fault, 0 = OK, 1 = fault
	EngineDamage			uint8       				   // Engine damage (percentage)
	GearBoxDamage			uint8       				   // Gear box damage (percentage)
	VehicleFiaFlags			ZoneFlag       				   // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ErsStoreEnergy			float32				           // ERS energy store in Joules
	ErsDeployMode			ERSMode       				   // ERS deployment mode, 0 = none, 1 = medium, 2 = overtake, 3 = hotlap
	ErsHarvestedThisLapMGUK float32						   // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32						   // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap float32						   // ERS energy deployed this lap
}

type PacketCarStatusData struct
{
	Header    		PacketHeader         // Header
	CarStatusData	[22]CarStatusData
}