package workerlib

import (
	"GoMapReduce/common/rpcdef"
	"fmt"
	"hash/fnv"
	"log"
	"net/rpc"
)

type KeyValue struct {
	Key   string
	Value string
}

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	CallExample()

}

func CallExample() {
	args := rpcdef.ExampleArgs{}

	args.X = 99

	reply := rpcdef.ExampleReply{}

	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

func call(rpcname string, args interface{}, reply interface{}) bool {
	sockname := rpcdef.CoordinatorSock()
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
