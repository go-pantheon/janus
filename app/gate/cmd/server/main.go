package main

import (
	"os"
	"os/signal"

	tcp "github.com/vulcan-frame/vulcan-gate/pkg/net/tcp/server"
)

func main() {
	server := tcp.NewServer()
	server.Start()

	signal.Notify(os.Interrupt, os.Kill)
	<-signal.Wait()
	server.Stop()
}
