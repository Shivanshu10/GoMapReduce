package rpc

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

func CoordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

func Serve(c any) {
	rpc.Register(c)
	rpc.HandleHTTP()
	e := http.ListenAndServe(":1234", nil)
	if e != nil {
		log.Fatal("listen error:", e)
	}
}

func Call(rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
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
