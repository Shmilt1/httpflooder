package main

import (
	"flag"
	"httpflooder/internals"
	"sync"
)

func main() {
	host := flag.String("h", "127.0.0.1", "Target host")
	port := flag.Int("p", 80, "Target port")
	duration := flag.Int("d", 60, "Duration of flood")
	interval := flag.Int("i", 0, "Interval per requests")
	threads := flag.Int("t", 2, "Threads")
	secure := flag.Bool("s", false, "Target is HTTPS")
	sockets := flag.Int("c", 1, "How many sockets to use")

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
			// ...
			flooder := internals.HttpFlooder{
				Host:     *host,
				Port:     *port,
				Duration: *duration,
				Interval: *interval,
				Secure:   *secure,
				Sockets:  *sockets,
				ThreadID: 0,
			}

			flooder.Flood()
			return
		}
		for i := 0; i < *threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				flooder := internals.HttpFlooder{
					Host:     *host,
					Port:     *port,
					Duration: *duration,
					Interval: *interval,
					Secure:   *secure,
					Sockets:  *sockets,
					ThreadID: i + 1,
				}

				flooder.Flood()
			}()
		}
	}

	wg.Wait()
}
