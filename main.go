package main

import (
	"github.com/apprentice3d/forge-photoII3D/server"
)

const port = ":3000"

func main() {
	server.StartServer(port)
}
