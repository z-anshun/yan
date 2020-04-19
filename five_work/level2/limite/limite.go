package limite

type ConnLimiter struct {
	coon chan int
}

func NewConn() *ConnLimiter {
	return &ConnLimiter{
		make(chan int, 20),
	}
}

func (c *ConnLimiter) PutConn() {
	for {
		if len(c.coon) < 20 {
			c.coon <- 1
			return
		} else {

			continue
		}

	}
}

//释放
func (c *ConnLimiter) ReleaseConn() {
	<-c.coon

}
