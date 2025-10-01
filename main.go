package main

import (
	"flag"
	"sync"

	"github.com/Shmilt1/httpflooder/internals"
)

func main() {
	host := flag.String("h", "127.0.0.1", "Target host")
	port := flag.Int("p", 80, "Target port")
	duration := flag.Int("d", 60, "Duration of flood")
	interval := flag.Int("i", 0, "Interval per requests")
	threads := flag.Int("t", 2, "Threads")
	secure := flag.Bool("s", false, "Target uses SSL/TLS")
	sockets := flag.Int("c", 1, "How many sockets to use")
	protocol := flag.String("m", "http", "Target protocol")

	flag.Parse()

	if *duration == 0 {
		*duration = 60
	}

	if *sockets == 0 {
		*sockets = 1
	}

	var wg sync.WaitGroup

	if *host != "" && *port != 0 {
		if *threads == 0 {
			flooder := internals.parseFlooderArgs(
				*protocol,
				*host,
				*port,
				*duration,
				*interval,
				*sockets,
				0,
				*secure
			)

			flooder.Flood()
			return
		}
		for i := 0; i < *threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				flooder := internals.parseFlooderArgs(
					*protocol,
					*host,
					*port,
					*duration,
					*interval,
					*sockets,
					i + 1,
					*secure
				)

				flooder.Flood()
			}()
		}
	}

	wg.Wait()
}
