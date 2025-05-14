package dns

import (
	"fmt"
	"net"
)

type DnsServer struct {
	port     int
	listener *net.UDPConn
}

func (s *DnsServer) Listen() {
	fmt.Printf("Starting DNS server at port %d\n", s.port)

	address := fmt.Sprintf("%s:%d", "127.0.0.1", s.port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	s.listener = listener

	for {
		err := s.handleConnection()
		if err != nil {
			break
		}
	}
}

func (s *DnsServer) handleConnection() error {
	buf := make([]byte, 512)

	size, source, err := s.listener.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return err
	}

	receivedData := string(buf[:size])
	fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

	// Parse request
	request := Message{}
	request.Unmarshall(buf[:size])
	fmt.Printf("Received message: %#v\n", request)

	responseData := []byte{}
	_, err = s.listener.WriteToUDP(responseData, source)
	if err != nil {
		fmt.Println("Failed to send response:", err)
	}

	return nil
}

func (s *DnsServer) Close() error {
	fmt.Println("Closing DNS server")
	return s.listener.Close()
}

func NewDnsServer(port int) *DnsServer {
	return &DnsServer{
		port: port,
	}
}
