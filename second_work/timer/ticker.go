package timer

import (
	"src/second_work/db"
	"src/second_work/limite"
	"src/second_work/model"
	"time"
)

func TickerStart() {
	//假设比赛时间进行3天

	t := time.NewTicker(time.Millisecond*1000*60*24*3 - 200)//0.2分钟就是12秒，给点时间创表
	//给点时间创表
	for {
		select {
		case <-t.C:
			{
				db.FinalMatch()
				////创建管道
				model.Coon=limite.CreatChan()
				model.Final = true          //那边可以开始了
				model.T = time.Now().Unix() //开始计时
				return
			}
		}
	}
}
