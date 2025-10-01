package internals

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseFlooderArgs(protocol, host string, port, duration, interval, sockets, threadId int, secure bool) Flooder {
	var flooder Flooder

	switch protocol {
		case "http":
			flooder = internals.HttpFlooder{
				Host:     host,
				Port:     port,
				Duration: duration,
				Interval: interval,
				Secure:   secure,
				Sockets:  sockets,
				ThreadID: threadId,
			}
			/*
			planned protocols:
			case "syn":
			case "udp":
			case "dns":
			case "tls":
			*/	
	}
	
	return flooder
}

func combineHost(host string, port int) string {
	parts := strings.Split(host, ":")
	if len(parts) == 8 {
		return "[" + host + "]:" + strconv.Itoa(port)
	}
	return host + ":" + strconv.Itoa(port)
}

const (
	INFO    = "+"
	WARNING = "!"
	ERROR   = "*"
	DEBUG   = "-"
)

// just prints sumethin
func print_sumthin(message, level string) {
	fmt.Println(+time.Now().Unix(), "["+level+"]", message)
}
