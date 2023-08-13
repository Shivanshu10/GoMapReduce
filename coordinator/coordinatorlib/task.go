package coordinatorlib

type Task struct {
	id             int
	file_name      []string
	func_file_name string
	status         TaskState
	task_type      TaskType
	worker_id      string
	bt             int64
}
