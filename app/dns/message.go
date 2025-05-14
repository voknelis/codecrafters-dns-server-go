package dns

type Message struct {
	Header *Header
}

func (m *Message) Unmarshall(buf []byte) error {
	header := NewHeader()
	err := header.Unmarshall(buf[:12])
	if err != nil {
		return err
	}

	m.Header = header
	return nil
}

func (m *Message) Marshall() []byte {
	buf := make([]byte, 512)
	copy(buf[:12], m.Header.Marshall())

	return buf
}

func NewMessage() *Message {
	return &Message{
		Header: NewHeader(),
	}
}
