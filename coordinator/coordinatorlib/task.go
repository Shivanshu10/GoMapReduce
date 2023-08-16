package coordinatorlib

type Task struct {
	Id             int
	File_name      []string
	Func_file_name string
	Status         TaskState
	Task_type      TaskType
	Worker_id      string
	Bt             int64
}
