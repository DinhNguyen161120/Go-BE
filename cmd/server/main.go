package main

import (
	"example.com/m/internal/initialize"
)

func main() {
	// r := routers.NewRouter()
	// r.Run() // r.Run(":8002") change port, listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	initialize.Run()
}
