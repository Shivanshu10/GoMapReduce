package coordinatorlib

type TaskState byte

const (
	IDLE    TaskState = 0
	RUNNING TaskState = 1
	FINISH  TaskState = 2
)
