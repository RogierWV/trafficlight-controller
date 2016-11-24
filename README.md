# Kruispunt controller
Controller for kruispunt assignment.

## Dependencies
- [Go](https://golang.org)
- [github.com/gorilla/websocket](https://github.com/gorilla/websocket)

## Build
`go build` to compile for current OS/arch or optionally set `GOOS` for different OS or `GOARCH` for different CPU architecture (see [this](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5)).

## Options
Usage of ./controller:
  `-addr string`
    	http service address (`"$ADDR:$PORT"`) (default `"0.0.0.0:3000"`)
  `-ft int`
    	frame type (`1` = text, `2` = binary) (default `1`)
  `-ontrtijd int`
    	ontruimingstijd (default `2`)
