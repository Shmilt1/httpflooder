package internals

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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
