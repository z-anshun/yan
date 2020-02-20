package taskrunner

type Runner struct {
	ControlChan controlChan
	Error       controlChan //两个chan 便于维护
	Data        dataChan
	dateSize    int
	longLived   bool //判断是否长期存活
	Dispatcher  fn
	Executor    fn
}

//构造函数
func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		ControlChan: make(chan string, 1),
		Error:       make(chan string, 1),
		Data:        make(chan interface{}, size),
		longLived:   longlived,
		dateSize:    size,
		Executor:    e,
		Dispatcher:  d,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.ControlChan)
			close(r.Data)
			close(r.Error)
		}
	}() //自己动
	for {
		select {
		//异步执行
		case c := <-r.ControlChan: //任务内容
		if c==READY_TO_DISPATCH{
			err:=r.Dispatcher(r.Data )
			if err!=nil{
				r.Error<-CLOSE
			}else{
				r.ControlChan<-READY_TO_EXECUTE
			}
		}
		if c==READY_TO_EXECUTE{
			err:=r.Executor(r.Data )
			if err!=nil{
				r.Error<-CLOSE
			}else{
				r.ControlChan<-READY_TO_DISPATCH
			}
		}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}
func (r *Runner)StartAll(){
	r.ControlChan <-READY_TO_DISPATCH
	r.startDispatch()
}