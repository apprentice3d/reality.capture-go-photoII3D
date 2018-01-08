package main

import (
	"github.com/apprentice3d/forge-photoII3D/server"
)

const PORT  = ":80"



func main() {
	server.StartServer(PORT)
}
