package strutil

import "bytes"

func ToString(data []byte) string {
	n := bytes.Index(data, []byte{0})
	if n>= 0 {
		return string(data[:n])
	} else {
		return string(data)
	}
}
