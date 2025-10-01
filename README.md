# httpflooder
simple http flooder

## table of contents
- [building](#building)
- [notice](#notice)
- [docs](#docs)
- [example](#example)
- [license](#license)

## building
you need to own the Go runtime if you want to compile / build this project as well.. it's written in Go.

run:
```bash
go build
```

## notice
you agree that when using this tool, the maintainer will NOT be held responsible for any misuse!!! ⚠️

please don't use this tool on anyone without explicit permission first!!! ⚠️

## docs
```
OPTIONS: [
  -h = target host,
  -p = target port,
  -d = duration of flood in seconds,
  -i = interval per request,
  -t = how many threads to use,
  -s = target uses TLS/SSL (will be replaced as 'tls' in the -m option later on),
  -m = target protocol
]
```

## example
Linux:
```bash
./httpflooder -d 120 -i 1 -t 4 -c 8 -h 127.0.0.1 -p 8080
```

Windows:
```bash
.\httpflooder.exe -d 120 -i 1 -t 4 -c 8 -h 127.0.0.1 -p 8080
```

## license
MIT License, read [LICENSE](LICENSE).
