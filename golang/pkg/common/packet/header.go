package packet

type PacketID uint8
const (
	Motion 		 		= 0 	// Contains all motion data for player’s car – only sent while player is in control
	Session 	 		= 1		// Data about the session – track, time left
	LapData 	 		= 2		// Data about all the lap times of cars in the session
	Event		 		= 3		// Various notable events that happen during a session
	Participants 		= 4		// List of participants in the session, mostly relevant for multiplayer
	CarSetups	 		= 5		// Packet detailing car setups for cars in the race
	CarTelemetry 		= 6		// Telemetry data for all cars
	CarStatus	 		= 7		// Status data for all cars such as damage
	FinalClassification = 8		// Final classification confirmation at the end of a race
	LobbyInfo			= 9		// Information about players in a multiplayer lobby
)

type PacketHeader struct {
	PacketFormat 	 uint16 			// 2020
	GameMajorVersion uint8				// Game major version - "X.00"
	GameMinorVersion uint8				// Game minor version - "1.XX"
	PacketVersion	 uint8				// Version of this packet type, all start from 1
	PacketID		 PacketID			// Identifier for the packet type, see below
	SessionUID		 uint64				// Unique identifier for the session
	SessionTime		 float32			// Session timestamp
	FrameIdentifier  uint32				// Identifier for the frame the data was retrieved on
	PlayerCarIndex   uint8				// Index of the player's car in the array

	// ADDED IN BETA 2:
	SecondaryPlayerCarIndex uint8		// Index of secondary player's car in the array (splitscreen), 255 if no second player
}
