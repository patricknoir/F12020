package parser

import (
	"bytes"
	"encoding/binary"
	"github.com/patricknoir/F12020/pkg/common/errors"
	"github.com/patricknoir/F12020/pkg/common/packet"
)

type ParsedPacket struct {
	PacketID packet.PacketID
	LapData *packet.LapDataPacket
}

func Parse(data []byte) (ParsedPacket, error) {
	pp := ParsedPacket{}
	header := packet.PacketHeader{}
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.LittleEndian, &header)
	if !errors.LogError(err) {
		pp.PacketID = header.PacketID
		r.Seek(0, 0)
		switch pp.PacketID {
		case packet.LapData:
			ld := packet.LapDataPacket{}
			err = binary.Read(r, binary.LittleEndian, &ld)
			pp.LapData = &ld
		}
	}
	return pp, err
}