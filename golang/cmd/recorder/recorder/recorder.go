package recorder

import (
	"github.com/patricknoir/F12020/pkg/common/errors"
	"net"
	"os"
	"strconv"
	"sync"
)

const maxBufferSize int = 2500

type Recorder struct {
	connection *net.UDPConn
	mutex sync.Mutex
	baseFilePath string
}

func New(host string, port int, baseFilePath string) (*Recorder, error) {
	recorder := Recorder{}
	recorder.baseFilePath = baseFilePath
	udpAddr, err := net.ResolveUDPAddr("udp", host + ":" + strconv.Itoa(port))
	errors.LogError(err)
	if err!=nil {
		return nil, err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	errors.LogError(err)
	if err!=nil {
		return nil, err
	}
	recorder.connection = udpConn
	return &recorder, nil
}

func Record(recorder *Recorder) {
	var n int = 0
	for {
		recorder.mutex.Lock()
		buffer := make([]byte, maxBufferSize)
		read, _, err := recorder.connection.ReadFromUDP(buffer)
		errors.LogError(err)
		data := buffer[:read]
		saveData(recorder.baseFilePath, data, n)
		n++
		recorder.mutex.Unlock()
	}
}

func saveData(baseFilePath string, data []byte, n int) (int, error) {
	filepath := baseFilePath + "/Packet-" + strconv.Itoa(n) + ".bin"
	f, err := os.Create(filepath)
	errors.LogError(err)
	defer f.Close()
	n2, err := f.Write(data)
	errors.LogError(err)
	f.Sync()
	return n2, nil
}
