package packet

/*
Session Packet
The session packet includes details about the current session in progress.

Frequency: 2 per second
Size: 251 bytes (Packet size updated in Beta 3)
Version: 1
 */

type ZoneFlag int8
const (
	Unknown = -1
	None 	= 0
	Green	= 1
	Blue	= 2
	Yellow	= 3
	Red		= 4
)

type MarshalZone struct
{
	Start	float32   // Fraction (0..1) of way through the lap the marshal zone starts
	Flag	ZoneFlag  // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type SessionType uint8
const (
	UnknownSession = 0
	P1 			   = 1
	P2			   = 2
	P3			   = 3
	ShortPractice  = 4
	Q1			   = 5
	Q2			   = 6
	Q3			   = 7
	ShortQualify   = 8
	OSQ			   = 9
	Race		   = 10
	SprintRace	   = 11
	TimeTrial	   = 12
)

type Weather uint8
const (
	Clear		= 0
	LightCloud	= 1
	Overcast	= 2
	LightRain	= 3
	HeavyRain	= 4
	Storm		= 5
)

type Formula uint8
const (
	ModernF1	= 0
	ClassicF1	= 1
	F2_2019		= 2
	GenericF1	= 3
	F2_2020		= 4
)

type WeatherForecastSample struct
{
	SessionType			SessionType      // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2, 12 = Time Trial
	TimeOffset			uint8            // Time in minutes the forecast is for
	Weather				Weather          // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature    int8             // Track temp. in degrees celsius
	AirTemperature 		int8             // Air temp. in degrees celsius
}

type SafetyCarStatus uint8
const (
	NoSafetyCar		 = 0
	FullSafetyCar	 = 1
	VirtualSafetyCar = 2
)

type NetworkGameMode uint8
const (
	Offline		= 0
	Online		= 1
)

type SessionDataPacket struct
{
	Header 						PacketHeader
	Weather     				Weather
	TrackTemperature 			int8          	 			// Track temp. in degrees celsius
	AirTemperature				int8             			// Air temp. in degrees celsius
	TotalLaps					uint8   		 			// Total number of laps in this race
	TrackLength					uint16           			// Track length in metres
	SessionType					SessionType
	TrackId 					int8             			// -1 for unknown, 0-21 for tracks, see appendix
	Formula						Formula
	SessionTimeLeft 			uint16           			// Time left in session in seconds
	SessionDuration 			uint16           			// Session duration in seconds
	PitSpeedLimit 				uint8            			// Pit speed limit in kilometres per hour
	GamePaused					uint8            			// Whether the game is paused
	IsSpectating				uint8            			// Whether the player is spectating
	SpectatorCarIndex			uint8	         			// Index of the car being spectated
	SliProNativeSupport			uint8			 			// SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones 			uint8            			// Number of marshal zones to follow
	MarshalZones				[21]uint8        			// List of marshal zones – max 21
    SafetyCarStatus				SafetyCarStatus
	NetworkGame					NetworkGameMode
	NumWeatherForecastSamples	uint8 			 			// Number of weather samples to follow
	WeatherForecastSamples		[20]WeatherForecastSample   // Array of weather forecast samples
}