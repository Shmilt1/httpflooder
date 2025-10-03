package internals

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseFlooderArgs(protocol, host string, port, duration, interval, sockets, threadId int, secure bool) Flooder {
	var flooder Flooder

	switch protocol {
	case "http":
		flooder = &HttpFlooder{
			Host:     host,
			Port:     port,
			Duration: duration,
			Interval: interval,
			Sockets:  sockets,
			ThreadID: threadId,
			Secure:   secure,
		}
	case "udp":
		flooder = &UdpFlooder{
			Host:     host,
			Port:     port,
			Duration: duration,
			Interval: interval,
			Sockets:  sockets,
			ThreadID: threadId,
		}
		/*
			planned protocols:
			case "syn":
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
