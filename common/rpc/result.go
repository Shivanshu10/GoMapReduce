package rpc

type ResultStatusCode byte

const (
	SUCCESS ResultStatusCode = 0
	FAIL    ResultStatusCode = 1
)

type Result struct {
	Status_Code ResultStatusCode
}
