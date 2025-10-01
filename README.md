![httpflooder](assets/httpflooder.png)
# httpflooder
simple http flooder

## table of contents
- [building](#building)
- [notice](#notice)
- [docs](#docs)
- [example](#example)
- [license](#license)

## building
you need to own the Go runtime if you want to compile / build this project as well.. it's written in Go,
you can download and install the Go runtime at: https://go.dev/dl/.

first clone the project and move into the project's directory:
```bash
git clone https://github.com/Shmilt1/httpflooder.git
cd httpflooder/
```

then run this in your terminal / shell to compile / build the project:
```bash
go build
```

You should now see (```httpflooder``` on Linux) or (```httpflooder.exe``` on Windows) as the output, now refer to [example](#example) or [docs](#docs). 

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
./httpflooder -d 120 -i 1 -t 4 -c 8 -h 127.0.0.1 -p 8080 -m http
```

Windows:
```bash
.\httpflooder.exe -d 120 -i 1 -t 4 -c 8 -h 127.0.0.1 -p 8080 -m http
```

## license
MIT License, read [LICENSE](LICENSE).
