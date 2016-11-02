# Kruispunt controller
Controller voor kruispunt opdracht.
## Dependencies
- [Go](https://golang.org)
- [github.com/gorilla/websocket](https://github.com/gorilla/websocket)
## Usage
`go build` to compile for current OS/arch (optionally set `GOOS` for different OS or `GOARCH` for different CPU architecture (see [this](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5)). Run with `./controller` for listening on port 3000, run with `--addr=$ADDRESS:$PORT` for anything else.