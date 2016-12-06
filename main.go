package main

import (
	"log"
	"log/syslog"

	"github.com/armon/go-socks5"
)

// Most of the job is done with https://github.com/armon/go-socks5
// Thanks armon.

func main() {
	// comment this logWritter part if you only want stdout logger
	logWriter, err := syslog.NewLogger(syslog.LOG_DAEMON|syslog.LOG_ERR, log.Lshortfile)
	if err != nil {
		panic(err)
	}

	// Create a SOCKS5 server
	conf := &socks5.Config{
		Logger: logWriter,
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on all ip4 interfaces port 1080
	if err := server.ListenAndServe("tcp4", ":1080"); err != nil {
		panic(err)
	}
}
