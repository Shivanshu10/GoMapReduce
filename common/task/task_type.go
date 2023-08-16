package task

type TaskType byte

const (
	MAP    TaskType = 0
	REDUCE TaskType = 1
)
