package main

import (
	gs "gameserver/src"
)

func main() {
	s := gs.StdServer()

	s.Start()
}
