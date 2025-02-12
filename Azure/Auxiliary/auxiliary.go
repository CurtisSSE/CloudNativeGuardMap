package Auxiliary

import "bytes"

func SecureWipe(buffer *bytes.Buffer) (overwrittenbuffer *bytes.Buffer) {
	for i := range buffer.Bytes() {
		buffer.Bytes()[i] = 0
	}
	buffer.Reset()
	return buffer
}
