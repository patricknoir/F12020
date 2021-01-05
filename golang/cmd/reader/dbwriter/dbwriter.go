package dbwriter

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/patricknoir/F12020/pkg/common/packet"
	"sync"
	"time"
)

type MetricsData struct {
	SessionUID 					uint64
	Frame						uint32
	SessionType					packet.SessionType
	SessionTypeName 			string
	TrackID						int8
	TrackName					string
	CarType						string
	PlayerCarIndex				uint8
	TrackLength					uint16
	TotalLaps	    			uint8
	TeamID						uint8
	TeamName					string
	CurrentLap					uint8
	SessionTime					float32
	Speed						uint16                 // Speed of car in kilometres per hour
	Throttle					float32                // Amount of throttle applied (0.0 to 1.0)
	Steer						float32                // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake						float32                // Amount of brake applied (0.0 to 1.0)
	Gear						int8
	EngineRPM					uint16                 // Engine RPM
	Drs							packet.DrsStatus              // 0 = off, 1 = on
	RevLightsPercent			uint8     		       // Rev lights indicator (percentage)
	BrakesTemperature			[4]uint16          	   // Brakes temperature (celsius)
	TyresSurfaceTemperature 	[4]uint8     	       // Tyres surface temperature (celsius)
	TyresInnerTemperature		[4]uint8      		   // Tyres inner temperature (celsius)
	EngineTemperature			uint16
	TyresPressure				[4]float32             // Tyres pressure (PSI)
	CurrentLapTime 				float32            	  // Current time around the lap in seconds
	Sector1TimeInMS				uint16           	  // Sector 1 time in milliseconds
	Sector2TimeInMS				uint16           	  // Sector 2 time in milliseconds
	LapDistance 				float32               // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance				float32               // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	Timestamp					time.Time
	Mutex 						sync.Mutex
}

var baseTime time.Time = time.Now()
var currentSession *MetricsData = nil

func UpdateMetricsDataWithSessionPacket(data packet.SessionDataPacket) error {
	session := getOrCreateSession(data.Header)
	session.SessionType = data.SessionType
	session.SessionTypeName = packet.SessionTypeMap[data.SessionType]
	session.TrackID = data.TrackId
	session.TrackName = packet.TrackIDMap[data.TrackId]
	session.CarType = packet.FormulaTypeMap[data.Formula]
	session.TrackLength = data.TrackLength
	session.TotalLaps = data.TotalLaps
	return nil
}

func UpdateMetricsDataWithParticipantsPacket(data packet.ParticipantsDataPacket) error {
	session := getOrCreateSession(data.Header)
	session.TeamID = data.Participants[data.Header.PlayerCarIndex].TeamId
	session.TeamName = packet.TeamIDMap[session.TeamID]
	return nil
}

func UpdateMetricsDataWithLapData(data packet.LapDataPacket) error {
	session := getOrCreateSession(data.Header)
	session.CurrentLapTime = data.Payload[session.PlayerCarIndex].CurrentLapTime
	session.Sector1TimeInMS = data.Payload[session.PlayerCarIndex].Sector1TimeInMS
	session.Sector2TimeInMS = data.Payload[session.PlayerCarIndex].Sector2TimeInMS
	session.LapDistance = data.Payload[session.PlayerCarIndex].LapDistance
	session.TotalDistance = data.Payload[session.PlayerCarIndex].TotalDistance
	session.CurrentLap = data.Payload[session.PlayerCarIndex].CurrentLapNum
	return nil
}

func UpdateMetricsDataWithTelemetry(data packet.CarTelemetryDataPacket) error {
	session := getOrCreateSession(data.Header)
	session.Speed = data.CarTelemetryData[session.PlayerCarIndex].Speed
	session.Throttle = data.CarTelemetryData[session.PlayerCarIndex].Throttle
	session.Brake = data.CarTelemetryData[session.PlayerCarIndex].Brake
	session.Steer = data.CarTelemetryData[session.PlayerCarIndex].Steer
	session.Gear = int8(data.CarTelemetryData[session.PlayerCarIndex].Gear)
	session.EngineRPM = data.CarTelemetryData[session.PlayerCarIndex].EngineRPM
	session.Drs = data.CarTelemetryData[session.PlayerCarIndex].Drs
	session.RevLightsPercent = data.CarTelemetryData[session.PlayerCarIndex].RevLightsPercent
	session.BrakesTemperature = data.CarTelemetryData[session.PlayerCarIndex].BrakesTemperature
	session.TyresSurfaceTemperature = data.CarTelemetryData[session.PlayerCarIndex].TyresSurfaceTemperature
	session.TyresInnerTemperature = data.CarTelemetryData[session.PlayerCarIndex].TyresInnerTemperature
	session.EngineTemperature = data.CarTelemetryData[session.PlayerCarIndex].EngineTemperature
	session.TyresPressure = data.CarTelemetryData[session.PlayerCarIndex].TyresPressure
	return nil
}

func getOrCreateSession(h packet.PacketHeader) *MetricsData {
	if currentSession==nil || currentSession.SessionUID!=h.SessionUID {
		currentSession = &MetricsData{}
		currentSession.SessionUID = h.SessionUID
		currentSession.PlayerCarIndex = h.PlayerCarIndex
	}

	if currentSession.Frame!= h.FrameIdentifier {
		writeMeasurement()
	}
	currentSession.SessionTime = h.SessionTime
	currentSession.Frame = h.FrameIdentifier
	currentSession.Timestamp = baseTime.Add(time.Millisecond * time.Duration(h.SessionTime*1000)) //have to extend this to 1000 as base unit is milliseconds and I want to plot Microseconds

	return currentSession
}

func writeMeasurement() error {
	_, writer := getOrCreateClient()
	p := influxdb2.NewPointWithMeasurement("telemetry").
		AddTag("uid", fmt.Sprintf("%d", currentSession.SessionUID)).
		AddTag("teamName", currentSession.TeamName).
		AddTag("trackName", currentSession.TrackName).
		AddTag("currentLapTime", fmt.Sprintf("%f", currentSession.CurrentLapTime)).
		AddTag("lapDistance", fmt.Sprintf("%f", currentSession.LapDistance)).
		AddTag("currentLap", fmt.Sprintf("%d", currentSession.CurrentLap)).
		AddField("speed", int(currentSession.Speed)).
		AddField("throttle", currentSession.Throttle).
		AddField("brake", currentSession.Brake).
		AddField("gear", int(currentSession.Gear)).
		AddField("rpm", int(currentSession.EngineRPM)).
		AddField("steer", currentSession.Steer).
		SetTime(currentSession.Timestamp)
	fmt.Printf("Printing datapoint UID: %d, track: %s, speed: %d, lapTime: %f, time: %s\n",
		currentSession.SessionUID, currentSession.TrackName, currentSession.Speed, currentSession.CurrentLapTime, currentSession.Timestamp.String())
	writer.WritePoint(p)
	return nil
}

const token = "f12020:password"
const bucket = "F12020"
const org = "hkubx"

var _client influxdb2.Client
var _writer api.WriteAPI

func getOrCreateClient() (influxdb2.Client, api.WriteAPI) {
	if _client==nil || _writer==nil {
		_client = influxdb2.NewClient("http://localhost:8086", token)
		_writer = _client.WriteAPI(org, bucket)
	}
	return _client, _writer
}

func ClearResources() {
	_writer.Flush()
	_client.Close()
}