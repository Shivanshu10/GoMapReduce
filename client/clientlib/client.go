package clientlib

import (
	"GoMapReduce/common/rpc"
	"log"
)

func AddTask(files []string, func_file string, n_reduce int64) bool {
	log.Println("Calling AddTask")
	result := rpc.AddTask(files, func_file, n_reduce)
	log.Println("AddTask finished")
	return result.Status_Code == rpc.SUCCESS
}
