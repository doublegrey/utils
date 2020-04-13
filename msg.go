package utils

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

type Message struct {
	Header Header
	Data   []byte
}
type Header struct {
	ConsumerID [16]byte
	ProducerID [16]byte
	DataLength uint32
}

func (m Message) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	m.Header.DataLength = uint32(len(m.Data))
	if err := binary.Write(buf, binary.BigEndian, &m.Header); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, &m.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (m Message) Decode(c net.Conn) error {
	buf := make([]byte, 36)
	if _, err := io.ReadFull(c, buf); err != nil {
		return err
	}
	if err := binary.Read(bytes.NewReader(buf), binary.BigEndian, &m.Header); err != nil { // FIXME: https://stackoverflow.com/questions/41400639/is-binary-read-slow
		return err
	}
	m.Data = make([]byte, m.Header.DataLength)
	_, err := io.ReadFull(c, m.Data)
	if err != nil {
		return err
	}
	return nil
}
