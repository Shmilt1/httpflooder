package main

import "httpflooder/internals"

func main() {
	flooder := &internals.HttpFlooder{
		Host:     "127.0.0.1",
		Port:     80,
		Secure:   false,
		Interval: 2,
		Duration: 60,
	}

	flooder.Flood()
}
