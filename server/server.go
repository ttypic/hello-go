// Package server greets from the server side.
package server

import hello "github.com/ttypic/hello-go/v2"

// SayHi prints a server prefix followed by the root greeting.
func SayHi() {
	print("[server] ")
	hello.SayHi()
}
