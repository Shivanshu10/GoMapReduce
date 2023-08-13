package coordinatorlib

import (
	"GoMapReduce/common/rpcdef"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Coordinator struct {
	tasks []Task
}

func (c *Coordinator) Example(args *rpcdef.ExampleArgs, reply *rpcdef.ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) PrintInfo() {
	for _, t := range c.tasks {
		log.Println(fmt.Sprintf("%#v", t))
	}
}

func (c *Coordinator) AddTask(args *rpcdef.AddTaskArgs, reply *rpcdef.Result) error {
	tasks_count := len(c.tasks)
	n_map := 0

	for _, file := range args.Files {
		c.tasks = append(c.tasks, Task{tasks_count + 1, []string{file}, args.Func_File, IDLE, MAP, "", 0})
		tasks_count += 1
		n_map += 1
	}

	for i := 0; i < args.N_Reduce; i++ {
		var fs []string
		for j := 0; j < n_map; j++ {
			map_res_file := fmt.Sprintf("mr-%v-%v", j, i)
			fs = append(fs, map_res_file)
		}
		c.tasks = append(c.tasks, Task{tasks_count + 1, fs, args.Func_File, IDLE, REDUCE, "", 0})
	}

	c.PrintInfo()

	reply.Status_Code = rpcdef.SUCCESS

	return nil
}

func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	sockname := rpcdef.CoordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}
	c.server()

	return &c
}
