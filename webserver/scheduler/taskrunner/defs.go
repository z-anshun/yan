package taskrunner

const (
	READY_TO_DISPATCH = "d" //读
	READY_TO_EXECUTE  = "e" //执行数据
	CLOSE             = "c"
)

type controlChan chan string

type dataChan chan interface{} //下发的数据

type fn func(dc dataChan) error
