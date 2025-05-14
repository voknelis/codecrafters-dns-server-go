package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/codecrafters-io/dns-server-starter-go/app/dns"
)

func main() {
	server := dns.NewDnsServer(2053)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-signalChannel
		server.Close()
	}()

	server.Listen()
}
