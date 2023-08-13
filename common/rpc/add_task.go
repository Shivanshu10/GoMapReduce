package rpc

type AddTaskArgs struct {
	Files     []string
	Func_File string
	N_Reduce  int64
}

func AddTask(files []string, func_file string, n_reduce int64) Result {
	args := AddTaskArgs{files, func_file, n_reduce}
	result := Result{}
	if Call("Coordinator.AddTask", &args, &result) {
		return result
	}
	result.Status_Code = FAIL
	return result
}
