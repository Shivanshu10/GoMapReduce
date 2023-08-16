package coordinatorlib

import (
	"GoMapReduce/common/rpc"
	"fmt"
	"log"
)

type Coordinator struct {
	tasks []Task
}

func (c *Coordinator) PrintInfo() {
	log.Println("PrintInfo Called")
	for _, t := range c.tasks {
		log.Println(fmt.Sprintf("%#v", t))
	}
}

func (c *Coordinator) AddTask(args *rpc.AddTaskArgs, reply *rpc.Result) error {
	log.Println("AddTask Called args: " + fmt.Sprintf("%#v", args))
	tasks_count := len(c.tasks)
	n_map := 0

	for _, file := range args.Files {
		c.tasks = append(c.tasks, Task{tasks_count + 1, []string{file}, args.Func_File, IDLE, MAP, "", 0})
		tasks_count += 1
		n_map += 1
	}

	var i int64
	for ; i < args.N_Reduce; i++ {
		var fs []string
		for j := 0; j < n_map; j++ {
			map_res_file := fmt.Sprintf("mr-%v-%v", j, i)
			fs = append(fs, map_res_file)
		}
		c.tasks = append(c.tasks, Task{tasks_count + 1, fs, args.Func_File, IDLE, REDUCE, "", 0})
	}

	c.PrintInfo()

	reply.Status_Code = rpc.SUCCESS

	log.Println("AddTask reply: " + fmt.Sprintf("%#v", reply))

	return nil
}

func (c *Coordinator) RequestTask(workerId string, reply *Task) error {
	log.Println("Request Task Args: " + fmt.Sprintf("%s", workerId))
	for idx, task := range c.tasks {
		if task.Status == IDLE {
			task.Worker_id = workerId
			task.Status = RUNNING
			c.tasks[idx] = task
			*reply = task
			break
		}
	}
	c.PrintInfo()
	return nil
}

func (c *Coordinator) Serve() {
	rpc.Serve(c)
}

func MakeCoordinator() *Coordinator {
	log.Println("Creating Coordinator")
	c := Coordinator{}
	c.Serve()
	log.Println("Serving Coordinator")

	return &c
}
