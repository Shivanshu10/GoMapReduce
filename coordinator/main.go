package main

import (
	"GoMapReduce/coordinator/coordinatorlib"
	"log"
)

func main() {
	log.Println("[*] Starting Coordinator Server")
	_ = coordinatorlib.MakeCoordinator()
}
