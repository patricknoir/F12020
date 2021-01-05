package main

import (
	"encoding/binary"
	"fmt"
	"github.com/patricknoir/F12020/cmd/reader/dbwriter"
	"github.com/patricknoir/F12020/pkg/common/errors"
	"github.com/patricknoir/F12020/pkg/common/packet"
	"os"
	"strconv"
)

const path string = "./time-trial/"

const maxIndex int = 37800

func main() {
	var basePath string = path
	frameMap := make(map[uint32][]string)
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}
	n := 0
	counter := make(map[string]int)
	for {
		f, err := os.Open(basePath + "Packet-" + strconv.Itoa(n) + ".bin")

		if err!=nil {
			fmt.Println(err)
			fmt.Printf("No more files to read, Last index: %d\n", n-1)
			f.Close()
			break				
		}
		header := packet.PacketHeader{}
		err = binary.Read(f, binary.LittleEndian, &header)
		if !errors.LogError(err) {
			//if n<400 {
			printPacketPayload(header, f)
			mapFrame(header, f, frameMap)
			//}
			countPackets(counter, &header)
		}
		f.Close()
		n++
	}
	//fmt.Printf("%+v\n", frameMap)
	//fmt.Printf("%+v\n", counter)
	dbwriter.ClearResources()
}

func mapFrame(h packet.PacketHeader, f *os.File, frameMap map[uint32][]string) {
	var p interface{}
	switch h.PacketID {
	case packet.Session:
		p, _ = handleSessionPacket(h, f)
	case packet.Motion:
		p, _ = handleMotionPacket(h, f)
	case packet.LobbyInfo:
		p, _ = handleLobbyInfoPacket(h, f)
	case packet.LapData:
		p, _ = handleLapDataPacket(h, f)
	case packet.FinalClassification:
		p, _ = handleFinalClassificationPacket(h, f)
	case packet.Event:
		p, _ = handleEventPacket(h, f)
	case packet.Participants:
		p, _ = handleParticipantsPacket(h, f)
	case packet.CarTelemetry:
		p, _ = handleCarTelemetryPacket(h, f)
	default:
		//fmt.Printf("skipping packet %d\n", id)
	}
	_ = p
	frameId := h.FrameIdentifier
	frameMap[frameId] = append(frameMap[frameId], packet.PacketIDMap[h.PacketID])
}

func printPacketPayload(h packet.PacketHeader, f *os.File) {
	var p interface{}
	switch h.PacketID {
		case packet.Session:
			session, _ := handleSessionPacket(h, f)
			dbwriter.UpdateMetricsDataWithSessionPacket(session)
			p = session
		case packet.Motion:
			p, _ = handleMotionPacket(h, f)
		case packet.LobbyInfo:
			p, _ = handleLobbyInfoPacket(h, f)
		case packet.LapData:
			lapData, _ := handleLapDataPacket(h, f)
			dbwriter.UpdateMetricsDataWithLapData(lapData)
			p = lapData
		case packet.FinalClassification:
			p, _ = handleFinalClassificationPacket(h, f)
		case packet.Event:
			p, _ = handleEventPacket(h, f)
		case packet.Participants:
			participants, _ := handleParticipantsPacket(h, f)
			dbwriter.UpdateMetricsDataWithParticipantsPacket(participants)
			p = participants
		case packet.CarTelemetry:
			telemetry, _ := handleCarTelemetryPacket(h, f)
			dbwriter.UpdateMetricsDataWithTelemetry(telemetry)
			p = telemetry
		default:
			//fmt.Printf("skipping packet %d\n", id)
	}

	_ = p
	//fmt.Println("Packet: " + packet.PacketIDMap[h.PacketID])
	//fmt.Printf("Packet: %s\t| Frame ID: %d\t| SessionTime: %f \n", packet.PacketIDMap[h.PacketID], h.FrameIdentifier, h.SessionTime)
	//out, err := json.MarshalIndent(p, "", "  ")
	//out, err := json.Marshal(p)
	//if !errors.LogError(err) && h.PacketID == packet.CarTelemetry {
	//	fmt.Println(string(out))
	//}
}

func handleCarTelemetryPacket(h packet.PacketHeader, f *os.File) (packet.CarTelemetryDataPacket, error) {
	t := packet.CarTelemetryDataPacket{}
	f.Seek(0, 0)
	err := binary.Read(f, binary.LittleEndian, &t)
	errors.LogError(err)
	return t, err
}

func handleSessionPacket(h packet.PacketHeader, f *os.File) (packet.SessionDataPacket, error) {
	s := packet.SessionDataPacket{}
	f.Seek(0, 0)
	err := binary.Read(f, binary.LittleEndian, &s)
	return s, err
}


func handleParticipantsPacket(h packet.PacketHeader, f *os.File) (packet.ParticipantsDataPacket, error) {
	p := packet.ParticipantsDataPacket{}
	p.Header = h
	err := binary.Read(f, binary.LittleEndian, &p.ActiveCars)
	if !errors.LogError(err) {
		err = binary.Read(f, binary.LittleEndian, &p.Participants)
		errors.LogError(err)
	}
	return p, err
}

func handleMotionPacket(h packet.PacketHeader, f *os.File) (packet.MotionDataPacket, error) {
	m := packet.MotionDataPacket{}
	m.Header = h
	err := binary.Read(f, binary.LittleEndian, &m.Payload)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.SuspensionPosition)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.SuspensionVelocity)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.SuspensionAcceleration)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.WheelSpeed)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.WheelSlip)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.LocalVelocityX)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.LocalVelocityY)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.LocalVelocityZ)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.AngularVelocityX)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.AngularVelocityY)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.AngularVelocityZ)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.AngularAccelerationX)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.AngularAccelerationY)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.AngularAccelerationZ)
	if errors.LogError(err) {
		return m, err
	}
	err = binary.Read(f, binary.LittleEndian, &m.FrontWheelsAngle)
	if errors.LogError(err) {
		return m, err
	}
	return m, err
}

func handleEventPacket(h packet.PacketHeader, f *os.File) (packet.EventDataPacket, error) {
	e := packet.EventDataPacket{}
	e.Header = h
	err := binary.Read(f, binary.LittleEndian, &e.EventStringCode)
	if !errors.LogError(err) {
		ed := packet.EventDetails{}
		e.EventDetails = ed
		switch e.EventStringCode.ToString() {
			case packet.FastestLap:
				err = binary.Read(f, binary.LittleEndian, &e.EventDetails.FastestLap)
			case packet.Retirement:
				err = binary.Read(f, binary.LittleEndian, &e.EventDetails.Retirement)
			case packet.TeamMateInPits:
				err = binary.Read(f, binary.LittleEndian, &e.EventDetails.TeamMateInPits)
			case packet.PenaltyIssued:
				err = binary.Read(f, binary.LittleEndian, &e.EventDetails.Penalty)
			case packet.RaceWinner:
				err = binary.Read(f, binary.LittleEndian, &e.EventDetails.RaceWinner)
			case packet.SpeedTrapTriggered:
				err = binary.Read(f, binary.LittleEndian, &e.EventDetails.SpeedTrap)
		}
		errors.LogError(err)
	}
	return e, err
}

func handleFinalClassificationPacket(h packet.PacketHeader, f *os.File) (packet.FinalClassificationDataPacket, error) {
	fc := packet.FinalClassificationDataPacket{}
	fc.Header = h
	err := binary.Read(f, binary.LittleEndian, &fc.NumCars)
	if !errors.LogError(err) {
		err = binary.Read(f, binary.LittleEndian, &fc.ClassificationData)
		errors.LogError(err)
	}
	return fc, err
}

func handleLobbyInfoPacket(h packet.PacketHeader, f *os.File) (packet.LobbyInfoDataPacket, error) {
	li := packet.LobbyInfoDataPacket{}
	li.Header = h
	err := binary.Read(f, binary.LittleEndian, &li.NumPlayers)
	if !errors.LogError(err) {
		err = binary.Read(f, binary.LittleEndian, &li.LobbyPlayers)
		errors.LogError(err)
	}
	return li, err
}

func handleLapDataPacket(h packet.PacketHeader, f *os.File) (packet.LapDataPacket, error) {
	ld := packet.LapDataPacket{}
	ld.Header = h
	err := binary.Read(f, binary.LittleEndian, &ld.Payload)
	return ld, err
}

func countPackets(counterMap map[string]int, h *packet.PacketHeader) {
	var packetName string = packet.PacketIDMap[h.PacketID]
	counterMap[packetName] = counterMap[packetName] + 1
}