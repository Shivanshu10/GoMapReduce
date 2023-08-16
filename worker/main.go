package main

import (
	"GoMapReduce/common/rpc"
	. "GoMapReduce/common/task"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	task := &Task{}
	rpc.Call("Coordinator.RequestTask", uuid.String(), &task)
	log.Println(fmt.Sprintf("%#v", task))

}
