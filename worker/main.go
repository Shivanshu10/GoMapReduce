package main

import (
	"GoMapReduce/common/rpc"
	"GoMapReduce/coordinator/coordinatorlib"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	task := &coordinatorlib.Task{}
	rpc.Call("Coordinator.RequestTask", uuid.String(), &task)
	log.Println(fmt.Sprintf("%#v", task))

}
