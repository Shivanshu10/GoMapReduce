package rpcdef

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type AddTaskArgs struct {
	Files     []string
	Func_File string
	N_Reduce  int
}

type Result struct {
	Status_Code ResultStatusCode
}

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

func CoordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

func Call(rpcname string, args interface{}, reply interface{}) bool {
	sockname := CoordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
