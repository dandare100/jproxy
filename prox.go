package main

//CGO_ENABLED=0 GOOS=linux go build -a -o prox  .
import (
	"fmt"
	"os"
	"os/signal"

	"github.com/google/tcpproxy"
)

func main() {

	args := os.Args

	if len(args) != 4 {
		fmt.Printf("Usage : %v lport rhost rport\n", args[0])
		os.Exit(69)
	}

	var p tcpproxy.Proxy

	p.AddRoute(":"+args[1], tcpproxy.To(args[2]+":"+args[3]))
	p.Run()

	defer p.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

}
