package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d:=func(da dataChan)error{
		for i:=0;i<30;i++{
			da<-i
			log.Printf("Dispacher sent :%v",i)
		}
		return  nil
	}
	e:=func(da dataChan)error{
		forloop:
		for{
			select {
			case d:=<-da:
				log.Printf("Executor received :%v",d)
			default:
				 break forloop
			}
		}
		return errors.New("Close")
	}
	runner:=NewRunner(30,false,d,e)
	go runner.StartAll()
	time.Sleep(3 *time.Second)
}
